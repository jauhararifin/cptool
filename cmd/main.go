package cmd

import (
	"fmt"
	"os"

	"github.com/jauhararifin/cptool/internal/core"
)

var cptool *core.CPTool
var err error

// Execute cobra command line interface
func Execute() {
	cptool, err = core.NewDefault()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cptool.Bootstrap()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCommand := initRootCommand()
	rootCommand.AddCommand(initCompileCommand())
	rootCommand.AddCommand(initRunCommand())
	rootCommand.AddCommand(initTestCommand())
	rootCommand.AddCommand(initLangCommand())

	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
