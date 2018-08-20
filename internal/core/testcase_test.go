package core

import (
	"path"
	"testing"
)

func TestGetTestCasesWithPrefix(t *testing.T) {
	cptool := newTest()
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.out"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.out"))

	testcases := cptool.getAllTestCaseWithPrefix("test")
	if len(testcases) != 2 {
		t.Error("getAllTestCaseWithPrefix should return 2 testcase")
	}

	tc1 := TestCase{
		Name:       "test.1",
		InputPath:  path.Join(cptool.workingDirectory, "test.1.in"),
		OutputPath: path.Join(cptool.workingDirectory, "test.1.out"),
	}
	if testcases[0] != tc1 {
		t.Error("getAllTestCaseWithPrefix should return ", tc1)
	}

	tc2 := TestCase{
		Name:       "test.2",
		InputPath:  path.Join(cptool.workingDirectory, "test.2.in"),
		OutputPath: path.Join(cptool.workingDirectory, "test.2.out"),
	}
	if testcases[1] != tc2 {
		t.Error("getAllTestCaseWithPrefix should return ", tc2)
	}
}

func TestGetTestCasesWithPrefixWithoutOutputFile(t *testing.T) {
	cptool := newTest()
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test/1.in"))

	testcases := cptool.getAllTestCaseWithPrefix("test")
	if len(testcases) > 0 {
		t.Error("getAllTestCaseWithPrefix should return nothing")
	}
}

func TestGetTestCasesByName(t *testing.T) {
	cptool := newTest()
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.out"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.out"))

	testcase, err := cptool.getTestCaseByName("test.1")
	if err != nil {
		t.Error("getTestCaseByName should not return error")
	}

	tc1 := TestCase{
		Name:       "test.1",
		InputPath:  path.Join(cptool.workingDirectory, "test.1.in"),
		OutputPath: path.Join(cptool.workingDirectory, "test.1.out"),
	}
	if testcase != tc1 {
		t.Error("getAllTestCaseWithPrefix should return ", tc1)
	}
}

func TestGetTestCasesByNameWithMissingTestCase(t *testing.T) {
	cptool := newTest()
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.out"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.out"))

	_, err := cptool.getTestCaseByName("test.3")
	if err == nil || err != ErrNoSuchTestCase {
		t.Error("getTestCaseByName should return ErrNoSuchTestCase error")
	}
}
