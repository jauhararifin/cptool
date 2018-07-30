package core

import (
	"github.com/spf13/afero"
)

// CleanCacheDirectory clean cptool cache directory that contains compiled solution and output.
func (cptool *CPTool) CleanCacheDirectory() error {
	compileDir := cptool.GetCompilationRootDir()
	if ok, err := afero.DirExists(cptool.fs, compileDir); ok && err == nil {
		err := cptool.fs.RemoveAll(compileDir)
		if err != nil {
			return err
		}
	}
	outputDir := cptool.GetOutputRootDir()
	if ok, err := afero.DirExists(cptool.fs, outputDir); ok && err == nil {
		err := cptool.fs.RemoveAll(outputDir)
		if err != nil {
			return err
		}
	}
	return nil
}
