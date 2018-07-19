package core

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/jauhararifin/cptool/internal/logger"
)

// ExecutionResult stores execution result
type ExecutionResult struct {
	CompilationResult
	Duration time.Duration
}

// Run will run solution. This will compile the solution first if its not compiled yet
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

	isTerminal := true
	if stat, err := os.Stdin.Stat(); err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		isTerminal = false
	}

	cmd := cptool.exec.CommandContext(ctx, language.RunScript, targetPath)
	cmd.SetStdin(stdin)
	cmd.SetStdout(stdout)
	cmd.SetStderr(stderr)

	if isTerminal {
		logger.PrintInfo("Program started")
	}
	start := time.Now()
	err = cmd.Run()
	duration := time.Since(start)

	if err != nil {
		return ExecutionResult{}, err
	}
	return ExecutionResult{
		CompilationResult: compilationResult,
		Duration:          duration,
	}, err
}

// RunByName will run solution. This will compile the solution first if its not compiled yet
func (cptool *CPTool) RunByName(
	ctx context.Context,
	languageName string,
	solutionName string,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
) (ExecutionResult, error) {
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return ExecutionResult{}, err
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return ExecutionResult{}, err
	}

	return cptool.Run(ctx, solution, stdin, stdout, stderr)
}
