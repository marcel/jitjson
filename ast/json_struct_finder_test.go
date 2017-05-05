package ast

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

// type MockFileSystemWalker struct {
// 	fileInfos []MockFileInfo
// }

// type MockFileInfo struct {
// 	name  string
// 	isDir bool
// }

// func (m *MockFileInfo) Name() string {
// 	return m.name
// }

// func (m *MockFileInfo) Size() int64 {
// 	return 1024
// }

// func (m *MockFileInfo) Mode() os.FileMode {
// 	if m.isDir {
// 		return os.ModePerm | os.ModeDir
// 	}

// 	return os.ModePerm
// }

// func (m *MockFileInfo) ModTime() time.Time {
// 	return time.Now()
// }

// func (m *MockFileInfo) IsDir() bool {
// 	return m.isDir
// }

// func (m *MockFileInfo) Sys() interface{} {
// 	return nil
// }

// func (m *MockFileSystemWalker) Walk(root string, walkFunc filepath.WalkFunc) {
// 	for _, fileInfo := range m.fileInfos {
// 		dir := filepath.Dir(fileInfo.Name())
// 		file := filepath.Base(fileInfo.Name())
// 		walkFunc(dir, fileInfo, nil)
// 	}
// }

type JSONStructFinderTestSuite struct {
	suite.Suite
	gocode []byte
}

func TestJSONStructFinderTestSuite(t *testing.T) {
	suite.Run(t, new(JSONStructFinderTestSuite))
}

func (s *JSONStructFinderTestSuite) TestFindInDir() {
	finder := NewJSONStructFinder()
	s.Equal(0, len(finder.StructDirectories()))
	s.Equal(0, len(finder.StructTypeSpecs()))

	finder.FindInDir("../fixtures")

	s.Equal(2, len(finder.StructDirectories()))
	s.Equal(9, len(finder.StructTypeSpecs()))
}

func (s *JSONStructFinderTestSuite) TestFindJSONStructFor() {
	// Golden path: Struct exists
	importPath := "github.com/marcel/jitjson/fixtures/media"
	structName := "Album"
	jsonStruct, err := FindJSONStructFor(importPath, structName)
	s.Nil(err)
	s.Equal(structName, jsonStruct.Name())
	s.Equal("media", jsonStruct.PackageName)

	// Error: Import path does not exist
	jsonStruct, err = FindJSONStructFor(importPath+"/does/not/exist", structName)
	s.Nil(jsonStruct)
	// N.B. We just check error message substring  because we don't want to reimplement the
	// full path resolution in the code
	s.Contains(err.Error(), "Search path")

	// Error: Struct does not exist
	jsonStruct, err = FindJSONStructFor(importPath, structName+"DoesNotExist")
	s.Nil(jsonStruct)
	s.Equal(ErrNonExistantJSONStruct(importPath, structName+"DoesNotExist"), err)

	// Error: GOPATH is not defined
	originalGoPathVal := os.Getenv("GOPATH")
	defer os.Setenv("GOPATH", originalGoPathVal)

	os.Setenv("GOPATH", "")
	jsonStruct, err = FindJSONStructFor(importPath, structName)
	s.Nil(jsonStruct)
	s.Equal(ErrGoPathUndefined, err)
}
