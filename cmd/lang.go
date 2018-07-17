package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func initLangCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "lang",
		Short:   "List all available languages",
		Version: cptool.GetVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			languages, _ := cptool.GetAllLanguages()

			if len(languages) == 0 {
				fmt.Println("No languages defined")
			}

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
}
