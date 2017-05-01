package jitjson

import (
	"go/ast"
)

type StructTypeSpec struct {
	PackageName string
	*ast.TypeSpec
}

func NewStructTypeSpec(packageName string, spec *ast.TypeSpec) *StructTypeSpec {
	return &StructTypeSpec{PackageName: packageName, TypeSpec: spec}
}

func (s *StructTypeSpec) Name() string {
	return s.TypeSpec.Name.Name
}
