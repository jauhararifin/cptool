package core

import (
	"errors"
	"os"
	"path"

	"github.com/jauhararifin/cptool/internal/logger"
)

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
func (cptool *CPTool) Compile(solution Solution, debug bool) error {
	language := solution.Language
	if debug && !language.Debuggable {
		logger.PrintError("language is not debuggable")
		return ErrLanguageNotDebuggable
	}

	logger.PrintInfo("compiling solution: ", solution.Name)
	targetDir := cptool.getCompiledDirectory(solution, debug)
	cptool.fs.MkdirAll(targetDir, os.ModePerm)

	targetPath := cptool.getCompiledTarget(solution, debug)
	info, err := cptool.fs.Stat(targetPath)
	if err == nil {
		compiledTime := info.ModTime()
		if compiledTime.After(solution.LastUpdated) {
			logger.PrintWarning("skipping compilation, solution already compiled")
			return nil
		}
	}

	commandPath := language.CompileScript
	if debug {
		commandPath = language.DebugScript
	}
	cmd := cptool.exec.Command(commandPath, solution.Path, targetPath)
	err = cmd.Run()
	if err != nil {
		logger.PrintError("compilation failed: ", err)
	} else {
		logger.PrintSuccess("compilation success: ", targetPath)
	}
	return err
}

// CompileByName will compile solution if not yet compiled
func (cptool *CPTool) CompileByName(languageName string, solutionName string, debug bool) error {
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		logger.PrintError("compilation failed: ", err)
		return err
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		logger.PrintError("compilation failed: ", err)
		return err
	}

	return cptool.Compile(solution, debug)
}
