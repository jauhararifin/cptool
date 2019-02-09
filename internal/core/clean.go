package core

import (
	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/spf13/afero"
)

// CleanCacheDirectory clean cptool cache directory that contains compiled solution and output.
func (cptool *CPTool) CleanCacheDirectory() error {
	compileDir := cptool.GetCompilationRootDir()
	if ok, err := afero.DirExists(cptool.fs, compileDir); ok && err == nil {
		if cptool.logger != nil {
			cptool.logger.Println(logger.VERBOSE, "Removing program compilation output", compileDir)
		}
		err := cptool.fs.RemoveAll(compileDir)
		if err != nil {
			return err
		}
	}
	outputDir := cptool.GetOutputRootDir()
	if ok, err := afero.DirExists(cptool.fs, outputDir); ok && err == nil {
		if cptool.logger != nil {
			cptool.logger.Println(logger.VERBOSE, "Removing program execution output", outputDir)
		}
		err := cptool.fs.RemoveAll(outputDir)
		if err != nil {
			return err
		}
	}
	return nil
}
