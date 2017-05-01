package jitjson

import (
	"go/ast"
)

type StructTypeSpec struct {
	*ast.TypeSpec
}

func NewStructTypeSpec(spec *ast.TypeSpec) *StructTypeSpec {
	return &StructTypeSpec{spec}
}

func (s *StructTypeSpec) Name() string {
	return s.TypeSpec.Name.Name
}
