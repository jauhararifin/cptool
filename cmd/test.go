package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jauhararifin/cptool/internal/core"
	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/spf13/cobra"
)

func parseSolutionAndTestcasePrefix(args []string) (string, core.Language, string) {
	if len(args) > 2 {
		solutionName := args[1]
		testcasePrefix := args[2]
		language, err := cptool.GetLanguageByName(args[0])
		if err != nil {
			logger.PrintError("cannot determine language")
			os.Exit(1)
		}
		return solutionName, language, testcasePrefix
	}
	solutionName := args[0]
	testcasePrefix := args[1]
	language, err := cptool.GetDefaultLanguage()
	if err != nil {
		logger.PrintError("cannot determine language")
		os.Exit(1)
	}
	return solutionName, language, testcasePrefix
}

func initTestCommand() *cobra.Command {
	var hideTime bool
	var timeout time.Duration

	cmd := &cobra.Command{
		Use:   "test [LANGUAGE] SOLUTION TESTCASE_PREFIX",
		Short: "Test competitive programming solution",
		Long: "Test competitive programming solution. The program will compiled first if not yet compiled. The program will run\n" +
			"with provided testcases. The program will be killed if still running after some period of time, you can change\n" +
			"this behaviour using --timeout option.",
		Args:    cobra.RangeArgs(2, 3),
		Version: cptool.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			solutionName, language, testcasePrefix := parseSolutionAndTestcasePrefix(args)

			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			startTime := time.Now()
			cptool.TestByName(ctx, language.Name, solutionName, testcasePrefix)
			if ctx.Err() != nil {
				logger.PrintWarning("Test stopped due to timeout")
			}
			duration := time.Since(startTime)
			if !hideTime {
				fmt.Printf("Ellapsed time: %.2f seconds\n", duration.Seconds())
			}
		},
	}

	cmd.Flags().BoolVar(&hideTime, "hide-time", false, "hide the time indicator when execution finished")
	cmd.Flags().DurationVarP(&timeout, "timeout", "t", 10*time.Second, "Stop all test if still running after TIME\n"+
		"TIME is a floating point number with an optional suffix:\n"+
		"'s' for seconds (the default), 'm' for minutes, 'h' for hours or 'd' for days\n"+
		"The default value of this option is 100s. This option only works if the program is\n"+
		"running using stdin from file and not from terminal.\n")

	return cmd
}
