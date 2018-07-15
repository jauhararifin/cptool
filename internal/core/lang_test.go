package core

import (
	"os"
	"testing"
)

func TestGetLanguageFromDirectory(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/langs/some_language_name", os.ModePerm)
	file, err := cptool.fs.Create("/langs/some_language_name/lang.conf")
	file.WriteString("verbose_name=\"Some Language\"\nextension=\"some_ext\"\n")
	cptool.fs.Create("/langs/some_language_name/compile")
	cptool.fs.Create("/langs/some_language_name/run")

	language, err := cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err != nil {
		t.Error(err)
	}
	if language.Name != "some_language_name" {
		t.Error("language name should be some_language_name")
	}
	if language.VerboseName != "Some Language" {
		t.Error("language verbose name should be \"Some Language\"")
	}
	if language.Extension != "some_ext" {
		t.Error("language extension should be some_ext")
	}
	if language.RunScript != "/langs/some_language_name/run" {
		t.Error("language run script should be \"/langs/some_language_name/run\"")
	}
	if language.CompileScript != "/langs/some_language_name/compile" {
		t.Error("language compile script should be \"/langs/some_language_name/compile\"")
	}
	if language.Debuggable {
		t.Error("language debuggable should be false")
	}
}

func TestGetLanguageFromDirectoryWithDebugScript(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/langs/some_language_name", os.ModePerm)
	file, err := cptool.fs.Create("/langs/some_language_name/lang.conf")
	file.WriteString("verbose_name=\"Some Language\"\nextension=\"some_ext\"\n")
	cptool.fs.Create("/langs/some_language_name/compile")
	cptool.fs.Create("/langs/some_language_name/run")
	cptool.fs.Create("/langs/some_language_name/debugcompile")

	language, err := cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err != nil {
		t.Error(err)
	}
	if language.Name != "some_language_name" {
		t.Error("language name should be some_language_name")
	}
	if language.VerboseName != "Some Language" {
		t.Error("language verbose name should be \"Some Language\"")
	}
	if language.Extension != "some_ext" {
		t.Error("language extension should be some_ext")
	}
	if language.RunScript != "/langs/some_language_name/run" {
		t.Error("language run script should be \"/langs/some_language_name/run\"")
	}
	if language.CompileScript != "/langs/some_language_name/compile" {
		t.Error("language compile script should be \"/langs/some_language_name/compile\"")
	}
	if language.DebugScript != "/langs/some_language_name/debugcompile" {
		t.Error("language compile script should be \"/langs/some_language_name/debugcompile\"")
	}
	if !language.Debuggable {
		t.Error("language debuggable should be true")
	}
}

func TestGetLanguageFromDirectoryWithoutConfFile(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/langs/some_language_name", os.ModePerm)
	cptool.fs.Create("/langs/some_language_name/compile")
	cptool.fs.Create("/langs/some_language_name/run")

	language, err := cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err != nil {
		t.Error(err)
	}
	if language.Name != "some_language_name" {
		t.Error("language name should be some_language_name")
	}
	if language.VerboseName != "some_language_name" {
		t.Error("language verbose name should be \"Some Language\"")
	}
	if language.Extension != "some_language_name" {
		t.Error("language extension should be some_ext")
	}
	if language.RunScript != "/langs/some_language_name/run" {
		t.Error("language run script should be \"/langs/some_language_name/run\"")
	}
	if language.CompileScript != "/langs/some_language_name/compile" {
		t.Error("language compile script should be \"/langs/some_language_name/compile\"")
	}
	if language.Debuggable {
		t.Error("language debuggable should be false")
	}
}

func TestGetLanguageFromDirectoryWithoutRunScript(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/langs/some_language_name", os.ModePerm)
	file, err := cptool.fs.Create("/langs/some_language_name/lang.conf")
	file.WriteString("verbose_name=\"Some Language\"\nextension=\"some_ext\"\n")
	cptool.fs.Create("/langs/some_language_name/compile")

	_, err = cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err == nil {
		t.Error("should return error when run script is missing")
	}

	cptool.fs.MkdirAll("/langs/some_language_name/run", os.ModePerm)
	_, err = cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err == nil {
		t.Error("should return error when run script file is a directory")
	}
}

func TestGetLanguageFromDirectoryWithoutCompileScript(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/langs/some_language_name", os.ModePerm)
	file, err := cptool.fs.Create("/langs/some_language_name/lang.conf")
	file.WriteString("verbose_name=\"Some Language\"\nextension=\"some_ext\"\n")
	cptool.fs.Create("/langs/some_language_name/run")

	_, err = cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err == nil {
		t.Error("should return error when compile script is missing")
	}

	cptool.fs.MkdirAll("/langs/some_language_name/compile", os.ModePerm)
	_, err = cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err == nil {
		t.Error("should return error when compile script file is a directory")
	}
}

func TestGetLanguageFromDirectoryWithInvalidConfFile(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/langs/some_language_name", os.ModePerm)
	file, err := cptool.fs.Create("/langs/some_language_name/lang.conf")
	file.WriteString("abcdefg")
	cptool.fs.Create("/langs/some_language_name/run")
	cptool.fs.Create("/langs/some_language_name/compile")

	_, err = cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err == nil {
		t.Error("should return error when configuration file is not in toml format")
	}
}

func TestGetLanguageFromDirectoryWithInvalidDirectory(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/langs", os.ModePerm)
	cptool.fs.Create("/langs/some_language_name")

	_, err := cptool.getLanguageFromDirectory("/langs/some_language_name")
	if err == nil {
		t.Error("should return invalid directory error")
	}
}
