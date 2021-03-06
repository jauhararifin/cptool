package cmd

import (
	"github.com/spf13/cobra"
)

func initRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cptool",
		Short: "Simple tool that help you compile and run your competitive programming solution",
		Long: "Simple and easy to use tool for compile and run your competitive programming\n" +
			"solution built in Go. Check github.com/jauhararifin/cptool for more information",
		Version: GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.PersistentFlags().BoolP("verbose", "v", false, "Give verbose output")
	return cmd
}
