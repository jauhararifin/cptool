package cmd

import (
	"github.com/spf13/cobra"
)

func initCompileCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "compile",
		Short:   "Compile competitive programming solution",
		Version: cptool.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}
