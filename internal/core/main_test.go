package core

import (
	"os"
	"os/user"
	"testing"

	"github.com/jauhararifin/cptool/internal/executioner"
	"github.com/spf13/afero"
)

func newTest() *CPTool {
	return &CPTool{
		MajorVersion: 1,
		MinorVersion: 2,
		PatchVersion: 5,

		languages: make(map[string]Language),

		exec: executioner.NewMemExec(),

		fs:               afero.NewMemMapFs(),
		workingDirectory: "/home/test/cptool",
		homeDirectory:    "/home/test/",
	}
}

func getCptoolMemExec(cptool *CPTool) *executioner.MemExec {
	memExec, _ := cptool.exec.(*executioner.MemExec)
	return memExec
}

func TestNewDefault(t *testing.T) {
	cptool, err := NewDefault()
	if err != nil {
		t.Error(err)
	}

	if cptool.MajorVersion != MajorVersion {
		t.Error("NewDefault should returns cptool with MajorVersion=" + string(MajorVersion))
	}
	if cptool.MinorVersion != MinorVersion {
		t.Error("NewDefault should returns cptool with MinorVersion=" + string(MinorVersion))
	}
	if cptool.PatchVersion != PatchVersion {
		t.Error("NewDefault should returns cptool with PatchVersion=" + string(PatchVersion))
	}

	currentDirectory, _ := os.Getwd()
	if cptool.workingDirectory != currentDirectory {
		t.Error("NewDefault should returns cptool with workingDirectory points to current working directory")
	}

	currentUser, _ := user.Current()
	currentHome := currentUser.HomeDir
	if cptool.homeDirectory != currentHome {
		t.Error("NewDefault should returns cptool with homeDirectory points to current home directory")
	}
}

func TestBootstrap(t *testing.T) {
	cptool, _ := NewDefault()
	err := cptool.Bootstrap()
	if err != nil {
		t.Error(err)
	}
}
