package cmd

import (
	"os"

	"github.com/jauhararifin/cptool/internal/core"
	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/spf13/cobra"
)

func newDefaultCptool(cmd *cobra.Command) (*core.CPTool, *logger.Logger) {
	loggingLevel := logger.INFO
	if cmd != nil {
		if val, _ := cmd.PersistentFlags().GetBool("verbose"); val {
			loggingLevel = logger.VERBOSE
		}
	}
	cptoolLogger := logger.New(os.Stderr, loggingLevel)

	cptool, err := core.New(nil, cptoolLogger)
	if err != nil {
		cptoolLogger.PrintError(err)
		os.Exit(-1)
	}

	err = cptool.Bootstrap()
	if err != nil {
		cptoolLogger.PrintError(err)
		os.Exit(-1)
	}

	return cptool, cptoolLogger
}
