package jitjson

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"text/template"
	"time"
)

type MetaCodeGenerator struct {
	StructDirectory
	tempDirName string
	bytes.Buffer
}

func NewMetaCodeGenerator(structDir StructDirectory) *MetaCodeGenerator {
	generator := new(MetaCodeGenerator)

	generator.StructDirectory = structDir
	generator.tempDirName = generator.makeTempDirName()

	return generator
}

func (m *MetaCodeGenerator) WriteFile() error {
	file, err := os.Create(m.TempFile())
	if err != nil {
		return err
	}
	defer file.Close()

	return m.WriteTo(file)
}

func (m *MetaCodeGenerator) MakeTempDir() error {
	return os.MkdirAll(m.tempDirName, 0700)
}

func (m *MetaCodeGenerator) WriteTo(writer io.Writer) error {
	templ := template.Must(template.New("gen_encoders").Parse(tmpl))

	return templ.Execute(writer, m.StructDirectory)
}

func (m *MetaCodeGenerator) DeleteOutdatedEncoderFile() error {
	err := os.Remove(filepath.Join(m.Directory, CodeGeneratorTargetFile))

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
	return exec.Command("go", "run", m.TempFile()).Run()
}

func (m *MetaCodeGenerator) CleanUp() {
	// TODO Needs to be platform independent
	err := exec.Command("rm", "-rf", m.tempDirName).Run()
	if err != nil {
		log.Fatal(err)
	}
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
	codeGen.EncodingBufferStructWrapper()	

	{{- range .Specs }}
	codeGen.JSONMarshalerInterfaceFor("{{.Name}}")
	codeGen.EncoderMethodFor({{.PackageName}}.{{.Name}}{})
	{{- end  }}

	codeGen.WriteFile()
}
`
