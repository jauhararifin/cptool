package cmd

import (
	"fmt"

	"github.com/jauhararifin/cptool/internal/core"
	"github.com/spf13/cobra"
)

var langCommand = &cobra.Command{
	Use:     "lang",
	Short:   "List all available languages",
	Version: core.GetVersion(),
	Run: func(cmd *cobra.Command, args []string) {
		languages, _ := core.GetAllLanguages()
		for _, lang := range languages {
			fmt.Printf("[ %s ]\n", lang.Name)
			fmt.Printf("  language name:  %s\n", lang.VerboseName)
			fmt.Printf("  file extension: %s\n", lang.Extension)
			fmt.Printf("  compile script: %s\n", lang.CompileScript)
			fmt.Printf("  run script:     %s\n", lang.RunScript)
			if lang.Debuggable {
				fmt.Printf("  debug script:   %s\n", lang.DebugScript)
			}
			fmt.Println()
		}
	},
}
