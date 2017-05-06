package ast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

// type WalkFunc func(path string, info os.FileInfo, err error) error
type FileSystemWalker interface {
	Walk(string, filepath.WalkFunc) error
}

type filePathWalker struct{}

func (f filePathWalker) Walk(root string, walkFunc filepath.WalkFunc) error {
	return filepath.Walk(root, walkFunc)
}

var defaultFileSystemWalker = filePathWalker{}

// JSONStructFinder recursively scans all directories starting at RootDir
// looking for go files and then traversing their AST in search of structs with
// json tags eligible for generating json encoders
type JSONStructFinder struct {
	// FileSystemWalker implements a Walk method which recursively traverses
	// directories looking for structs
	FileSystemWalker
	// WalkFunc is the callback which is passed every file discovered by the
	// FileSystemWalker
	filepath.WalkFunc

	structDirectories map[string]StructDirectory

	*token.FileSet

	currentDirectory string
}

type StructDirectory struct {
	ProjectRoot string
	PackageRoot string
	Package     string
	Directory   string
	ImportPath  string
	Specs       []StructTypeSpec
}

func NewJSONStructFinder() *JSONStructFinder {
	finder := new(JSONStructFinder)

	finder.WalkFunc = finder.findInFile
	finder.FileSystemWalker = defaultFileSystemWalker
	finder.FileSet = token.NewFileSet()

	finder.structDirectories = make(map[string]StructDirectory)

	return finder
}

var (
	errorPrefix              = "FindJSONStructFor:"
	ErrGoPathUndefined       = fmt.Errorf("%s $GOPATH is undefined", errorPrefix)
	ErrNonExistantSearchPath = func(path string) error {
		return fmt.Errorf("%s Search path '%s' does not exist", errorPrefix, path)
	}
	ErrNonExistantJSONStruct = func(importPath string, structName string) error {
		return fmt.Errorf("%s Could not find JSON struct spec for %s.%s{}", errorPrefix, importPath, structName)
	}
)

func FindJSONStructFor(importPath string, name string) (*StructTypeSpec, error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return nil, ErrGoPathUndefined
	}

	rootDir := filepath.Join(goPath, "src", importPath)
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		return nil, ErrNonExistantSearchPath(rootDir)
	}

	finder := NewJSONStructFinder()
	finder.FindInDir(rootDir)

	for _, structDir := range finder.StructDirectories() {
		if structDir.ImportPath == importPath {
			for _, typeSpec := range structDir.Specs {
				if typeSpec.Name() == name {
					return &typeSpec, nil
				}
			}
		}
	}

	return nil, ErrNonExistantJSONStruct(importPath, name)
}

func (s *JSONStructFinder) StructDirectories() []StructDirectory {
	structDirs := []StructDirectory{}

	for _, structDir := range s.structDirectories {
		structDirs = append(structDirs, structDir)
	}

	return structDirs
}

func (s *JSONStructFinder) StructTypeSpecs() []StructTypeSpec {
	specs := []StructTypeSpec{}

	for _, structDir := range s.StructDirectories() {
		specs = append(specs, structDir.Specs...)
	}

	return specs
}

func (s *JSONStructFinder) FindInDir(rootDir string) error {
	return s.FileSystemWalker.Walk(rootDir, s.WalkFunc)
}

func (s *JSONStructFinder) findInFile(directoryPath string, info os.FileInfo, err error) error {
	if info.IsDir() {
		if info.Name() == "vendor" {
			return filepath.SkipDir
		}

		return nil
	}

	if !strings.HasSuffix(info.Name(), ".go") || strings.HasSuffix(info.Name(), "_test.go") {
		return nil
	}

	absolutePath, err := filepath.Abs(directoryPath)
	if err != nil {
		return err
	}

	s.currentDirectory = filepath.Dir(absolutePath)

	file := filepath.Join(s.currentDirectory, info.Name())

	fileNode, err := parser.ParseFile(s.FileSet, file, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	for _, spec := range s.FindInAST(fileNode) {
		s.add(&spec)
	}

	return nil
}

func (s *JSONStructFinder) add(spec *StructTypeSpec) {
	structDir, found := s.structDirectories[spec.Directory]

	if !found {
		fullPath, _ := filepath.Abs(spec.Directory)
		pathParts := strings.Split(filepath.Dir(fullPath), "/src/")

		packageRoot := pathParts[len(pathParts)-1]

		dirName := filepath.Base(spec.Directory)
		importPath := filepath.Join(packageRoot, dirName)

		structDir = StructDirectory{
			PackageRoot: packageRoot,
			Directory:   spec.Directory,
			Package:     spec.PackageName,
			ImportPath:  importPath,
		}
	}

	structDir.Specs = append(structDir.Specs, *spec)

	s.structDirectories[spec.Directory] = structDir
}

func (s *JSONStructFinder) FindInAST(fileNode *ast.File) []StructTypeSpec {
	structs := []StructTypeSpec{}

	packageName := fileNode.Name.Name

	structFinder := func(node ast.Node) bool {
		typeSpec, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if structType, ok := typeSpec.Type.(*ast.StructType); ok {
			// We can't code-gen structs that aren't exported because we
			// have to run the code gen in a separate package (could in theory get
			// around this by temporarily defining a exported struct that wraps the
			// unexported one but...)
			firstRune, _ := utf8.DecodeRuneInString(typeSpec.Name.Name)
			if !unicode.IsUpper(firstRune) {
				return true
			}

			for _, field := range structType.Fields.List {
				if field.Tag != nil && field.Tag.Kind == token.STRING {
					if strings.Contains(field.Tag.Value, "json:") {
						spec := StructTypeSpec{
							Directory:   s.currentDirectory,
							PackageName: packageName,
							TypeSpec:    typeSpec,
						}

						structs = append(structs, spec)
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
