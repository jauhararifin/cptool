package core

import (
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
