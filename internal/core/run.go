package core

import "github.com/jauhararifin/cptool/internal/logger"

// Run will run solution. This will compile the solution first if its not compiled yet
func (cptool *CPTool) Run(language Language, solution Solution) error {
	targetPath := cptool.getCompiledTarget(language, solution, false)
	err := cptool.Compile(language, solution, false)
	if err != nil {
		return err
	}

	cmd := cptool.exec.Command(language.RunScript, targetPath)
	err = cmd.Run()
	if err != nil {
		logger.PrintError("program exited with error")
	} else {
		logger.PrintSuccess("program executed with no error")
	}
	return err
}

// RunByName will run solution. This will compile the solution first if its not compiled yet
func (cptool *CPTool) RunByName(languageName string, solutionName string) error {
	language, err := cptool.GetLanguageByName(languageName)
	if err != nil {
		return err
	}
	solution, err := cptool.GetSolution(solutionName, language)
	if err != nil {
		return err
	}

	return cptool.Run(language, solution)
}
