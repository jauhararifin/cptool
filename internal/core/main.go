package core

import (
	"os"
	"os/user"

	"github.com/jauhararifin/cptool/internal/executioner"
	"github.com/spf13/afero"
)

// MajorVersion incidates current cptool major version
const MajorVersion = 1

// MinorVersion incidates current cptool minor version
const MinorVersion = 0

// PatchVersion incidates current cptool patch version
const PatchVersion = 6

// CPTool stores information about this tool. The information includes version, all known languages, and configuration directories.
type CPTool struct {
	MajorVersion int
	MinorVersion int
	PatchVersion int

	languages map[string]Language

	exec executioner.Exec

	fs                  afero.Fs
	workingDirectory    string
	cptoolHomeDirectory string
	homeDirectory       string
}

// NewDefault create new default cptool instance. This instance contains information about the tool version,
// configuration direcories, and languages.
func NewDefault() (*CPTool, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	return &CPTool{
		MajorVersion: MajorVersion,
		MinorVersion: MinorVersion,
		PatchVersion: PatchVersion,

		languages: make(map[string]Language),

		exec: executioner.NewOSExec(),

		fs:                  afero.NewOsFs(),
		workingDirectory:    cwd,
		cptoolHomeDirectory: os.Getenv("CPTOOL_HOME"),
		homeDirectory:       user.HomeDir,
	}, nil
}

// Bootstrap will bootstrap cptool. The bootstrap process will load all language from known directories.
func (cptool *CPTool) Bootstrap() error {
	cptool.loadAllLanguages()
	return nil
}
