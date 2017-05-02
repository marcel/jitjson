package main

import (
	"fmt"
	"os"

	"github.com/marcel/jitjson"
)

func main() {
	finder := jitjson.NewJSONStructFinder()

	err := finder.FindInDir(os.Args[1])
	if err != nil {
		panic(err)
	}

	for _, dirs := range finder.StructDirectories() {
		fmt.Println(dirs.Directory)
		for _, spec := range dirs.Specs {
			fmt.Printf("\t%s.%s\n", dirs.Package, spec.Name())
		}
	}
}
