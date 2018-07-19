package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jauhararifin/cptool/internal/core"
	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/spf13/cobra"
)

func parseSolution(args []string) (string, core.Language) {
	if len(args) > 1 {
		solutionName := args[1]
		language, err := cptool.GetLanguageByName(args[0])
		if err != nil {
			logger.PrintError("cannot determine language")
			os.Exit(1)
		}
		return solutionName, language
	}
	solutionName := args[0]
	language, err := cptool.GetDefaultLanguage()
	if err != nil {
		logger.PrintError("cannot determine language")
		os.Exit(1)
	}
	return solutionName, language
}

func initCompileCommand() *cobra.Command {
	var debug bool

	cmd := &cobra.Command{
		Use:     "compile [LANGUAGE] SOLUTION",
		Short:   "Compile competitive programming solution",
		Version: cptool.GetVersion(),
		Args:    cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			solutionName, language := parseSolution(args)
			logger.PrintInfo("Compiling solution: ", solutionName)
			result, err := cptool.CompileByName(context.Background(), language.Name, solutionName, debug)
			if err != nil {
				logger.PrintError(err)
				os.Exit(1)
			}
			if result.Skipped {
				logger.PrintWarning("Compilation skipped because solution already compiled")
			}
			fmt.Printf("Compiled program : %s\n", result.TargetPath)
			fmt.Printf("Done in %.2f seconds\n", result.Duration.Seconds())
		},
	}

	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "compile your solution as debug mode")

	return cmd
}
