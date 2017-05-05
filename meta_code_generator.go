package jitjson

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"text/template"
	"time"

	"github.com/marcel/jitjson/ast"
)

type FileSystemInterface interface {
	Create(name string) (File, error)
	MkdirAll(path string, perm os.FileMode) error
	Remove(name string) error
	RmRF(dirName string) error
	ExecGo(file string) error
}

type File interface {
	io.Closer
	io.Reader
	io.Writer
}

type fileSystem struct{}

func (f fileSystem) Create(name string) (File, error) {
	fd, err := os.Create(name)

	return fd, err
}

func (f fileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (f fileSystem) Remove(name string) error {
	return os.Remove(name)
}

func (f fileSystem) RmRF(dirName string) error {
	// TODO Needs to be platform independent
	return exec.Command("rm", "-rf", dirName).Run()
}

func (f fileSystem) ExecGo(file string) error {
	return exec.Command("go", "run", file).Run()
}

var DefaultFileSystemInterface = fileSystem{}

type MetaCodeGenerator struct {
	ast.StructDirectory
	fileSystem  FileSystemInterface
	tempDirName string
	bytes.Buffer
}

func NewMetaCodeGenerator(structDir ast.StructDirectory) *MetaCodeGenerator {
	generator := new(MetaCodeGenerator)

	generator.StructDirectory = structDir
	generator.fileSystem = DefaultFileSystemInterface
	generator.tempDirName = generator.makeTempDirName()

	return generator
}

func (m *MetaCodeGenerator) WriteFile() error {
	file, err := m.fileSystem.Create(m.TempFile())
	if err != nil {
		return err
	}
	defer file.Close()

	return m.WriteTo(file)
}

var tempDirFileMode os.FileMode = 0700

func (m *MetaCodeGenerator) MakeTempDir() error {
	return m.fileSystem.MkdirAll(m.tempDirName, tempDirFileMode)
}

func (m *MetaCodeGenerator) WriteTo(writer io.Writer) error {
	templ := template.Must(template.New("gen_encoders").Parse(tmpl))

	return templ.Execute(writer, m.StructDirectory)
}

func (m *MetaCodeGenerator) PathToTargetFile() string {
	return filepath.Join(m.Directory, CodeGeneratorTargetFile)
}

func (m *MetaCodeGenerator) DeleteOutdatedEncoderFile() error {
	err := m.fileSystem.Remove(m.PathToTargetFile())

	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		// File didn't exist, no need to propagate as error
		return nil
	}

	return err
}

func (m *MetaCodeGenerator) Exec() error {
	m.DeleteOutdatedEncoderFile()
	err := m.MakeTempDir()
	if err != nil {
		return err
	}

	defer m.CleanUp()

	err = m.WriteFile()
	if err != nil {
		return err
	}

	// TODO After generating the json_encoders.go file we should try a 'go build'
	// and if that returns with a non-zero exit status then the json_encoders
	// should be deleted and the error should be returned and displayed
	return m.fileSystem.ExecGo(m.TempFile())
}

func (m *MetaCodeGenerator) CleanUp() error {
	err := m.fileSystem.RmRF(m.tempDirName)
	if err != nil {
		return err
	}

	return nil
}

func (m *MetaCodeGenerator) makeTempDirName() string {
	tempDir := fmt.Sprintf("json_encoders_generator_%d", time.Now().UnixNano())

	return filepath.Join(m.Directory, tempDir)
}

func (m *MetaCodeGenerator) TempFile() string {
	return filepath.Join(m.tempDirName, "gen_encoders.go")
}

// TODO Needs to either write structEncoders into a file for the project (in
// order to have no external dependencies), or import it from jitjson
var tmpl = `package main

import (
	"github.com/marcel/jitjson"
	"{{.PackageRoot}}/{{.Package}}"
)

func main() {
	codeGen := jitjson.NewCodeGenerator("{{.Directory}}", "{{.Package}}")
	codeGen.PackageDeclaration()
	codeGen.ImportDeclaration()	
	codeGen.SetBufferPoolVar()
	codeGen.EncodingBufferStructWrapper()	

	{{- range .Specs }}
	codeGen.JSONMarshalerInterfaceFor("{{.Name}}")
	codeGen.EncoderMethodFor({{.PackageName}}.{{.Name}}{})
	{{- end  }}

	codeGen.WriteFile()
}
`
