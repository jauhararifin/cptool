package core

import (
	"errors"
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/jauhararifin/cptool/internal/executioner"
)

var compileTestLanguage = Language{
	Name:          "some_lang",
	Extension:     "lang",
	VerboseName:   "Some Language",
	CompileScript: "/compile",
	RunScript:     "/run",
	DebugScript:   "",
	Debuggable:    false,
}

var compileTestLanguageDebuggable = Language{
	Name:          "some_lang",
	Extension:     "lang",
	VerboseName:   "Some Language",
	CompileScript: "/compile",
	RunScript:     "/run",
	DebugScript:   "/debug",
	Debuggable:    true,
}

func TestCompile(t *testing.T) {
	cptool := newTest()
	cptool.languages["some_lang"] = compileTestLanguage
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "a.lang"))
	if err != nil {
		t.Fail()
	}
	executed := false
	memexec := getCptoolMemExec(cptool)
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == compileTestLanguage.CompileScript {
			executed = true
		}
		return nil
	}

	err = cptool.CompileByName("some_lang", "a", false)
	if err != nil {
		t.Error("Compile should compile code successfully succesfully")
	}
	if !executed {
		t.Error("Compile should execute compilation script")
	}
}

func TestCompileWithDebug(t *testing.T) {
	cptool := newTest()
	executed := false
	memexec := getCptoolMemExec(cptool)
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		fmt.Println(m.GetPath(), compileTestLanguageDebuggable.DebugScript)
		if m.GetPath() == compileTestLanguageDebuggable.DebugScript {
			executed = true
		}
		return nil
	}
	cptool.languages["some_lang"] = compileTestLanguageDebuggable
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "a.lang"))
	if err != nil {
		t.Fail()
	}

	err = cptool.CompileByName("some_lang", "a", true)
	if err != nil {
		t.Error("Compile should compile code successfully succesfully")
	}
	if !executed {
		t.Error("Compile should execute compilation script")
	}
}
func TestCompileWithError(t *testing.T) {
	cptool := newTest()
	executed := false
	memexec := getCptoolMemExec(cptool)
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == compileTestLanguage.CompileScript {
			executed = true
		}
		return errors.New("just some error")
	}
	cptool.languages["some_lang"] = compileTestLanguage
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "a.lang"))
	if err != nil {
		t.Fail()
	}

	err = cptool.CompileByName("some_lang", "a", false)
	if err == nil {
		t.Error("Compile should return error")
	}
	if !executed {
		t.Error("Compile should execute compilation script")
	}
}

func TestCompileWithDebugInNonDebuggableLanguage(t *testing.T) {
	cptool := newTest()
	executed := false
	memexec := getCptoolMemExec(cptool)
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		executed = true
		return nil
	}
	cptool.languages["some_lang"] = compileTestLanguage
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "a.lang"))
	if err != nil {
		t.Fail()
	}

	err = cptool.CompileByName("some_lang", "a", true)
	if err == nil || err != ErrLanguageNotDebuggable {
		t.Error("Compile should return ErrLanguageNotDebuggable")
	}
	if executed {
		t.Error("Compile should not execute the script")
	}
}

func TestCompileWithMissingLanguage(t *testing.T) {
	cptool := newTest()
	cptool.languages["some_lang"] = compileTestLanguage
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "a.lang"))
	if err != nil {
		t.Fail()
	}

	err = cptool.CompileByName("some_lang_not_found", "b", false)
	if err == nil || err != ErrNoSuchLanguage {
		t.Error("Compile should return ErrNoSuchLanguage")
	}
}

func TestCompileWithMissingSolution(t *testing.T) {
	cptool := newTest()
	cptool.languages["some_lang"] = compileTestLanguage
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "a.lang"))
	if err != nil {
		t.Fail()
	}

	err = cptool.CompileByName("some_lang", "b", false)
	if err == nil || err != ErrNoSuchSolution {
		t.Error("Compile should return ErrNoSuchSolution")
	}
}

func TestCompileWithDueDate(t *testing.T) {
	cptool := newTest()
	executed := false
	memexec := getCptoolMemExec(cptool)
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		executed = true
		return nil
	}
	cptool.languages["some_lang"] = compileTestLanguage
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "a.lang"))
	if err != nil {
		t.Fail()
	}
	targetPath := path.Join(cptool.workingDirectory, ".cptool/solutions/a/some_lang/program")
	_, err = cptool.fs.Create(targetPath)
	if err != nil {
		t.Fail()
	}
	targetTime := time.Now().Add(time.Minute)
	cptool.fs.Chtimes(targetPath, targetTime, targetTime)

	err = cptool.CompileByName("some_lang", "a", false)
	if err != nil {
		t.Error("Compile should not return error")
	}
	if executed {
		t.Error("Compile should not execute the script")
	}
}
