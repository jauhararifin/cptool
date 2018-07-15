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
const PatchVersion = 0

// CPTool represent this tool
type CPTool struct {
	MajorVersion int
	MinorVersion int
	PatchVersion int

	languages map[string]Language

	exec executioner.Exec

	fs               afero.Fs
	workingDirectory string
	homeDirectory    string
}

// NewDefault create new default cptool instance
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

		fs:               afero.NewOsFs(),
		workingDirectory: cwd,
		homeDirectory:    user.HomeDir,
	}, nil
}
