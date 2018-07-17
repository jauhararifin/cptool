package core

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

// TestCase represent testcase
type TestCase struct {
	Name       string
	InputPath  string
	OutputPath string
}

// ErrNoSuchTestCase indicates testcase not found
var ErrNoSuchTestCase = errors.New("No such testcase")

func (cptool *CPTool) getTestCaseByName(testcaseName string) (TestCase, error) {
	for _, testCase := range cptool.getAllTestCaseWithPrefix(testcaseName) {
		if testCase.Name == testcaseName {
			return testCase, nil
		}
	}
	return TestCase{}, ErrNoSuchTestCase
}

func (cptool *CPTool) getAllTestCaseWithPrefix(testcasePrefix string) []TestCase {
	testCases := make([]TestCase, 0)
	cwd := filepath.Clean(cptool.workingDirectory)
	afero.Walk(cptool.fs, cwd, func(testPath string, info os.FileInfo, err error) error {
		relativePath := filepath.Clean(testPath)[len(cwd):]
		if len(relativePath) == 0 {
			return nil
		}
		if relativePath[0] == '/' {
			relativePath = relativePath[1:]
		}
		if !info.IsDir() && strings.HasPrefix(relativePath, testcasePrefix) && filepath.Ext(testPath) == ".in" {
			testName := relativePath[:len(relativePath)-3]
			outputFilePath := path.Join(cptool.workingDirectory, testName+".out")
			info, err := cptool.fs.Stat(outputFilePath)
			if err != nil || info.IsDir() {
				return nil
			}
			testCases = append(testCases, TestCase{
				Name:       testName,
				InputPath:  testPath,
				OutputPath: outputFilePath,
			})
		}
		return nil
	})
	return testCases
}