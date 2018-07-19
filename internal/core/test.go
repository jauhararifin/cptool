package core

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/udhos/equalfile"
)

func (cptool *CPTool) getOutputTarget(solution Solution, testCase TestCase) string {
	return path.Join(cptool.workingDirectory, ".cptool/outputs", solution.Name, solution.Language.Name, testCase.Name)
}

// Test will run solution using some testcases.
func (cptool *CPTool) Test(ctx context.Context, language Language, solution Solution, testPrefix string) error {
	testCases := cptool.getAllTestCaseWithPrefix(testPrefix)
	for _, testCase := range testCases {
		logger.PrintInfo("Running testcase ", testCase.Name)
		outputFilePath := cptool.getOutputTarget(solution, testCase)
		if err := cptool.fs.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm); err != nil {
			return err
		}
		outputFile, err := cptool.fs.Create(outputFilePath)
		if err != nil {
			logger.PrintError("Cannot open output file ", outputFilePath)
			return err
		}

		inputFile, err := cptool.fs.Open(testCase.InputPath)
		if err != nil {
			logger.PrintError("Cannot open input file ", testCase.InputPath)
			return err
		}
		err = cptool.Run(ctx, solution, inputFile, outputFile, os.Stderr)
		if err != nil {
			logger.PrintError("Testcase ", testCase.Name, " skipped, due to runtime error")
			return err
		}

		expectedOutputFile, err := cptool.fs.Open(testCase.OutputPath)
		if err != nil {
			logger.PrintError("Cannot open expected output file ", testCase.OutputPath)
			return err
		}

		same, err := equalfile.CompareReader(outputFile, expectedOutputFile)
		if err != nil {
			logger.PrintError("Cannot comparing output file with expected output: ", err)
			return err
		}
		if !same {
			logger.PrintError("Program's output differ with expected output")
		} else {
			logger.PrintSuccess("Program's output match with expected output")
		}
		fmt.Println()
	}
	return nil
}

// TestByName will run solution using some testcases.
func (cptool *CPTool) TestByName(ctx context.Context, languageName string, solutionName string, testPrefix string) error {
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return err
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return err
	}

	return cptool.Test(ctx, language, solution, testPrefix)
}
