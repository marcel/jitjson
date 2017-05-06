package main

import (
	"fmt"
	goast "go/ast"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/alecthomas/kingpin"
	jitast "github.com/marcel/jitjson/ast"
	"github.com/marcel/jitjson/codegen"
)

func rootDirArg(cmdClause *kingpin.CmdClause) *string {
	return cmdClause.Arg(
		"root-dir",
		"Directory to start searching for structs from. Defaults to '.'",
	).Default(".").String()
}

var (
	cli = kingpin.New("jitjson", "Finds structs with json tags and generates efficient (non-reflection) based JSON encoders")

	// TODO Add flag to configure the buffer size for the buffer pool
	// TODO option to include certain structs by name even if they don't have json tags
	gen        = cli.Command("gen", "Generate json encoders")
	genRootDir = rootDirArg(gen)

	list        = cli.Command("list", "List eligible structs that were found")
	listRootDir = rootDirArg(list)
	listFull    = list.Flag("full", "List full information for each struct including its fields").Short('f').Bool()

	clean        = cli.Command("clean", "Delete all auto-generated source files")
	cleanRootDir = rootDirArg(clean)

	dump        = cli.Command("dump", "Dump auto-generated source code")
	dumpRootDir = rootDirArg(dump)
	dumpFilter  = dump.Flag("filter", "Only dump code from packages matching filter").Short('f').String()

	files        = cli.Command("files", "List json encoder files that have been generated")
	filesRootDir = rootDirArg(files)
)

func main() {
	err := Run(os.Args[1:], os.Stdout)

	if err != nil {
		panic(err)
	}
}

func Run(args []string, out io.Writer) error {
	var err error

	cmdLine := kingpin.MustParse(cli.Parse(args))
	finder := jitast.NewJSONStructFinder()

	err = finder.FindInDir(rootDir())
	if err != nil {
		return err
	}

	var command Command

	switch cmdLine {
	case gen.FullCommand():
		command = new(GenCommand)
	case list.FullCommand():
		command = &ListCommand{Full: *listFull}
	case clean.FullCommand():
		command = new(CleanCommand)
	case dump.FullCommand():
		command = &DumpCommand{Filter: *dumpFilter}
	case files.FullCommand():
		command = new(FilesCommand)
	}

	err = CommandRunner{command}.Run(finder, out)

	return err
}

func rootDir() string {
	for _, dir := range []*string{genRootDir, listRootDir, cleanRootDir, dumpRootDir, filesRootDir} {
		if *dir != "" {
			return *dir
		}
	}

	return ""
}

type CommandRunner struct {
	Command
}

type Command interface {
	Run(*jitast.JSONStructFinder, io.Writer) error
}

type GenCommand struct{}

func (c *GenCommand) Run(finder *jitast.JSONStructFinder, out io.Writer) error {
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

func (c *ListCommand) Run(finder *jitast.JSONStructFinder, out io.Writer) error {
	for _, dirs := range finder.StructDirectories() {
		fmt.Fprintln(out, filepath.Join(dirs.PackageRoot, dirs.Package))
		for _, spec := range dirs.Specs {
			fmt.Fprintf(out, "\t%s.%s\n", dirs.Package, spec.Name())
			if c.Full {
				structType, _ := spec.Type.(*goast.StructType)
				for _, field := range structType.Fields.List {
					if len(field.Names) == 0 {
						continue
					}

					name := field.Names[0].Name
					firstRune, _ := utf8.DecodeRuneInString(name)

					if unicode.IsUpper(firstRune) {
						fmt.Fprintf(out, "\t\t%s\n", name)
					}
				}
			}
		}
	}

	return nil
}

type CleanCommand struct{}

func (c *CleanCommand) Run(finder *jitast.JSONStructFinder, out io.Writer) error {
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

func (c *DumpCommand) Run(finder *jitast.JSONStructFinder, out io.Writer) error {
	for _, structDir := range finder.StructDirectories() {
		if c.Filter != "" {
			if !strings.Contains(structDir.Directory, c.Filter) {
				continue
			}
		}
		metaCodeGen := codegen.NewMetaJSONEncoders(structDir)
		metaCodeGen.WriteTo(out)
	}

	return nil
}

type FilesCommand struct{}

func (c *FilesCommand) Run(finder *jitast.JSONStructFinder, out io.Writer) error {
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

			fmt.Fprintln(out, relativePath)
		}
	}

	return nil
}
