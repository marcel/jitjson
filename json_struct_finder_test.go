package jitjson

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/suite"
)

type JSONStructFinderTestSuite struct {
	suite.Suite
	gocode []byte
}

func TestJSONStructFinderTestSuite(t *testing.T) {
	suite.Run(t, new(JSONStructFinderTestSuite))
}

func (s *JSONStructFinderTestSuite) SetupSuite() {
	code, err := ioutil.ReadFile("fixtures/structs.go")
	s.Nil(err)

	s.gocode = code
}

func (s *JSONStructFinderTestSuite) TestStructs() {
	structFinder := new(JSONStructFinder)
	structs := structFinder.Structs(s.gocode)
	s.Equal(3, len(structs))
}
