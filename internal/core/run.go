package core

import (
	"io"

	"github.com/jauhararifin/cptool/internal/logger"
)

// Run will run solution. This will compile the solution first if its not compiled yet
func (cptool *CPTool) Run(solution Solution, stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	language := solution.Language
	targetPath := cptool.getCompiledTarget(solution, false)
	err := cptool.Compile(solution, false)
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

	return cptool.Run(solution, stdin, stdout, stderr)
}
