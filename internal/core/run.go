package core

import (
	"io"

	"github.com/jauhararifin/cptool/internal/logger"
)

// Run will run solution. This will compile the solution first if its not compiled yet
func (cptool *CPTool) Run(language Language, solution Solution, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	targetPath := cptool.getCompiledTarget(language, solution, false)
	err := cptool.Compile(language, solution, false)
	if err != nil {
		return err
	}

	cmd := cptool.exec.Command(language.RunScript, targetPath)
	cmd.SetStdin(stdin)
	cmd.SetStdout(stdout)
	cmd.SetStderr(stderr)
	err = cmd.Run()
	if err != nil {
		logger.PrintError("program exited with error")
	} else {
		logger.PrintSuccess("program executed with no error")
	}
	return err
}

// RunByName will run solution. This will compile the solution first if its not compiled yet
func (cptool *CPTool) RunByName(languageName string, solutionName string, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return err
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return err
	}

	return cptool.Run(language, solution, stdin, stdout, stderr)
}
