package jitjson

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MetaCodeGeneratorTestSuite struct {
	suite.Suite
	generator *MetaCodeGenerator
}

func TestMetaCodeGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(MetaCodeGeneratorTestSuite))
}

func (s *MetaCodeGeneratorTestSuite) SetupTest() {
	structDir := StructDirectory{
		ProjectRoot: "/path/to/project/src/github.com/marcel/jitson",
		PackageRoot: "github.com/marcel/jitjson/fixtures",
		Package:     "media",
		Directory:   "/path/to/project/src/github.com/marcel/jitson/fixtures/media",
		Specs:       []StructTypeSpec{*Spec("Album")},
	}

	s.generator = NewMetaCodeGenerator(structDir)
}

func (s *MetaCodeGeneratorTestSuite) TestMetaCodeGenerator() {
	buf := new(bytes.Buffer)

	s.Nil(s.generator.WriteTo(buf))

	expected :=
		`package main

import (
	"github.com/marcel/jitjson"
	"github.com/marcel/jitjson/fixtures/media"
)

func main() {
	codeGen := jitjson.NewCodeGenerator("/path/to/project/src/github.com/marcel/jitson/fixtures/media", "media")
	codeGen.PackageDeclaration()
	codeGen.ImportDeclaration()	
	codeGen.SetBufferPoolVar()
	codeGen.EncodingBufferStructWrapper()
	codeGen.JSONMarshalerInterfaceFor("Album")
	codeGen.EncoderMethodFor(media.Album{})

	codeGen.WriteFile()
}
`
	s.Equal(expected, buf.String())
}
