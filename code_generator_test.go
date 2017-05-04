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

  // "some_bool":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x22, 0x3a})
  e.Bool(testJSONStruct.SomeBool)
  e.Comma()

  // "some_int":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x22, 0x3a})
  e.Int(testJSONStruct.SomeInt)
  e.Comma()

  // "some_int_8":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x38, 0x22, 0x3a})
  e.Int8(testJSONStruct.SomeInt8)
  e.Comma()

  // "some_int_16":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x31, 0x36, 0x22, 0x3a})
  e.Int16(testJSONStruct.SomeInt16)
  e.Comma()

  // "some_int_32":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x33, 0x32, 0x22, 0x3a})
  e.Int32(testJSONStruct.SomeInt32)
  e.Comma()

  // "some_int_64":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x36, 0x34, 0x22, 0x3a})
  e.Int64(testJSONStruct.SomeInt64)
  e.Comma()

  // "some_uint_":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x5f, 0x22, 0x3a})
  e.Uint(testJSONStruct.SomeUint)
  e.Comma()

  // "some_uint_8":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x5f, 0x38, 0x22, 0x3a})
  e.Uint8(testJSONStruct.SomeUint8)
  e.Comma()

  // "some_uint_16":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x5f, 0x31, 0x36, 0x22, 0x3a})
  e.Uint16(testJSONStruct.SomeUint16)
  e.Comma()

  // "some_uint_32":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x5f, 0x33, 0x32, 0x22, 0x3a})
  e.Uint32(testJSONStruct.SomeUint32)
  e.Comma()

  // "some_uint_64":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x5f, 0x36, 0x34, 0x22, 0x3a})
  e.Uint64(testJSONStruct.SomeUint64)
  e.Comma()

  // "some_float_32":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x33, 0x32, 0x22, 0x3a})
  e.Float32(float32(testJSONStruct.SomeFloat32))
  e.Comma()

  // "some_float_64":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x36, 0x34, 0x22, 0x3a})
  e.Float64(float64(testJSONStruct.SomeFloat64))
  e.Comma()

  // "some_string":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x3a})
  e.String(testJSONStruct.SomeString)
  e.Comma()

  // "another_struct":
  e.Write([]byte{0x22, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x3a})
  e.anothertestjsonstructStruct(testJSONStruct.AnotherStruct)
  e.Comma()

  // "some_struct_slice":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range testJSONStruct.SomeStructSlice {
    if index != 0 { e.Comma() }
    e.anotherTestJSONStructStruct(element)
  }
  e.WriteByte(']')
  e.Comma()

  // "some_array":
  e.Write([]byte{0x22, 0x73, 0x6f, 0x6d, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22, 0x3a})
  e.WriteByte('[')
  for index, element := range testJSONStruct.SomeArray {
    if index != 0 { e.Comma() }
    e.String(string(element))
  }
  e.WriteByte(']')
  e.Comma()

  // "nonameoverride":
  e.Write([]byte{0x22, 0x6e, 0x6f, 0x6e, 0x61, 0x6d, 0x65, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x22, 0x3a})
  e.String(testJSONStruct.NoNameOverride)
  e.Comma()

  // "implements_marshaler":
  e.Write([]byte{0x22, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x65, 0x72, 0x22, 0x3a})
  jsonBytes, err := testJSONStruct.ImplementsMarshaler.MarshalJSON()
  if err != nil {
    panic(err)
  }
  e.Write(jsonBytes)
  e.Comma()

  // "slice_of_impls_marsh":
  e.Write([]byte{0x22, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x73, 0x5f, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x22, 0x3a})
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
