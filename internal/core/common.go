package core

import (
	"os"
	"os/user"
	"path"
)

// GetConfigurationPaths returns all paths that considered contain cptool configuration
func GetConfigurationPaths() []string {
	paths := make([]string, 0)

	cwd, err := os.Getwd()
	if err == nil {
		paths = append(paths, path.Join(cwd, ".cptool"))
	}

	usr, err := user.Current()
	if err == nil {
		paths = append(paths, path.Join(usr.HomeDir, ".cptool"))
	}

	paths = append(paths, "/etc/cptool/")

	return paths
}
