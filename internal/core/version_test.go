package core

import (
	"fmt"
	"testing"
)

func TestGetVersion(t *testing.T) {
	cptool := newTest()
	version := cptool.GetVersion()
	expected := fmt.Sprint("v", cptool.MajorVersion, ".", cptool.MinorVersion, ".", cptool.PatchVersion)
	if version != expected {
		t.Errorf("version should be %s, got %s", expected, version)
	}
}
