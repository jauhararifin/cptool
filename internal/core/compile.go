package core

import (
	"context"
	"errors"
	"os"
	"path"
	"time"
)

// CompilationResult store compilation result
type CompilationResult struct {
	Skipped    bool
	TargetPath string
	Duration   time.Duration
}

// ErrLanguageNotDebuggable indicates that the language is not debuggable
var ErrLanguageNotDebuggable = errors.New("Language is not debuggable")

// GetCompilationRootDir returns root directory for compilation
func (cptool *CPTool) getCompilationRootDir() string {
	return path.Join(cptool.workingDirectory, ".cptool/solutions")
}

// GetCompiledDirectory returns directory path where compiled program exists
func (cptool *CPTool) getCompiledDirectory(solution Solution, debug bool) string {
	language := solution.Language
	return path.Join(cptool.getCompilationRootDir(), solution.Name, language.Name)
}

// GetCompiledTarget returns file path to compiled program
func (cptool *CPTool) getCompiledTarget(solution Solution, debug bool) string {
	dir := cptool.getCompiledDirectory(solution, debug)
	if debug {
		return path.Join(dir, "program_debug")
	}
	return path.Join(dir, "program")
}

// Compile will compile solution if not yet compiled
func (cptool *CPTool) Compile(ctx context.Context, solution Solution, debug bool) (CompilationResult, error) {
	language := solution.Language
	if debug && !language.Debuggable {
		return CompilationResult{}, ErrLanguageNotDebuggable
	}

	targetDir := cptool.getCompiledDirectory(solution, debug)
	cptool.fs.MkdirAll(targetDir, os.ModePerm)

	targetPath := cptool.getCompiledTarget(solution, debug)
	info, err := cptool.fs.Stat(targetPath)
	if err == nil {
		compiledTime := info.ModTime()
		if compiledTime.After(solution.LastUpdated) {
			return CompilationResult{
				Skipped:    true,
				TargetPath: targetPath,
			}, nil
		}
	}

	commandPath := language.CompileScript
	if debug {
		commandPath = language.DebugScript
	}
	cmd := cptool.exec.CommandContext(ctx, commandPath, solution.Path, targetPath)
	err = cmd.Run()
	if err != nil {
		return CompilationResult{}, err
	}
	return CompilationResult{
		Skipped:    false,
		TargetPath: targetPath,
	}, nil
}

// CompileByName will compile solution if not yet compiled
func (cptool *CPTool) CompileByName(ctx context.Context, languageName string, solutionName string, debug bool) (CompilationResult, error) {
	start := time.Now()
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return CompilationResult{}, err
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return CompilationResult{}, err
	}
	result, err := cptool.Compile(ctx, solution, debug)
	if err != nil {
		return CompilationResult{}, err
	}
	result.Duration = time.Since(start)
	return result, nil
}
