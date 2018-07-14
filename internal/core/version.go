package core

import (
	"fmt"
)

// MajorVersion incidates current cptool major version
const MajorVersion = 1

// MinorVersion incidates current cptool minor version
const MinorVersion = 0

// PatchVersion incidates current cptool patch version
const PatchVersion = 0

// GetVersion returns current version as string
func GetVersion() string {
	return fmt.Sprintf("v%d.%d.%d", MajorVersion, MinorVersion, MinorVersion)
}
