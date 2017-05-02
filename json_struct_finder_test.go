package jitjson

import (
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

	finder.FindInDir("fixtures")

	s.Equal(2, len(finder.StructDirectories()))
	s.Equal(8, len(finder.StructTypeSpecs()))
}
