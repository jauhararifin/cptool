package core

import "github.com/spf13/afero"

func newTest() *CPTool {
	return &CPTool{
		MajorVersion: 1,
		MinorVersion: 2,
		PatchVersion: 5,

		fs:               afero.NewMemMapFs(),
		workingDirectory: "/home/test/cptool",
		homeDirectory:    "/home/test/",
	}
}
