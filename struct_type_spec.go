package jitjson

import (
	"go/ast"
)

type StructTypeSpec struct {
	PackageName string
	Directory   string
	*ast.TypeSpec
}

func (s *StructTypeSpec) Name() string {
	return s.TypeSpec.Name.Name
}
