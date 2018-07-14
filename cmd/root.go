package cmd

import (
	"fmt"
	"os"

	"github.com/jauhararifin/cptool/internal/version"
	"github.com/spf13/cobra"
)

var mainCommand = &cobra.Command{
	Use:   "cptool",
	Short: "Simple tool that help you compile and run your competitive programming solution",
	Long: "Simple and easy to use tool for compile and run your competitive programming\n" +
		"solution built in Go. Check github.com/jauhararifin/cptool for more information",
	Version: version.GetVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	mainCommand.AddCommand(compileCommand)
	mainCommand.AddCommand(runCommand)
	mainCommand.AddCommand(testCommand)
}

// Execute cobra command line interface
func Execute() {
	if err := mainCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
