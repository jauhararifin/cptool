package core

import (
	"os"
	"path"
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

func TestLoadAllLanguage(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/etc/cptool/langs/lang_a", os.ModePerm)
	file, _ := cptool.fs.Create("/etc/cptool/langs/lang_a/lang.conf")
	file.WriteString("verbose_name=\"Language A\"\nextension=\"a\"\n")
	cptool.fs.Create("/etc/cptool/langs/lang_a/compile")
	cptool.fs.Create("/etc/cptool/langs/lang_a/run")

	cptool.loadAllLanguages()

	if len(cptool.languages) != 1 {
		t.Error("cptool loaded languages should be 1")
	}

	languageA := Language{
		Name:          "lang_a",
		Extension:     "a",
		VerboseName:   "Language A",
		CompileScript: "/etc/cptool/langs/lang_a/compile",
		RunScript:     "/etc/cptool/langs/lang_a/run",
		DebugScript:   "",
		Debuggable:    false,
	}
	lang, ok := cptool.languages["lang_a"]
	if !ok {
		t.Error("cptool should contain language lang_a")
	}
	if lang != languageA {
		t.Error("language lang_a of cptool should be", languageA, ", but found", lang)
	}
}

func TestLoadAllLanguagesWithMultipleLanguages(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/etc/cptool/langs/lang_a", os.ModePerm)
	file, _ := cptool.fs.Create("/etc/cptool/langs/lang_a/lang.conf")
	file.WriteString("verbose_name=\"Language A\"\nextension=\"a\"\n")
	cptool.fs.Create("/etc/cptool/langs/lang_a/compile")
	cptool.fs.Create("/etc/cptool/langs/lang_a/run")

	cptool.fs.MkdirAll("/etc/cptool/langs/lang_b", os.ModePerm)
	file, _ = cptool.fs.Create("/etc/cptool/langs/lang_b/lang.conf")
	file.WriteString("verbose_name=\"Language B\"\nextension=\"b\"\n")
	cptool.fs.Create("/etc/cptool/langs/lang_b/compile")
	cptool.fs.Create("/etc/cptool/langs/lang_b/run")

	cptool.fs.MkdirAll(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c"), os.ModePerm)
	file, _ = cptool.fs.Create(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/lang.conf"))
	file.WriteString("verbose_name=\"Language C\"\nextension=\"c\"\n")
	cptool.fs.Create(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/compile"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/run"))

	cptool.loadAllLanguages()

	if len(cptool.languages) != 3 {
		t.Error("cptool loaded languages should be 3")
	}

	languageA := Language{
		Name:          "lang_a",
		Extension:     "a",
		VerboseName:   "Language A",
		CompileScript: "/etc/cptool/langs/lang_a/compile",
		RunScript:     "/etc/cptool/langs/lang_a/run",
		DebugScript:   "",
		Debuggable:    false,
	}
	lang, ok := cptool.languages["lang_a"]
	if !ok {
		t.Error("cptool should contain language lang_a")
	}
	if lang != languageA {
		t.Error("language lang_a of cptool should be", languageA, ", but found", lang)
	}

	languageB := Language{
		Name:          "lang_b",
		Extension:     "b",
		VerboseName:   "Language B",
		CompileScript: "/etc/cptool/langs/lang_b/compile",
		RunScript:     "/etc/cptool/langs/lang_b/run",
		DebugScript:   "",
		Debuggable:    false,
	}
	lang, ok = cptool.languages["lang_b"]
	if !ok {
		t.Error("cptool should contain language lang_b")
	}
	if lang != languageB {
		t.Error("language lang_a of cptool should be", languageB, ", but found", lang)
	}

	languageC := Language{
		Name:          "lang_c",
		Extension:     "c",
		VerboseName:   "Language C",
		CompileScript: path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/compile"),
		RunScript:     path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/run"),
		DebugScript:   "",
		Debuggable:    false,
	}
	lang, ok = cptool.languages["lang_c"]
	if !ok {
		t.Error("cptool should contain language lang_c")
	}
	if lang != languageC {
		t.Error("language lang_c of cptool should be", languageC, ", but found", lang)
	}
}

func TestLoadAllLanguagesWithInvalidLanguages(t *testing.T) {
	cptool := newTest()
	cptool.fs.MkdirAll("/etc/cptool/langs/lang_a", os.ModePerm)
	cptool.fs.MkdirAll("/etc/cptool/langs/lang_b", os.ModePerm)
	cptool.fs.MkdirAll(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c"), os.ModePerm)
	cptool.loadAllLanguages()
	if len(cptool.languages) > 0 {
		t.Error("cptool should not loaded any invalid languages")
	}
}

func TestLoadAllLanguagesWithMultipleName(t *testing.T) {
	cptool := newTest()

	cptool.fs.MkdirAll("/etc/cptool/langs/lang_c", os.ModePerm)
	file, _ := cptool.fs.Create("/etc/cptool/langs/lang_c/lang.conf")
	file.WriteString("verbose_name=\"Language A\"\nextension=\"a\"\n")
	cptool.fs.Create("/etc/cptool/langs/lang_c/compile")
	cptool.fs.Create("/etc/cptool/langs/lang_c/run")

	cptool.fs.MkdirAll(path.Join(cptool.homeDirectory, ".cptool/langs/lang_c"), os.ModePerm)
	file, _ = cptool.fs.Create(path.Join(cptool.homeDirectory, ".cptool/langs/lang_c/lang.conf"))
	file.WriteString("verbose_name=\"Language C\"\nextension=\"c\"\n")
	cptool.fs.Create(path.Join(cptool.homeDirectory, ".cptool/langs/lang_c/compile"))
	cptool.fs.Create(path.Join(cptool.homeDirectory, ".cptool/langs/lang_c/run"))

	cptool.fs.MkdirAll(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c"), os.ModePerm)
	file, _ = cptool.fs.Create(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/lang.conf"))
	file.WriteString("verbose_name=\"Language C\"\nextension=\"c\"\n")
	cptool.fs.Create(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/compile"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/run"))

	cptool.loadAllLanguages()

	if len(cptool.languages) != 1 {
		t.Error("cptool loaded languages should be 1")
	}

	languageC := Language{
		Name:          "lang_c",
		Extension:     "c",
		VerboseName:   "Language C",
		CompileScript: path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/compile"),
		RunScript:     path.Join(cptool.workingDirectory, ".cptool/langs/lang_c/run"),
		DebugScript:   "",
		Debuggable:    false,
	}
	lang, ok := cptool.languages["lang_c"]
	if !ok {
		t.Error("cptool should contain language lang_c")
	}
	if lang != languageC {
		t.Error("language lang_c of cptool should be", languageC, ", but found", lang)
	}
}
