package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/spf13/cobra"
)

func initRunCommand() *cobra.Command {
	var timeout time.Duration
	var hideTime bool

	cmd := &cobra.Command{
		Use:   "run [LANGUAGE] [SOLUTION]",
		Short: "Run competitive programming solution",
		Long: "Run competitive programming solution. The program will compiled first if not yet compiled. The program\n" +
			"will be killed if still running after some period of time, you can change this behaviour using --timeout\n" +
			"option",
		Args:    cobra.RangeArgs(1, 2),
		Version: cptool.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			solutionName, language := parseSolution(args)

			isTerminal := true
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				isTerminal = false
			}

			if isTerminal && timeout != 10*time.Second {
				logger.PrintWarning("Timeout flag only works on piped stdin")
			}

			var ctx context.Context
			var cancel context.CancelFunc
			if isTerminal {
				ctx, cancel = context.WithCancel(context.Background())
			} else {
				ctx, cancel = context.WithTimeout(context.Background(), timeout)
			}
			defer cancel()

			startTime := time.Now()
			cptool.RunByName(ctx, language.Name, solutionName, os.Stdin, os.Stdout, os.Stderr)
			if ctx.Err() != nil {
				logger.PrintWarning("program stopped due to timeout")
			}
			duration := time.Since(startTime)
			if !hideTime {
				fmt.Printf("Ellapsed time: %.2f seconds\n", duration.Seconds())
			}
		},
	}

	cmd.Flags().BoolVar(&hideTime, "hide-time", false, "hide the time indicator when execution finished")
	cmd.Flags().DurationVarP(&timeout, "timeout", "t", 10*time.Second, "Kill program if still running after TIME\n"+
		"TIME is a floating point number with an optional suffix:\n"+
		"'s' for seconds (the default), 'm' for minutes, 'h' for hours or 'd' for days\n"+
		"The default value of this option is 10s. This option only works if the program is\n"+
		"running using stdin from file and not from terminal.\n")

	return cmd
}
