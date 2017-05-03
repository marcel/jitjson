package jitjson

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CodeGeneratorTestSuite struct {
	suite.Suite
	generator *CodeGenerator
}

func TestCodeGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(CodeGeneratorTestSuite))
}

func (s *CodeGeneratorTestSuite) SetupTest() {
	s.generator = NewCodeGenerator("somedir", "somepackage")
}

func (s *CodeGeneratorTestSuite) TestPackageDeclaration() {
	s.generator.PackageDeclaration()
	s.Equal("package somepackage\n\n", s.generator.String())
}

func (s *CodeGeneratorTestSuite) TestImportDeclaration() {
	s.generator.ImportDeclaration()
	s.Equal("import \"github.com/marcel/jitjson/encoding\"\n\n", s.generator.String())
}

func (s *CodeGeneratorTestSuite) TestEncodingBufferStructWrapper() {
	s.generator.EncodingBufferStructWrapper()

	expected :=
		`type encodingBuffer struct {
	*encoding.Buffer
}

`

	s.Equal(expected, s.generator.String())
}

func (s *CodeGeneratorTestSuite) TestJSONMarshlerInterface() {
	s.generator.JSONMarshalerInterfaceFor("SomeStructName")

	expected :=
		`func (s SomeStructName) MarshalJSON() ([]byte, error) {
	underlying := bufferPool.GetBuffer()
	buf := encodingBuffer{Buffer: underlying}
	defer func() {
		underlying.Reset()
		bufferPool.PutBuffer(underlying)
	}()

	buf.somestructnameStruct(s)
	return buf.Bytes(), nil
}

`
	s.Equal(expected, s.generator.String())
}

// TODO TestWriteFile

type TestJSONStruct struct {
	SomeBool            bool                    `json:"some_bool"`
	SomeInt             int                     `json:"some_int"`
	SomeInt8            int8                    `json:"some_int_8"`
	SomeInt16           int16                   `json:"some_int_16"`
	SomeInt32           int32                   `json:"some_int_32"`
	SomeInt64           int64                   `json:"some_int_64"`
	SomeUint            uint                    `json:"some_uint_"`
	SomeUint8           uint8                   `json:"some_uint_8"`
	SomeUint16          uint16                  `json:"some_uint_16"`
	SomeUint32          uint32                  `json:"some_uint_32"`
	SomeUint64          uint64                  `json:"some_uint_64"`
	SomeFloat32         float32                 `json:"some_float_32"`
	SomeFloat64         float64                 `json:"some_float_64"`
	SomeString          string                  `json:"some_string"`
	AnotherStruct       AnotherTestJSONStruct   `json:"another_struct"`
	SomeStructSlice     []AnotherTestJSONStruct `json:"some_struct_slice"`
	SomeArray           [4]string               `json:"some_array"`
	NoNameOverride      string                  `json:""`
	ImplementsMarshaler ImplementsMarshaler     `json:"implements_marshaler"`
	SliceOfImplsMarsh   []ImplementsMarshaler   `json:"slice_of_impls_marsh"`
}

type AnotherTestJSONStruct struct {
	StringField string `json:"string_field"`
}

type ImplementsMarshaler struct {
	SomeField int `json:"some_field"`
}

func (i ImplementsMarshaler) MarshalJSON() ([]byte, error) {
	return []byte{}, nil
}

func (s *CodeGeneratorTestSuite) TestEncoderMethodFor() {
	jsonStruct := TestJSONStruct{}

	expected :=
		`func (e *encodingBuffer) testJSONStructStruct(testJSONStruct TestJSONStruct) {
  e.OpenBrace()

  e.Attr("some_bool")
  e.Bool(bool(testJSONStruct.SomeBool))
  e.Comma()

  e.Attr("some_int")
  e.Int(int(testJSONStruct.SomeInt))
  e.Comma()

  e.Attr("some_int_8")
  e.Int8(int8(testJSONStruct.SomeInt8))
  e.Comma()

  e.Attr("some_int_16")
  e.Int16(int16(testJSONStruct.SomeInt16))
  e.Comma()

  e.Attr("some_int_32")
  e.Int32(int32(testJSONStruct.SomeInt32))
  e.Comma()

  e.Attr("some_int_64")
  e.Int64(int64(testJSONStruct.SomeInt64))
  e.Comma()

  e.Attr("some_uint_")
  e.Uint(uint(testJSONStruct.SomeUint))
  e.Comma()

  e.Attr("some_uint_8")
  e.Uint8(uint8(testJSONStruct.SomeUint8))
  e.Comma()

  e.Attr("some_uint_16")
  e.Uint16(uint16(testJSONStruct.SomeUint16))
  e.Comma()

  e.Attr("some_uint_32")
  e.Uint32(uint32(testJSONStruct.SomeUint32))
  e.Comma()

  e.Attr("some_uint_64")
  e.Uint64(uint64(testJSONStruct.SomeUint64))
  e.Comma()

  e.Attr("some_float_32")
  e.Float32(float32(testJSONStruct.SomeFloat32))
  e.Comma()

  e.Attr("some_float_64")
  e.Float64(float64(testJSONStruct.SomeFloat64))
  e.Comma()

  e.Attr("some_string")
  e.String(string(testJSONStruct.SomeString))
  e.Comma()

  e.Attr("another_struct")
  e.anothertestjsonstructStruct(testJSONStruct.AnotherStruct)
  e.Comma()

  e.Attr("some_struct_slice")
  e.WriteByte('[')
  for index, element := range testJSONStruct.SomeStructSlice {
    if index != 0 { e.Comma() }
    e.anotherTestJSONStructStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("some_array")
  e.WriteByte('[')
  for index, element := range testJSONStruct.SomeArray {
    if index != 0 { e.Comma() }
    e.String(string(element))
  }
  e.WriteByte(']')
  e.Comma()

  e.Attr("nonameoverride")
  e.String(string(testJSONStruct.NoNameOverride))
  e.Comma()

  e.Attr("implements_marshaler")
  jsonBytes, err := testJSONStruct.ImplementsMarshaler.MarshalJSON()
  if err != nil {
    panic(err)
  }
  e.Write(jsonBytes)
  e.Comma()

  e.Attr("slice_of_impls_marsh")
  e.WriteByte('[')
  for index, element := range testJSONStruct.SliceOfImplsMarsh {
    if index != 0 { e.Comma() }
    jsonBytes, err := element.MarshalJSON()
    if err != nil {
      panic(err)
    }
    e.Write(jsonBytes)

  e.CloseBrace()
}

`
	s.Nil(s.generator.EncoderMethodFor(jsonStruct))
	s.Equal(expected, s.generator.String())

	// Passing in a pointer to a struct also works
	s.generator.Reset()
	s.Nil(s.generator.EncoderMethodFor(&jsonStruct))
	s.Equal(expected, s.generator.String())

	// Structs with no json tags are skipped
	s.generator.Reset()
	s.Nil(s.generator.EncoderMethodFor(struct{ SomeField string }{}))
	s.Equal("", s.generator.String())

	// Unsupported kinds of types return an error
	s.generator.Reset()
	s.NotNil(s.generator.EncoderMethodFor(1))
	s.NotNil(s.generator.EncoderMethodFor("foo"))
	s.NotNil(s.generator.EncoderMethodFor(1.23))
	s.NotNil(s.generator.EncoderMethodFor([]int{1, 2, 3}))
	s.Equal("", s.generator.String())
}
