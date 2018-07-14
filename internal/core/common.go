package core

import (
	"path"
)

// GetConfigurationPaths returns all paths that considered contain cptool configuration
func (cptool *CPTool) GetConfigurationPaths() []string {
	paths := make([]string, 0)
	paths = append(paths, path.Join(cptool.workingDirectory, ".cptool"))
	paths = append(paths, path.Join(cptool.homeDirectory, ".cptool"))
	paths = append(paths, "/etc/cptool/")
	return paths
}
