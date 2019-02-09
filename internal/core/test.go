package core

import (
	"context"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/jauhararifin/cptool/internal/logger"
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

// TestCaseResult stores the result of testing a single test case. This contains information about the result of a test, the duration,
// and the test case. This also contains error if there is an error when testing the test case. The result of testing a test case can
// be classified into three category: "TestCaseSkipped", "TestCaseFailed", "TestCaseSuccess". Skipped means that there is an error
// (maybe IO error or something), that made the test skipped. When skipped, the Err property will set to error that made the test skipped.
// Failed means that the test run successfully but the solution's output is differ with expected output. Success means that test run
// successfully and gives output as expected.
type TestCaseResult struct {
	Testcase TestCase
	Duration time.Duration
	Status   int
	Err      error
}

// TestResult store test results of several test case. When testing using many testcase, this struct will returned after all test have been
// done. TestCaseResults's contain result of every single test case that tested. Duration contains durations of testing all test cases.
// UnsuccessfullTestsCount contains the number of unsuccessfull test case
type TestResult struct {
	TestCaseResults         []TestCaseResult
	Duration                time.Duration
	UnsuccessfullTestsCount uint
}

// Test will run solution using some testcases. A test case is a pair of text file that defines input and expected output of a test case.
// A file named "example.in" and "example.out" in current working directory considered as a test case named "example". This method will
// tests the given solution using all test cases with Name attribute that stars with `testPrefix`.
func (cptool *CPTool) Test(
	ctx context.Context,
	solution Solution,
	testPrefix string,
) (TestResult, error) {
	testCases := cptool.getAllTestCaseWithPrefix(testPrefix)
	if cptool.logger != nil {
		for _, tc := range testCases {
			cptool.logger.Println(logger.VERBOSE, "Test case found:", tc.Name)
			cptool.logger.Println(logger.VERBOSE, "Using input file:", tc.InputPath)
			cptool.logger.Println(logger.VERBOSE, "Using output file:", tc.OutputPath)
		}
	}

	results := TestResult{}
	startTime := time.Now()
	for _, testCase := range testCases {
		if cptool.logger != nil {
			cptool.logger.Println(logger.VERBOSE, "Testing program using test case:", testCase.Name)
		}
		success, duration, err := cptool.runSingleTest(ctx, solution, testCase)
		if err != nil {
			if cptool.logger != nil {
				cptool.logger.Println(logger.VERBOSE, "Program execution return an error")
			}
			results.TestCaseResults = append(results.TestCaseResults, TestCaseResult{
				Testcase: testCase,
				Status:   TestCaseSkipped,
				Err:      err,
			})
			results.UnsuccessfullTestsCount++
		} else {
			status := TestCaseFailed
			if success {
				status = TestCaseSuccess
			}
			if cptool.logger != nil {
				if success {
					cptool.logger.Println(logger.VERBOSE, "Test case passed")
				} else {
					cptool.logger.Println(logger.VERBOSE, "Test case failed")
				}
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

// TestByName will test solution using some test cases. This method will search the language and solution by its name and then
// call Test method. This method will return an error if the language or solution with it's name doesn't exist.
func (cptool *CPTool) TestByName(
	ctx context.Context,
	languageName string,
	solutionName string,
	testPrefix string,
) (TestResult, error) {
	if cptool.logger != nil {
		cptool.logger.Println(logger.VERBOSE, "Testing using language:", languageName)
	}
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return TestResult{}, err
	}

	if cptool.logger != nil {
		cptool.logger.Println(logger.VERBOSE, "Testing using solution:", solutionName)
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return TestResult{}, err
	}

	return cptool.Test(ctx, solution, testPrefix)
}

// GetOutputRootDir returns directory of all tested solution's output.
func (cptool *CPTool) GetOutputRootDir() string {
	return path.Join(cptool.workingDirectory, ".cptool/outputs")
}

func (cptool *CPTool) getOutputTarget(solution Solution, testCase TestCase) string {
	return path.Join(cptool.GetOutputRootDir(), solution.Name, solution.Language.Name, testCase.Name)
}

func (cptool *CPTool) runSingleTest(ctx context.Context, solution Solution, testCase TestCase) (bool, time.Duration, error) {
	outputFilePath := cptool.getOutputTarget(solution, testCase)
	if err := cptool.fs.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm); err != nil {
		return false, 0, err
	}
	outputFile, err := cptool.fs.Create(outputFilePath)
	if outputFile != nil {
		defer outputFile.Close()
	}
	if err != nil {
		return false, 0, err
	}
	inputFile, err := cptool.fs.Open(testCase.InputPath)
	if inputFile != nil {
		defer inputFile.Close()
	}
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
	if expectedOutputFile != nil {
		defer expectedOutputFile.Close()
	}
	if err != nil {
		return false, 0, err
	}
	same, _ := equalfile.CompareReader(outputFile, expectedOutputFile)
	if !same {
		return false, duration, nil
	}
	return true, duration, nil
}
