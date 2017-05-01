package jitjson

import (
	"io/ioutil"
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
	code, err := ioutil.ReadFile("fixtures/structs.go")
	if err != nil {
		panic(err)
	}

	finder := new(JSONStructFinder)

	specs := finder.Structs(code)

	specFixtures = NewSpecFixture()

	for i := range specs {
		specFixtures.Add(&specs[i])
	}

	Spec = func(name string) *StructTypeSpec {
		return specFixtures.Spec(name)
	}
}
