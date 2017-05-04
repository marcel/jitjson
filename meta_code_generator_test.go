package jitjson

import (
	"bytes"
	"errors"
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/suite"
)

type mockFileSystem struct {
	madeDirs []struct {
		path string
		perm os.FileMode
	}
	removed  []string
	rmRfed   []string
	goExeced []string

	errorToReturn     error
	fileErrorToReturn error
	filesCreated      []*mockFile
}

func NewMockFileSystem() *mockFileSystem {
	return new(mockFileSystem)
}

func (m *mockFileSystem) Create(name string) (File, error) {
	file := &mockFile{
		name:          name,
		errorToReturn: m.fileErrorToReturn,
	}

	m.filesCreated = append(m.filesCreated, file)

	return file, m.errorToReturn
}

func (m *mockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	m.madeDirs = append(m.madeDirs, struct {
		path string
		perm os.FileMode
	}{path, perm})

	return m.errorToReturn
}

func (m *mockFileSystem) Remove(name string) error {
	m.removed = append(m.removed, name)

	return m.errorToReturn
}

func (m *mockFileSystem) RmRF(dirName string) error {
	m.rmRfed = append(m.rmRfed, dirName)

	return m.errorToReturn
}

func (m *mockFileSystem) ExecGo(file string) error {
	m.goExeced = append(m.goExeced, file)

	return m.errorToReturn
}

type mockFile struct {
	name      string
	wasClosed bool
	bytes.Buffer
	errorToReturn error
}

func (m *mockFile) Close() error {
	m.wasClosed = true

	return m.errorToReturn
}

func (m *mockFile) Read(p []byte) (n int, err error) {
	bytesRead, _ := m.Buffer.Read(p)

	return bytesRead, m.errorToReturn
}

func (m *mockFile) Write(p []byte) (n int, err error) {
	bytesWritten, _ := m.Buffer.Write(p)
	return bytesWritten, m.errorToReturn
}

type MetaCodeGeneratorTestSuite struct {
	suite.Suite
	generator *MetaCodeGenerator
	fs        *mockFileSystem
}

func TestMetaCodeGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(MetaCodeGeneratorTestSuite))
}

func (s *MetaCodeGeneratorTestSuite) SetupTest() {
	structDir := StructDirectory{
		ProjectRoot: "/path/to/project/src/github.com/marcel/jitson",
		PackageRoot: "github.com/marcel/jitjson/fixtures",
		Package:     "media",
		Directory:   "/path/to/project/src/github.com/marcel/jitson/fixtures/media",
		Specs:       []StructTypeSpec{*Spec("Album")},
	}

	s.generator = NewMetaCodeGenerator(structDir)
	s.fs = NewMockFileSystem()
	s.generator.fileSystem = s.fs
}

func (s *MetaCodeGeneratorTestSuite) TestWriteFile() {
	s.Equal(0, len(s.fs.filesCreated))

	s.Nil(s.generator.WriteFile())

	s.Equal(1, len(s.fs.filesCreated))

	file := s.fs.filesCreated[0]
	s.True(file.wasClosed)

	expected := new(bytes.Buffer)
	s.generator.WriteTo(expected)

	s.Equal(expected.String(), file.String())
}

func (s *MetaCodeGeneratorTestSuite) TestWriteFileFails() {
	s.fs.errorToReturn = errors.New("Create failed")

	s.Equal(0, len(s.fs.filesCreated))

	s.NotNil(s.generator.WriteFile())

	s.Equal(1, len(s.fs.filesCreated))

	file := s.fs.filesCreated[0]
	s.False(file.wasClosed)

	s.Equal("", file.String())
}

func (s *MetaCodeGeneratorTestSuite) TestExec() {
	emptyFS := NewMockFileSystem()
	s.Equal(emptyFS, s.fs)

	s.Nil(s.generator.Exec())

	// Delete outdated encoder file
	s.Equal(1, len(s.fs.removed))
	outdatedEncoderFileRemoved := s.generator.PathToTargetFile()
	s.Equal(outdatedEncoderFileRemoved, s.fs.removed[0])

	// Make temp dir
	s.Equal(1, len(s.fs.madeDirs))
	tempDirMade := s.generator.tempDirName

	expectedMadeDir := struct {
		path string
		perm os.FileMode
	}{tempDirMade, tempDirFileMode}

	s.Equal(expectedMadeDir, s.fs.madeDirs[0])

	// Write file
	s.Equal(1, len(s.fs.filesCreated))

	file := s.fs.filesCreated[0]
	s.True(file.wasClosed)

	expectedFileData := new(bytes.Buffer)
	s.generator.WriteTo(expectedFileData)

	s.Equal(expectedFileData.String(), file.String())

	// Exec gened go file
	s.Equal(1, len(s.fs.goExeced))
	expectedGoFile := s.generator.TempFile()

	s.Equal(expectedGoFile, s.fs.goExeced[0])

	// Clean up
	s.Equal(1, len(s.fs.rmRfed))
	expectedRmRF := s.generator.tempDirName

	s.Equal(expectedRmRF, s.fs.rmRfed[0])
}

func (s *MetaCodeGeneratorTestSuite) TestDeleteOutdateEncoderFileThatDoesNotExist() {
	fileDoesNotExistError := new(os.PathError)
	fileDoesNotExistError.Err = syscall.ENOENT
	s.fs.errorToReturn = fileDoesNotExistError

	s.Nil(s.generator.DeleteOutdatedEncoderFile())

	otherPathError := new(os.PathError)
	otherPathError.Err = syscall.ENOBUFS
	s.fs.errorToReturn = otherPathError

	s.Equal(otherPathError, s.generator.DeleteOutdatedEncoderFile())
}

func (s *MetaCodeGeneratorTestSuite) TestMetaCodeGenerator() {
	buf := new(bytes.Buffer)

	s.Nil(s.generator.WriteTo(buf))

	expected :=
		`package main

import (
	"github.com/marcel/jitjson"
	"github.com/marcel/jitjson/fixtures/media"
)

func main() {
	codeGen := jitjson.NewCodeGenerator("/path/to/project/src/github.com/marcel/jitson/fixtures/media", "media")
	codeGen.PackageDeclaration()
	codeGen.ImportDeclaration()	
	codeGen.SetBufferPoolVar()
	codeGen.EncodingBufferStructWrapper()
	codeGen.JSONMarshalerInterfaceFor("Album")
	codeGen.EncoderMethodFor(media.Album{})

	codeGen.WriteFile()
}
`
	s.Equal(expected, buf.String())
}
