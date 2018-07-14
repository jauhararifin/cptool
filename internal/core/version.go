package core

import (
	"fmt"
)

// GetVersion returns current version as string
func (cptool *CPTool) GetVersion() string {
	return fmt.Sprintf("v%d.%d.%d", cptool.MajorVersion, cptool.MinorVersion, cptool.PatchVersion)
}
