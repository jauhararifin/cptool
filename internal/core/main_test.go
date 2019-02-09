package core

import (
	"bytes"
	"os"
	"os/user"
	"testing"

	"github.com/jauhararifin/cptool/internal/executioner"
	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/spf13/afero"
)

func newTest() *CPTool {
	return &CPTool{
		languages: make(map[string]Language),

		exec: executioner.NewMemExec(),

		fs:                  afero.NewMemMapFs(),
		workingDirectory:    "/home/test/cptool",
		cptoolHomeDirectory: "/home/test/.cptool",
		homeDirectory:       "/home/test/",

		logger: logger.New(new(bytes.Buffer), 100),
	}
}

func newDefault() (*CPTool, error) {
	return New(nil, nil)
}

func getCptoolMemExec(cptool *CPTool) *executioner.MemExec {
	memExec, _ := cptool.exec.(*executioner.MemExec)
	return memExec
}

func TestNewDefault(t *testing.T) {
	cptool, err := newDefault()
	if err != nil {
		t.Error(err)
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
	cptool, _ := newDefault()
	err := cptool.Bootstrap()
	if err != nil {
		t.Error(err)
	}
}
