package cmd

import (
	"fmt"
	"os"
)

// MajorVersion incidates current cptool major version
const MajorVersion = 1

// MinorVersion incidates current cptool minor version
const MinorVersion = 0

// PatchVersion incidates current cptool patch version
const PatchVersion = 9

// GetVersion returns current version as string. Example: "v.1.0.1"
func GetVersion() string {
	return fmt.Sprintf("v%d.%d.%d", MajorVersion, MinorVersion, PatchVersion)
}

// Execute cobra command line interface
func Execute() {
	rootCommand := initRootCommand()
	rootCommand.AddCommand(initCompileCommand())
	rootCommand.AddCommand(initRunCommand())
	rootCommand.AddCommand(initTestCommand())
	rootCommand.AddCommand(initLangCommand())
	rootCommand.AddCommand(initCleanCommand())

	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
