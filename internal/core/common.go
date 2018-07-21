package core

import (
	"path"
)

func appendPaths(paths []string, path string) []string {
	for _, p := range paths {
		if p == path {
			return paths
		}
	}
	return append(paths, path)
}

// GetConfigurationPaths returns all paths that considered contain cptool configuration
func (cptool *CPTool) GetConfigurationPaths() []string {
	paths := make([]string, 0)
	paths = appendPaths(paths, path.Join(cptool.workingDirectory, ".cptool"))
	paths = appendPaths(paths, path.Join(cptool.homeDirectory, ".cptool"))
	paths = appendPaths(paths, "/etc/cptool/")
	return paths
}
