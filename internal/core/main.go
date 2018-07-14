package core

import (
	"os"
	"os/user"

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

		fs:               afero.NewOsFs(),
		workingDirectory: cwd,
		homeDirectory:    user.HomeDir,
	}, nil
}
