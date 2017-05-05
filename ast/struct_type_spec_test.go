package ast

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StructTypeSpecTestSuite struct {
	suite.Suite
}

func TestStructTypeSpecTestSuite(t *testing.T) {
	suite.Run(t, new(StructTypeSpecTestSuite))
}

func (s *StructTypeSpecTestSuite) TestName() {
	spec, err := FindJSONStructFor("github.com/marcel/jitjson/fixtures/media", "Album")
	s.Nil(err)

	s.Equal("Album", spec.Name())
}
