package cmd

import (
	"github.com/spf13/cobra"
)

func initCleanCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "clean",
		Short:   "Clean cptool cache directories",
		Version: GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			cptool, logger := newDefaultCptool(cmd)
			err := cptool.CleanCacheDirectory()
			if err != nil {
				logger.PrintError(err)
			}
		},
	}
	return cmd
}
