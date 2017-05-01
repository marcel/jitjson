package jitjson

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type JSONStructFinder struct {
}

func (s JSONStructFinder) Structs(gocode []byte) []StructTypeSpec {
	fs := token.NewFileSet()
	fileNode, err := parser.ParseFile(fs, "file name goes here", gocode, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	structs := []StructTypeSpec{}

	packageName := fileNode.Name.Name

	structFinder := func(node ast.Node) bool {
		typeSpec, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if structType, ok := typeSpec.Type.(*ast.StructType); ok {
			for _, field := range structType.Fields.List {
				if field.Tag != nil && field.Tag.Kind == token.STRING {
					if strings.Contains(field.Tag.Value, "json:") {
						structs = append(structs, StructTypeSpec{packageName, typeSpec})
						return true
					}
				}
			}
		}

		return true
	}

	ast.Inspect(fileNode, structFinder)

	return structs
}
