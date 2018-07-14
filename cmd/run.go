package cmd

import (
	"github.com/jauhararifin/cptool/internal/core"
	"github.com/spf13/cobra"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Run competitive programming solution",
	Long: "Run competitive programming solution. The program will compiled first if not yet compiled. The program\n" +
		"will be killed if still running after some period of time, you can change this behaviour using --timeout\n" +
		"option",
	Version: core.GetVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
