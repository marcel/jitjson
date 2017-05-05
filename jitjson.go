package main

import (
	"fmt"
	goast "go/ast"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/alecthomas/kingpin"
	jitast "github.com/marcel/jitjson/ast"
	"github.com/marcel/jitjson/codegen"
)

var (
	cli = kingpin.New("jitjson", "Finds structs with json tags and generates efficient (non-reflection) based JSON encoders")

	// TODO Add flag to configure the buffer size for the buffer pool
	gen        = cli.Command("gen", "Generate json encoders")
	genRootDir = gen.Arg("root-dir", "Directory to start searching for structs from. Defaults to '.'").Default(".").String()

	list        = cli.Command("list", "List eligible structs that were found")
	listRootDir = list.Arg("root-dir", "Directory to start searching for structs from. Defaults to '.'").Default(".").String()
	listFull    = list.Flag("full", "List full information for each struct including its fields").Short('f').Bool()

	clean        = cli.Command("clean", "Delete all auto-generated source files")
	cleanRootDir = clean.Arg("root-dir", "Directory to start searching for structs from. Defaults to '.'").Default(".").String()

	dump        = cli.Command("dump", "Dump auto-generated source code")
	dumpRootDir = dump.Arg("root-dir", "Directory to start searching for structs from. Defaults to '.'").Default(".").String()
	dumpFilter  = dump.Flag("filter", "Only dump code from packages matching filter").Short('f').String()

	files        = cli.Command("files", "List json encoder files that have been generated")
	filesRootDir = files.Arg("root-dir", "Directory to start searching for structs from. Defaults to '.'").Default(".").String()
)

// TODO
// option to include certain structs by name even if they don't have json tags

func main() {
	var err error

	switch kingpin.MustParse(cli.Parse(os.Args[1:])) {
	case gen.FullCommand():
		err = new(GenCommand).Run(FinderFrom(*genRootDir))
	case list.FullCommand():
		listCmd := &ListCommand{Full: *listFull}
		err = listCmd.Run(FinderFrom(*listRootDir))
	case clean.FullCommand():
		err = new(CleanCommand).Run(FinderFrom(*cleanRootDir))
	case dump.FullCommand():
		dumpCmd := &DumpCommand{Filter: *dumpFilter}
		err = dumpCmd.Run(FinderFrom(*dumpRootDir))
	case files.FullCommand():
		err = new(FilesCommand).Run(FinderFrom(*filesRootDir))
	}

	if err != nil {
		log.Fatal(err)
	}
}

func FinderFrom(rootDir string) *jitast.JSONStructFinder {
	finder := jitast.NewJSONStructFinder()

	err := finder.FindInDir(rootDir)
	if err != nil {
		log.Fatal(err)
	}

	return finder
}

type Command interface {
	Run(*jitast.JSONStructFinder) error
}

type GenCommand struct{}

func (c *GenCommand) Run(finder *jitast.JSONStructFinder) error {
	for _, structDir := range finder.StructDirectories() {
		metaCodeGen := codegen.NewMetaJSONEncoders(structDir)

		err := metaCodeGen.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}

type ListCommand struct {
	Full bool
}

func (c *ListCommand) Run(finder *jitast.JSONStructFinder) error {
	for _, dirs := range finder.StructDirectories() {
		fmt.Println(filepath.Join(dirs.PackageRoot, dirs.Package))
		for _, spec := range dirs.Specs {
			fmt.Printf("\t%s.%s\n", dirs.Package, spec.Name())
			if c.Full {
				structType, _ := spec.Type.(*goast.StructType)
				for _, field := range structType.Fields.List {
					if len(field.Names) == 0 {
						continue
					}

					name := field.Names[0].Name
					firstRune, _ := utf8.DecodeRuneInString(name)

					if unicode.IsUpper(firstRune) {
						fmt.Printf("\t\t%s\n", name)
					}
				}
			}
		}
	}

	return nil
}

type CleanCommand struct{}

func (c *CleanCommand) Run(finder *jitast.JSONStructFinder) error {
	for _, structDir := range finder.StructDirectories() {
		metaGen := codegen.NewMetaJSONEncoders(structDir)
		err := metaGen.DeleteOutdatedEncoderFile()

		if err != nil {
			return err
		}
	}

	return nil
}

type DumpCommand struct {
	Filter string
}

func (c *DumpCommand) Run(finder *jitast.JSONStructFinder) error {
	for _, structDir := range finder.StructDirectories() {
		if c.Filter != "" {
			if !strings.Contains(structDir.Directory, c.Filter) {
				continue
			}
		}
		metaCodeGen := codegen.NewMetaJSONEncoders(structDir)
		metaCodeGen.WriteTo(os.Stdout)
	}

	return nil
}

type FilesCommand struct{}

func (c *FilesCommand) Run(finder *jitast.JSONStructFinder) error {
	for _, structDir := range finder.StructDirectories() {
		metaGen := codegen.NewMetaJSONEncoders(structDir)
		if _, err := os.Stat(metaGen.PathToTargetFile()); !os.IsNotExist(err) {
			currentDir, err := os.Getwd()
			if err != nil {
				return err
			}

			relativePath, err := filepath.Rel(currentDir, metaGen.PathToTargetFile())
			if err != nil {
				return err
			}

			fmt.Println(relativePath)
		}
	}

	return nil
}
