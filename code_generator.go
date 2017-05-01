package jitjson

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type CodeGenerator struct {
	bytes.Buffer
	structValue reflect.Value
}

func (c *CodeGenerator) PackageDelcaration() {
	// TODO needs to be the actual right package name
	c.WriteString("package jitjson\n\n")
}

func (c *CodeGenerator) ImportDeclaration() {
	// TODO this is just a test impl, needs to be updated eventually
	c.WriteString("import . \"github.com/marcel/jitjson/fixtures\"\n\n")
}

type encodableStructSpec struct {
	value      reflect.Value
	jsonFields []reflect.StructField
}

func newEncodableStructSpec(value reflect.Value) *encodableStructSpec {
	return &encodableStructSpec{value, []reflect.StructField{}}
}

func (c *CodeGenerator) EncoderMethodFor(jsonStruct interface{}) error {
	value := reflect.ValueOf(jsonStruct)

	var structSpec *encodableStructSpec

	switch value.Kind() {
	default:
		fmt.Println("unrecognized kind", value.Kind())
		return nil
	case reflect.Ptr:
		c.EncoderMethodFor(value.Elem())
	case reflect.Struct:
		structSpec = newEncodableStructSpec(value)
		for i := 0; i < value.NumField(); i++ {
			// TODO Handled structField.Anonymous == true
			structField := value.Type().Field(i)

			if structField.Tag.Get("json") != "" {
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

func (c *CodeGenerator) generateMethodForStruct(structSpec *encodableStructSpec) {
	c.structValue = structSpec.value

	c.methodDeclaration()

	c.encoderInvoke("openBrace")

	for index, field := range structSpec.jsonFields {
		if index != 0 {
			c.encoderInvoke("comma")
		}
		c.fieldEncodingFor(field)
	}

	c.WriteString("\n")
	c.encoderInvoke("closeBrace")
	c.endMethod()
}

func (c *CodeGenerator) encoderInvoke(method string) {
	c.WriteString(fmt.Sprintf("  e.%s()\n", method))
}

func (c *CodeGenerator) methodDeclaration() {
	methodDecl := fmt.Sprintf(
		"func (e *structEncoder) %sStruct(%s %s) {\n",
		c.structName(), c.structName(), c.structTypeName(),
	)

	c.WriteString(methodDecl)
}

func (c *CodeGenerator) fieldEncodingFor(field reflect.StructField) {
	var attrName string

	attrName = field.Tag.Get("json")
	if attrName == "" {
		attrName = strings.ToLower(field.Name)
	}

	c.WriteString(fmt.Sprintf("\n  e.attr(\"%s\")\n", attrName))

	// TODO Needs to be refactored to recursively support nested collections like
	// a slice of a slice, etc
	switch field.Type.Kind() {
	default:
		log.Println("Unsupported field kind", field.Type.Kind())
	case reflect.Bool:
		c.invokeEncoderForFieldType("bool", field)
	case reflect.String:
		c.stringFieldEncoding(field)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		c.intFieldEncoding(field)
	case reflect.Struct:
		c.structFieldEncoding(field)
	case reflect.Slice, reflect.Array:
		c.sliceFieldEncoding(field)
	}
}

func (c *CodeGenerator) dispatch(field reflect.StructField) string {
	return fmt.Sprintf("%s.%s", c.structName(), field.Name)
}

func (c *CodeGenerator) stringFieldEncoding(field reflect.StructField) {
	c.invokeEncoderForFieldType("string", field)
}

func (c *CodeGenerator) intFieldEncoding(field reflect.StructField) {
	var specializedIntEncoder string

	switch field.Type.Kind() {
	case reflect.Int64:
		specializedIntEncoder = "int64"
	case reflect.Int32:
		specializedIntEncoder = "int32"
	case reflect.Int16:
		specializedIntEncoder = "int16"
	case reflect.Int8:
		specializedIntEncoder = "int8"
	case reflect.Int:
		specializedIntEncoder = "int"
	}

	code := fmt.Sprintf("  e.%s(%s(%s))\n",
		specializedIntEncoder, specializedIntEncoder, c.dispatch(field),
	)

	c.WriteString(code)
}

func (c *CodeGenerator) structFieldEncoding(field reflect.StructField) {
	targetStruct := strings.ToLower(field.Name)

	code := fmt.Sprintf("  e.%sStruct(%s)\n", targetStruct, c.dispatch(field))
	c.WriteString(code)
}

func (c *CodeGenerator) sliceFieldEncoding(field reflect.StructField) {
	c.WriteString("  e.WriteByte('[')\n")
	forLoopLine := fmt.Sprintf(
		"  for index, element := range %s {\n", c.dispatch(field),
	)

	c.WriteString(forLoopLine)
	c.WriteString("    if index != 0 { e.comma() }\n")

	elementType := field.Type.Elem()

	switch elementType.Kind() {
	case reflect.Struct:
		structName := strings.ToLower(elementType.Name())
		code := fmt.Sprintf("    e.%sStruct(element)\n", structName)
		c.WriteString(code)
	default:
		encoderFromKind := strings.ToLower(elementType.Kind().String())
		code := fmt.Sprintf("    e.%s(element)\n", encoderFromKind)
		c.WriteString(code)
	}

	c.WriteString("  }\n")
	c.WriteString("  e.WriteByte(']')\n")
}

func (c *CodeGenerator) invokeEncoderForFieldType(fieldType string, field reflect.StructField) {
	code := fmt.Sprintf("  e.%s(%s)\n", fieldType, c.dispatch(field))
	c.WriteString(code)
}

func (c *CodeGenerator) endMethod() {
	c.WriteString("}\n\n")
}

func (c *CodeGenerator) structTypeName() string {
	return c.structValue.Type().Name()
}

func (c *CodeGenerator) structName() string {
	// TODO this needs to be only the first letter to lower
	return strings.ToLower(c.structTypeName())
}
