package cmd

import (
	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/spf13/cobra"
)

func initCleanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clean",
		Short:   "Clean cptool cache directories",
		Version: cptool.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			err := cptool.CleanCacheDirectory()
			if err != nil {
				logger.PrintError(err)
			}
		},
	}
	return cmd
}
