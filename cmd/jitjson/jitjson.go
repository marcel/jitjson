package main

import (
	"fmt"

	"github.com/marcel/jitjson"
	"github.com/marcel/jitjson/fixtures"
)

func main() {
	codeGen := new(jitjson.CodeGenerator)

	codeGen.PackageDelcaration()
	codeGen.ImportDeclaration()

	codeGen.EncoderMethodFor(fixtures.Album{})
	codeGen.EncoderMethodFor(fixtures.Artist{})
	codeGen.EncoderMethodFor(fixtures.Track{})

	fmt.Println(string(codeGen.Bytes()))
}
