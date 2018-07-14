package cmd

import (
	"github.com/jauhararifin/cptool/internal/version"
	"github.com/spf13/cobra"
)

var compileCommand = &cobra.Command{
	Use:     "compile",
	Short:   "Compile competitive programming solution",
	Version: version.GetVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
