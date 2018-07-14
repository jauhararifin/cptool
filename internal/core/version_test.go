package core

import (
	"fmt"
	"testing"
)

func TestGetVersion(t *testing.T) {
	version := newTest().GetVersion()
	if version != fmt.Sprint("v", MajorVersion, ".", MinorVersion, ".", PatchVersion) {
		t.Fail()
	}
}
