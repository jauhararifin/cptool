package version

import (
	"fmt"
)

// GetMajorVersion returns current major version
func GetMajorVersion() int {
	return 1
}

// GetMinorVersion returns current minor version
func GetMinorVersion() int {
	return 0
}

// GetPatchVersion returns current patch version
func GetPatchVersion() int {
	return 0
}

// GetVersion returns current version as string
func GetVersion() string {
	return fmt.Sprintf("v%d.%d.%d", GetMajorVersion(), GetMinorVersion(), GetPatchVersion())
}
