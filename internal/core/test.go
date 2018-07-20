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
		tcResult := TestCaseResult{Testcase: testCase}

		outputFilePath := cptool.getOutputTarget(solution, testCase)
		if err := cptool.fs.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm); err != nil {
			tcResult.Status = TestCaseSkipped
			tcResult.Err = err
			results.TestCaseResults = append(results.TestCaseResults, tcResult)
			results.UnsuccessfullTestsCount++
			continue
		}
		outputFile, err := cptool.fs.Create(outputFilePath)
		if err != nil {
			tcResult.Status = TestCaseSkipped
			tcResult.Err = err
			results.TestCaseResults = append(results.TestCaseResults, tcResult)
			results.UnsuccessfullTestsCount++
			continue
		}

		inputFile, err := cptool.fs.Open(testCase.InputPath)
		if err != nil {
			tcResult.Status = TestCaseSkipped
			tcResult.Err = err
			results.TestCaseResults = append(results.TestCaseResults, tcResult)
			results.UnsuccessfullTestsCount++
			continue
		}
		startTime := time.Now()
		_, err = cptool.Run(ctx, solution, inputFile, outputFile, os.Stderr)
		duration := time.Since(startTime)
		if err != nil {
			tcResult.Status = TestCaseSkipped
			tcResult.Err = err
			results.TestCaseResults = append(results.TestCaseResults, tcResult)
			results.UnsuccessfullTestsCount++
			continue
		}

		expectedOutputFile, err := cptool.fs.Open(testCase.OutputPath)
		if err != nil {
			tcResult.Status = TestCaseSkipped
			tcResult.Err = err
			results.TestCaseResults = append(results.TestCaseResults, tcResult)
			results.UnsuccessfullTestsCount++
			continue
		}

		same, err := equalfile.CompareReader(outputFile, expectedOutputFile)
		if err != nil {
			tcResult.Status = TestCaseSkipped
			tcResult.Err = err
			results.TestCaseResults = append(results.TestCaseResults, tcResult)
			results.UnsuccessfullTestsCount++
			continue
		}
		tcResult.Duration = duration
		if !same {
			tcResult.Status = TestCaseFailed
			results.TestCaseResults = append(results.TestCaseResults, tcResult)
			results.UnsuccessfullTestsCount++
		} else {
			tcResult.Status = TestCaseSuccess
			results.TestCaseResults = append(results.TestCaseResults, tcResult)
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
