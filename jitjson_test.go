package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/marcel/jitjson/ast"
	"github.com/stretchr/testify/suite"
)

// type JitjsonCommandsTestSuite struct {
// 	suite.Suite
// }

// func TestJitjsonCommandsTestSuite(t *testing.T) {
// 	suite.Run(t, new(JitjsonCommandsTestSuite))
// }

// func (s *JitjsonCommandsTestSuite) SetupTest() {

// }

// func (s *JitjsonCommandsTestSuite) TestJitjsonCommands() {
// }

type JitjsonIntegrationTestSuite struct {
	suite.Suite
	tmpDir string
	finder *ast.JSONStructFinder
}

func TestJitjsonIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(JitjsonIntegrationTestSuite))
}

func (s *JitjsonIntegrationTestSuite) SetupTest() {
	dir, err := ioutil.TempDir(".", "temptest")
	s.Nil(err)
	s.tmpDir = dir
	err = os.Link("fixtures/navigation/navigation.go", filepath.Join(s.tmpDir, "navigation.go"))
	s.Nil(err)
	s.finder = ast.NewJSONStructFinder()
	s.finder.FindInDir(s.tmpDir)
}

func (s *JitjsonIntegrationTestSuite) TearDownTest() {
	os.RemoveAll(s.tmpDir)
}

var expectedListOutput = `github.com/marcel/jitjson/navigation
	navigation.Route
	navigation.Leg
	navigation.Step
	navigation.Location
	navigation.Address
`

var expectedFullListOutput = `github.com/marcel/jitjson/navigation
	navigation.Route
		Summary
		Legs
	navigation.Leg
		Distance
		Duration
		StartAddress
		EndAddress
		StartLocation
		EndLocation
		Steps
	navigation.Step
		Distance
		Duration
		StartLocation
		EndLocation
		Instructions
	navigation.Location
		Lat
		Lng
	navigation.Address
		Number
		Street
		City
		State
		ZipCode
		Country
`

func (s *JitjsonIntegrationTestSuite) runCmd(cmd Command) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := CommandRunner{cmd}.Run(s.finder, buf)
	return buf, err
}

func (s *JitjsonIntegrationTestSuite) TestList() {
	out, err := s.runCmd(new(ListCommand))
	s.Nil(err)
	s.Equal(expectedListOutput, out.String())

	out, err = s.runCmd(&ListCommand{Full: true})
	s.Nil(err)
	s.Equal(expectedFullListOutput, out.String())
}

func (s *JitjsonIntegrationTestSuite) TestGenCleanFiles() {
	// Nothing has happend yet so no files expected
	out, err := s.runCmd(new(FilesCommand))
	s.Nil(err)
	s.Equal("", out.String())

	// No files yet but clean just exists if it doesn't find the files
	out, err = s.runCmd(new(CleanCommand))
	s.Nil(err)
	s.Equal("", out.String())

	// Files get generated now
	out, err = s.runCmd(new(GenCommand))
	s.Nil(err)
	s.Equal("", out.String())

	// So files get listed
	out, err = s.runCmd(new(FilesCommand))
	s.Nil(err)
	s.Equal(filepath.Join(s.tmpDir, "json_encoders.go\n"), out.String())

	// Files get removed now
	out, err = s.runCmd(new(CleanCommand))
	s.Nil(err)
	s.Equal("", out.String())

	// So they are no longer listed
	out, err = s.runCmd(new(FilesCommand))
	s.Nil(err)
	s.Equal("", out.String())
}

func (s *JitjsonIntegrationTestSuite) Run(args []string) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	err := Run(args, out)

	return out, err
}

func (s *JitjsonIntegrationTestSuite) TestRun() {
	out, err := s.Run([]string{"list", s.tmpDir})
	s.Nil(err)

	s.Equal(expectedListOutput, out.String())

	out, err = s.Run([]string{"clean", s.tmpDir})
	s.Nil(err)

	out, err = s.Run([]string{"gen", s.tmpDir})
	s.Nil(err)

	out, err = s.Run([]string{"files", s.tmpDir})
	s.Nil(err)
	s.Equal(filepath.Join(s.tmpDir, "json_encoders.go\n"), out.String())

	out, err = s.Run([]string{"clean", s.tmpDir})
	s.Nil(err)

	out, err = s.Run([]string{"files", s.tmpDir})
	s.Nil(err)
	s.Equal("", out.String())

	out, err = s.Run([]string{"dump", s.tmpDir})
	s.Nil(err)
	s.Contains(out.String(), "package main")
}
