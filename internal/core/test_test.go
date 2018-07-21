package core

import (
	"context"
	"path"
	"testing"
	"time"

	"github.com/jauhararifin/cptool/internal/executioner"
)

func TestTest(t *testing.T) {
	cptool := newTest()
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.out"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.out"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "solution.lang"))
	solution := Solution{
		Name:        "solution",
		Language:    compileTestLanguage,
		Path:        path.Join(cptool.workingDirectory, "solution.lang"),
		LastUpdated: time.Now(),
	}

	_, err := cptool.Test(context.Background(), compileTestLanguage, solution, "test")
	if err != nil {
		t.Error(err)
	}
}

func prepareTestCase(cptool *CPTool, inputStr, expectedOutputStr, outputStr string) (Solution, TestCase) {
	solution := Solution{
		Name:        "solution",
		Language:    compileTestLanguage,
		Path:        "/solution",
		LastUpdated: time.Now(),
	}

	testCase := TestCase{
		Name:       "tc1",
		InputPath:  path.Join(cptool.workingDirectory, "tc1.in"),
		OutputPath: path.Join(cptool.workingDirectory, "tc1.out"),
	}

	input, _ := cptool.fs.Create(testCase.InputPath)
	input.WriteString(inputStr)
	input.Close()

	output, _ := cptool.fs.Create(testCase.OutputPath)
	output.WriteString(expectedOutputStr)
	output.Close()

	memexec := getCptoolMemExec(cptool)
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == solution.Language.CompileScript {
			return nil
		}
		_, err := m.Stdout.Write([]byte(outputStr))
		return err
	}

	return solution, testCase
}

func TestRunSingleTestCase(t *testing.T) {
	// ctx context.Context, solution Solution, testCase TestCase
	cptool := newTest()
	solution, testCase := prepareTestCase(cptool, "input", "expected_output", "expected_output")
	cptool.languages[solution.Language.Name] = solution.Language
	success, _, err := cptool.runSingleTest(context.Background(), solution, testCase)
	if err != nil {
		t.Error(err)
	}
	if !success {
		t.Error("RunSingleTestCase should returns success")
	}
}

func TestRunSingleTestCaseFailed(t *testing.T) {
	// ctx context.Context, solution Solution, testCase TestCase
	cptool := newTest()
	solution, testCase := prepareTestCase(cptool, "input", "expected_output", "some_different_output_with_exptected_output")
	cptool.languages[solution.Language.Name] = solution.Language
	success, _, err := cptool.runSingleTest(context.Background(), solution, testCase)
	if err != nil {
		t.Error(err)
	}
	if success {
		t.Error("RunSingleTestCase should returns failed")
	}
}
