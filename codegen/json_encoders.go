package codegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type JSONEncoders struct {
	Directory string
	Package   string
	bytes.Buffer
	structValue reflect.Value
}

func NewJSONEncoders(directory string, packageName string) *JSONEncoders {
	return &JSONEncoders{Directory: directory, Package: packageName}
}

func (c *JSONEncoders) PackageDeclaration() {
	c.WriteString(fmt.Sprintf("package %s\n\n", c.Package))
}

func (c *JSONEncoders) ImportDeclaration() {
	c.WriteString("import \"github.com/marcel/jitjson/encoding\"\n\n")
}

// TODO Add func NewEncodingBuffer to wrap bufferPool.GetBuffer()
func (c *JSONEncoders) SetBufferPoolVar() {
	c.WriteString("var bufferPool = encoding.NewSyncPool(4096)\n\n")
}

func (c *JSONEncoders) EncodingBufferStructWrapper() {
	c.WriteString("type encodingBuffer struct {\n\t*encoding.Buffer\n}\n\n")
}

func (c *JSONEncoders) JSONMarshalerInterfaceFor(structName string) {
	buf := bytes.Buffer{}

	buf.WriteString(fmt.Sprintf("func (s %s) MarshalJSON() ([]byte, error) {\n", structName))
	buf.WriteString("\tunderlying := bufferPool.GetBuffer()\n")
	buf.WriteString("\tbuf := encodingBuffer{Buffer: underlying}\n")
	buf.WriteString("\tdefer func() {\n")
	buf.WriteString("\t\tunderlying.Reset()\n")
	buf.WriteString("\t\tbufferPool.PutBuffer(underlying)\n")
	buf.WriteString("\t}()\n\n")

	buf.WriteString(fmt.Sprintf("\tbuf.%sStruct(s)\n", strings.ToLower(structName)))
	buf.WriteString("\treturn buf.Bytes(), nil\n")
	buf.WriteString("}\n\n")

	c.Write(buf.Bytes())
}

var JSONEncodersTargetFile = "json_encoders.go"

func (c *JSONEncoders) WriteFile() error {
	targetPath := filepath.Join(c.Directory, JSONEncodersTargetFile)
	return ioutil.WriteFile(targetPath, c.Bytes(), 0644)
}

type encodableStructSpec struct {
	value      reflect.Value
	jsonFields []reflect.StructField
}

func newEncodableStructSpec(value reflect.Value) *encodableStructSpec {
	return &encodableStructSpec{value, []reflect.StructField{}}
}

func (c *JSONEncoders) EncoderMethodFor(jsonStruct interface{}) error {
	value := reflect.ValueOf(jsonStruct)

	var structSpec *encodableStructSpec

	switch value.Kind() {
	default:
		return fmt.Errorf("EncoderMethodFor: unsupported kind '%s'", value.Kind().String())
	case reflect.Ptr:
		return c.EncoderMethodFor(value.Elem().Interface())
	case reflect.Struct:
		structSpec = newEncodableStructSpec(value)
		for i := 0; i < value.NumField(); i++ {
			// TODO Handled structField.Anonymous == true
			structField := value.Type().Field(i)

			if _, ok := structField.Tag.Lookup("json"); ok {
				structSpec.jsonFields = append(structSpec.jsonFields, structField)
			}
		}
	}

	if len(structSpec.jsonFields) == 0 {
		return nil
	}

	c.generateMethodForStruct(structSpec)

	return nil
}

func (c *JSONEncoders) generateMethodForStruct(structSpec *encodableStructSpec) {
	c.structValue = structSpec.value

	c.methodDeclaration()

	c.encoderInvoke("OpenBrace")

	for index, field := range structSpec.jsonFields {
		if index != 0 {
			c.encoderInvoke("Comma")
		}
		c.fieldEncodingFor(field)
	}

	c.WriteString("\n")
	c.encoderInvoke("CloseBrace")
	c.endMethod()
}

func (c *JSONEncoders) encoderInvoke(method string) {
	c.WriteString(fmt.Sprintf("  e.%s()\n", method))
}

func (c *JSONEncoders) methodDeclaration() {
	methodDecl := fmt.Sprintf(
		"func (e *encodingBuffer) %sStruct(%s %s) {\n",
		c.structName(), c.structName(), c.structTypeName(),
	)

	c.WriteString(methodDecl)
}

func (c *JSONEncoders) fieldEncodingFor(field reflect.StructField) {
	var attrName string

	attrName = field.Tag.Get("json")
	if attrName == "" {
		attrName = strings.ToLower(field.Name)
	}

	quotedAttr := strconv.Quote(attrName) + ":"
	c.WriteString(fmt.Sprintf("\n  // %s\n  e.Write(%#v)\n", quotedAttr, []byte(quotedAttr)))

	// TODO Needs to be refactored to recursively support nested collections like
	// a slice of a slice, etc
	switch field.Type.Kind() {
	default:
		log.Println("Unsupported field kind", field.Type.Kind())
	case reflect.Bool:
		c.invokeEncoderForFieldType("Bool", field)
	case reflect.String:
		c.stringFieldEncoding(field)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		c.intFieldEncoding(field)
	case reflect.Float32, reflect.Float64:
		c.floatFieldEncoding(field)
	case reflect.Struct:
		c.structFieldEncoding(field)
	case reflect.Slice, reflect.Array:
		c.sliceFieldEncoding(field)
	}
}

func (c *JSONEncoders) dispatch(field reflect.StructField) string {
	return fmt.Sprintf("%s.%s", c.structName(), field.Name)
}

func (c *JSONEncoders) stringFieldEncoding(field reflect.StructField) {
	c.invokeEncoderForFieldType("Quote", field)
}

func (c *JSONEncoders) intFieldEncoding(field reflect.StructField) {
	var specializedIntEncoder string

	switch field.Type.Kind() {
	case reflect.Int64:
		specializedIntEncoder = "Int64"
	case reflect.Uint64:
		specializedIntEncoder = "Uint64"
	case reflect.Int32:
		specializedIntEncoder = "Int32"
	case reflect.Uint32:
		specializedIntEncoder = "Uint32"
	case reflect.Int16:
		specializedIntEncoder = "Int16"
	case reflect.Uint16:
		specializedIntEncoder = "Uint16"
	case reflect.Int8:
		specializedIntEncoder = "Int8"
	case reflect.Uint8:
		specializedIntEncoder = "Uint8"
	case reflect.Int:
		specializedIntEncoder = "Int"
	case reflect.Uint:
		specializedIntEncoder = "Uint"
	}

	var code string
	if field.Type.String() == field.Type.Kind().String() {
		code = fmt.Sprintf("  e.%s(%s)\n",
			specializedIntEncoder, c.dispatch(field))
	} else {
		code = fmt.Sprintf("  e.%s(%s(%s))\n",
			specializedIntEncoder, strings.ToLower(specializedIntEncoder), c.dispatch(field))
	}
	c.WriteString(code)
}

func (c *JSONEncoders) floatFieldEncoding(field reflect.StructField) {
	var specializedFloatEncoder string

	switch field.Type.Kind() {
	case reflect.Float64:
		specializedFloatEncoder = "Float64"
	case reflect.Float32:
		specializedFloatEncoder = "Float32"
	}

	code := fmt.Sprintf("  e.%s(%s(%s))\n",
		specializedFloatEncoder, strings.ToLower(specializedFloatEncoder), c.dispatch(field))

	c.WriteString(code)

}

func (c *JSONEncoders) structFieldEncoding(field reflect.StructField) {
	// TODO Consolidate this and the duplicate code in sliceFieldEncoding
	jsonMarshalerType := reflect.TypeOf(new(json.Marshaler)).Elem()
	if field.Type.Implements(jsonMarshalerType) {
		buf := bytes.Buffer{}
		buf.WriteString(fmt.Sprintf("  jsonBytes, err := %s.MarshalJSON()\n", c.dispatch(field)))
		buf.WriteString("  if err != nil {\n    panic(err)\n  }\n")
		buf.WriteString("  e.Write(jsonBytes)\n")
		c.Write(buf.Bytes())
		return
	}

	targetStruct := strings.ToLower(field.Type.Name())

	code := fmt.Sprintf("  e.%sStruct(%s)\n", targetStruct, c.dispatch(field))
	c.WriteString(code)
}

func (c *JSONEncoders) lowerCase(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func (c *JSONEncoders) sliceFieldEncoding(field reflect.StructField) {
	c.WriteString("  e.WriteByte('[')\n")
	forLoopLine := fmt.Sprintf(
		"  for index, element := range %s {\n", c.dispatch(field),
	)

	c.WriteString(forLoopLine)
	c.WriteString("    if index != 0 { e.Comma() }\n")

	elementType := field.Type.Elem()
	switch elementType.Kind() {
	case reflect.Struct:
		jsonMarshalerType := reflect.TypeOf(new(json.Marshaler)).Elem()
		if elementType.Implements(jsonMarshalerType) {
			code :=
				`    jsonBytes, err := element.MarshalJSON()
    if err != nil {
      panic(err)
    }
    e.Write(jsonBytes)
`
			c.WriteString(code)
			return
		}
		structName := c.lowerCase(elementType.Name())
		code := fmt.Sprintf("    e.%sStruct(element)\n", structName)
		c.WriteString(code)
	default:
		var code string
		// TODO get rid of this special case to work around Quote() asymetry
		if elementType.Kind() == reflect.String {
			code = "    e.Quote(string(element))\n"
		} else {
			encoderFromKind := strings.Title(elementType.Kind().String())
			code = fmt.Sprintf("    e.%s(%s(element))\n", encoderFromKind, strings.ToLower(encoderFromKind))
		}
		c.WriteString(code)
	}

	c.WriteString("  }\n")
	c.WriteString("  e.WriteByte(']')\n")
}

func (c *JSONEncoders) invokeEncoderForFieldType(fieldType string, field reflect.StructField) {
	var code string

	if field.Type.Kind() == reflect.String {
		if field.Type.String() == field.Type.Kind().String() {
			code = fmt.Sprintf("  e.%s(%s)\n", fieldType, c.dispatch(field))
		} else {
			code = fmt.Sprintf("  e.%s(string(%s))\n", fieldType, c.dispatch(field))
		}
	} else {
		code = fmt.Sprintf("  e.%s(%s(%s))\n", fieldType, strings.ToLower(fieldType), c.dispatch(field))
	}
	c.WriteString(code)
}

func (c *JSONEncoders) endMethod() {
	c.WriteString("}\n\n")
}

func (c *JSONEncoders) structTypeName() string {
	return c.structValue.Type().Name()
}

func (c *JSONEncoders) structName() string {
	return c.lowerCase(c.structTypeName())
}
