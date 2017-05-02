package jitjson

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var specFixtures *SpecFixture

func NewSpecFixture() *SpecFixture {
	return &SpecFixture{
		Specs: make(map[string]*StructTypeSpec),
	}
}

var Spec func(string) *StructTypeSpec

type StructTypeSpecTestSuite struct {
	suite.Suite
}

func TestStructTypeSpecTestSuite(t *testing.T) {
	suite.Run(t, new(StructTypeSpecTestSuite))
}

func (s *StructTypeSpecTestSuite) TestName() {
	s.Equal("Album", Spec("Album").Name())
}

type SpecFixture struct {
	Specs map[string]*StructTypeSpec
}

func (sf *SpecFixture) Add(spec *StructTypeSpec) {
	sf.Specs[spec.Name()] = spec
}

func (sf *SpecFixture) Spec(name string) *StructTypeSpec {
	spec, ok := sf.Specs[name]

	if !ok {
		panic("No spec fixture named: " + name)
	}

	return spec
}

func init() {
	finder := NewJSONStructFinder()
	finder.FindInDir("fixtures")

	specFixtures = NewSpecFixture()

	specs := finder.StructTypeSpecs()
	for i := range specs {
		specFixtures.Add(&specs[i])
	}

	Spec = func(name string) *StructTypeSpec {
		return specFixtures.Spec(name)
	}
}
