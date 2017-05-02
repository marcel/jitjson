package main

import (
	"os"

	"github.com/marcel/jitjson"
)

func main() {
	finder := jitjson.NewJSONStructFinder()

	rootDir := os.Args[1]

	err := finder.FindInDir(rootDir)
	if err != nil {
		panic(err)
	}

	for _, structDir := range finder.StructDirectories() {
		metaCodeGen := jitjson.NewMetaCodeGenerator(structDir)
		metaCodeGen.WriteTo(os.Stdout)
	}
}
