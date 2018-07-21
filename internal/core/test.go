package core

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/udhos/equalfile"
)

const (
	// TestCaseSkipped indicates the testcase is skipped
	TestCaseSkipped = iota

	// TestCaseFailed indicates the testcase is failed
	TestCaseFailed = iota

	// TestCaseSuccess indicates the testcase is success
	TestCaseSuccess = iota
)

// TestCaseResult stores the result of a test case
type TestCaseResult struct {
	Testcase TestCase
	Duration time.Duration
	Status   int
	Err      error
}

// TestResult store test results
type TestResult struct {
	TestCaseResults         []TestCaseResult
	Duration                time.Duration
	UnsuccessfullTestsCount uint
}

func (cptool *CPTool) getOutputTarget(solution Solution, testCase TestCase) string {
	return path.Join(cptool.workingDirectory, ".cptool/outputs", solution.Name, solution.Language.Name, testCase.Name)
}

func (cptool *CPTool) runSingleTest(ctx context.Context, solution Solution, testCase TestCase) (bool, time.Duration, error) {
	outputFilePath := cptool.getOutputTarget(solution, testCase)
	if err := cptool.fs.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm); err != nil {
		return false, 0, err
	}
	outputFile, err := cptool.fs.Create(outputFilePath)
	defer outputFile.Close()
	if err != nil {
		return false, 0, err
	}
	inputFile, err := cptool.fs.Open(testCase.InputPath)
	defer inputFile.Close()
	if err != nil {
		return false, 0, err
	}
	startTime := time.Now()
	_, err = cptool.Run(ctx, solution, inputFile, outputFile, os.Stderr)
	duration := time.Since(startTime)
	if err != nil {
		return false, 0, err
	}
	_, err = outputFile.Seek(0, 0)
	if err != nil {
		return false, 0, err
	}
	expectedOutputFile, err := cptool.fs.Open(testCase.OutputPath)
	defer expectedOutputFile.Close()
	if err != nil {
		return false, 0, err
	}
	same, _ := equalfile.CompareReader(outputFile, expectedOutputFile)
	if !same {
		return false, duration, nil
	}
	return true, duration, nil
}

// Test will run solution using some testcases.
func (cptool *CPTool) Test(
	ctx context.Context,
	language Language,
	solution Solution,
	testPrefix string,
) (TestResult, error) {
	testCases := cptool.getAllTestCaseWithPrefix(testPrefix)
	results := TestResult{}
	startTime := time.Now()
	for _, testCase := range testCases {
		succes, duration, err := cptool.runSingleTest(ctx, solution, testCase)
		if err != nil {
			results.TestCaseResults = append(results.TestCaseResults, TestCaseResult{
				Testcase: testCase,
				Status:   TestCaseSkipped,
				Err:      err,
			})
			results.UnsuccessfullTestsCount++
		} else {
			status := TestCaseFailed
			if succes {
				status = TestCaseSuccess
			}
			results.TestCaseResults = append(results.TestCaseResults, TestCaseResult{
				Testcase: testCase,
				Status:   status,
				Duration: duration,
			})
		}
	}
	results.Duration = time.Since(startTime)
	return results, nil
}

// TestByName will run solution using some testcases.
func (cptool *CPTool) TestByName(
	ctx context.Context,
	languageName string,
	solutionName string,
	testPrefix string,
) (TestResult, error) {
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return TestResult{}, err
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return TestResult{}, err
	}

	return cptool.Test(ctx, language, solution, testPrefix)
}
