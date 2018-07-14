package core

import (
	"path"
	"testing"
)

func TestGetConfigurationPaths(t *testing.T) {
	cptool := newTest()

	paths := cptool.GetConfigurationPaths()

	if len(paths) != 3 {
		t.Error("configuration paths should contain 3 elements")
	}

	cwd := cptool.workingDirectory
	if paths[0] != path.Join(cwd, ".cptool") {
		t.Errorf("first path should contain current working directory path, found: %s", paths[0])
	}

	home := cptool.homeDirectory
	if paths[1] != path.Join(home, ".cptool") {
		t.Errorf("second path should contain current home directory, found: %s", paths[1])
	}

	if paths[2] != "/etc/cptool" && paths[2] != "/etc/cptool/" {
		t.Errorf("third path should contain /etc/cptool, found: %s", paths[2])
	}
}
