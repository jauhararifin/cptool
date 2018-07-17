package cmd

import (
	"github.com/spf13/cobra"
)

func initTestCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test competitive programming solution",
		Long: "Test competitive programming solution. The program will compiled first if not yet compiled. The program will run\n" +
			"with provided testcases. The program will be killed if still running after some period of time, you can change\n" +
			"this behaviour using --timeout option.",
		Version: cptool.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}
