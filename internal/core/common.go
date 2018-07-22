package core

import (
	"path"
)

// GetConfigurationPaths returns all paths to the directory that considered contain cptool configuration.
// This directory contains language definition, and some configuration like default language. The paths that
// considered are "/etc/cptool", "~/.cptool", your $CPTOOL_HOME directory and .cptool directory in your
// current working directory.
func (cptool *CPTool) GetConfigurationPaths() []string {
	paths := make([]string, 0)
	paths = appendPaths(paths, path.Join(cptool.workingDirectory, ".cptool"))
	paths = appendPaths(paths, cptool.cptoolHomeDirectory)
	paths = appendPaths(paths, path.Join(cptool.homeDirectory, ".cptool"))
	paths = appendPaths(paths, "/etc/cptool/")
	return paths
}

func appendPaths(paths []string, path string) []string {
	for _, p := range paths {
		if p == path {
			return paths
		}
	}
	return append(paths, path)
}
