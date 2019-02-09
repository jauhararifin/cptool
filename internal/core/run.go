package core

import (
	"context"
	"io"
	"time"

	"github.com/jauhararifin/cptool/internal/logger"
)

// ExecutionResult stores execution result
type ExecutionResult struct {
	CompilationResult
	Duration time.Duration
}

// Run will run solution. This method will execute the solution using the run script that defined in language.
// Before the execution begin, this method will compile the solution first by calling Compile method. When there is
// no error occured, this method return ExecutionResult that contains CompilationResult and execution duration.
func (cptool *CPTool) Run(
	ctx context.Context,
	solution Solution,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
) (ExecutionResult, error) {
	language := solution.Language
	targetPath := cptool.getCompiledTarget(solution, false)
	compilationResult, err := cptool.Compile(ctx, solution, false)
	if err != nil {
		return ExecutionResult{}, err
	}

	cmd := cptool.exec.CommandContext(ctx, language.RunScript, targetPath)
	cmd.SetStdin(stdin)
	cmd.SetStdout(stdout)
	cmd.SetStderr(stderr)

	start := time.Now()
	err = cmd.Run()
	duration := time.Since(start)

	if err != nil {
		if cptool.logger != nil {
			cptool.logger.Println(logger.VERBOSE, "Program execution error:", err)
		}
		return ExecutionResult{}, err
	}
	return ExecutionResult{
		CompilationResult: compilationResult,
		Duration:          duration,
	}, err
}

// RunByName will run solution. This method will search the language and solution by its name and then call Run method.
// This method will return an error if the language or solution with it's name doesn't exist.
func (cptool *CPTool) RunByName(
	ctx context.Context,
	languageName string,
	solutionName string,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
) (ExecutionResult, error) {
	if cptool.logger != nil {
		cptool.logger.Println(logger.VERBOSE, "Run solution with language:", languageName)
	}
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return ExecutionResult{}, err
	}

	if cptool.logger != nil {
		cptool.logger.Println(logger.VERBOSE, "Run solution:", solutionName)
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return ExecutionResult{}, err
	}

	return cptool.Run(ctx, solution, stdin, stdout, stderr)
}
