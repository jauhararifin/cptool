package core

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/jauhararifin/cptool/internal/logger"
)

// CompilationResult store the result of compiling solution. The Skipped property indicates whether the compilation
// process is skipped. The compilation can be skipped when the most up to date solution is already compiled.
// TargetPath property contain the path to compiled program. Duration property indicates the duration of compilation
// process.
type CompilationResult struct {
	Skipped      bool
	TargetPath   string
	Duration     time.Duration
	ErrorMessage string
}

// ErrLanguageNotDebuggable indicates that the language is not debuggable. This happens at compilation process when
// the language definition doesn't have debugcompile script to compile the solution with debug mode.
var ErrLanguageNotDebuggable = errors.New("Language is not debuggable")

// Compile will compile solution if not yet compiled. The compilation prosess will execute compile script of the
// language. It will use debugcompile script when debug parameter is true. When debug is true, but the language is
// not debuggable (doesn't contain debugcompile script), an ErrLanguageNotDebuggable error will returned. This function
// will execute the compilation script (could be compile/debugcompile) that defined in language definition. This execution
// could be skipped when the solution already compiled before.
func (cptool *CPTool) Compile(ctx context.Context, solution Solution, debug bool) (CompilationResult, error) {
	language := solution.Language
	if debug && !language.Debuggable {
		return CompilationResult{}, ErrLanguageNotDebuggable
	}

	targetDir := cptool.getCompiledDirectory(solution, debug)
	cptool.fs.MkdirAll(targetDir, os.ModePerm)

	targetPath := cptool.getCompiledTarget(solution, debug)
	if cptool.logger != nil {
		cptool.logger.Println(logger.VERBOSE, "Compiling to: ", targetPath)
	}

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
	if cptool.logger != nil {
		cptool.logger.Println(logger.VERBOSE, "Compiling using script: ", commandPath)
	}

	cmd := cptool.exec.CommandContext(ctx, commandPath, solution.Path, targetPath)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return CompilationResult{}, err
	}

	err = cmd.Start()
	if err != nil {
		return CompilationResult{}, err
	}

	compilationError, err := ioutil.ReadAll(stderr)
	if err != nil {
		return CompilationResult{}, err
	}

	err = cmd.Wait()
	if err != nil {
		if cptool.logger != nil {
			cptool.logger.Print(logger.VERBOSE, "Compilation script execution giving error result")
		}
		return CompilationResult{ErrorMessage: string(compilationError)}, err
	}

	return CompilationResult{
		Skipped:    false,
		TargetPath: targetPath,
	}, nil
}

// CompileByName will compile solution if not yet compiled. This method will search the language and solution by its name
// and then call Compile method. This method will return an error if the language or solution with it's name doesn't exist.
func (cptool *CPTool) CompileByName(ctx context.Context, languageName string, solutionName string, debug bool) (CompilationResult, error) {
	start := time.Now()

	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return CompilationResult{}, err
	}
	if cptool.logger != nil {
		cptool.logger.Println(logger.VERBOSE, "Compiling using language:", language.Name)
	}

	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return CompilationResult{}, err
	}
	if cptool.logger != nil {
		cptool.logger.Println(logger.VERBOSE, "Compiling solution:", solution.Name)
	}

	result, err := cptool.Compile(ctx, solution, debug)
	if err != nil {
		return result, err
	}
	result.Duration = time.Since(start)
	return result, nil
}

// GetCompilationRootDir returns directory of all compiled solutions.
func (cptool *CPTool) GetCompilationRootDir() string {
	return path.Join(cptool.workingDirectory, ".cptool/solutions")
}

func (cptool *CPTool) getCompiledDirectory(solution Solution, debug bool) string {
	language := solution.Language
	return path.Join(cptool.GetCompilationRootDir(), solution.Name, language.Name)
}

func (cptool *CPTool) getCompiledTarget(solution Solution, debug bool) string {
	dir := cptool.getCompiledDirectory(solution, debug)
	if debug {
		return path.Join(dir, "program_debug")
	}
	return path.Join(dir, "program")
}
