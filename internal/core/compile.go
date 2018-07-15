package core

import (
	"os"
	"os/exec"
	"path"

	"github.com/jauhararifin/cptool/internal/logger"
)

// GetCompilationRootDir returns root directory for compilation
func (cptool *CPTool) getCompilationRootDir() string {
	return ".cptool/solution"
}

// GetCompiledDirectory returns directory path where compiled program exists
func (cptool *CPTool) getCompiledDirectory(language Language, solution Solution, debug bool) string {
	return path.Join(cptool.getCompilationRootDir(), solution.Name, language.Name)
}

// GetCompiledTarget returns file path to compiled program
func (cptool *CPTool) getCompiledTarget(language Language, solution Solution, debug bool) string {
	dir := cptool.getCompiledDirectory(language, solution, debug)
	if debug {
		return path.Join(dir, "program_debug")
	}
	return path.Join(dir, "program")
}

// Compile will compile solution if not yet compiled
func (cptool *CPTool) Compile(language Language, solutionName string, debug bool) error {
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return err

	}
	logger.PrintInfo("compiling solution ", solution.Name)
	targetDir := cptool.getCompiledDirectory(language, solution, debug)
	cptool.fs.MkdirAll(targetDir, os.ModePerm)

	targetPath := cptool.getCompiledTarget(language, solution, debug)
	info, err := cptool.fs.Stat(targetPath)
	if err == nil {
		compiledTime := info.ModTime()
		if compiledTime.Before(solution.LastUpdated) {
			logger.PrintWarning("skipping compilation, solution already compiled")
			return nil
		}
	}

	cmd := exec.Command(language.CompileScript, solution.Path, targetPath)
	err = cmd.Run()
	if err != nil {
		logger.PrintError("compilation failed ", err)
	} else {
		logger.PrintSuccess("compilation success: ", targetPath)
	}
	return err
}
