package core

import (
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestCleanCacheDirectory(t *testing.T) {
	cptool := newTest()
	cptool.fs.MkdirAll(cptool.GetCompilationRootDir(), os.ModePerm)
	cptool.fs.MkdirAll(cptool.GetOutputRootDir(), os.ModePerm)
	err := cptool.CleanCacheDirectory()
	if err != nil {
		t.Error(err)
	}
	if ok, _ := afero.DirExists(cptool.fs, cptool.GetCompilationRootDir()); ok {
		t.Error("CleanCacheDirectory should remove compilation directory")
	}
	if ok, _ := afero.DirExists(cptool.fs, cptool.GetOutputRootDir()); ok {
		t.Error("CleanCacheDirectory should remove output directory")
	}
}

func TestCleanCacheDirectoryAlreadyEmpty(t *testing.T) {
	cptool := newTest()
	err := cptool.CleanCacheDirectory()
	if err != nil {
		t.Error(err)
	}
	if ok, _ := afero.DirExists(cptool.fs, cptool.GetCompilationRootDir()); ok {
		t.Error("CleanCacheDirectory should remove compilation directory")
	}
	if ok, _ := afero.DirExists(cptool.fs, cptool.GetOutputRootDir()); ok {
		t.Error("CleanCacheDirectory should remove output directory")
	}
}
