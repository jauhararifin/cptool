package core

import (
	"errors"
	"path"
	"testing"
	"time"

	"github.com/jauhararifin/cptool/internal/executioner"
)

func TestRun(t *testing.T) {
	cptool := newTest()
	memexec := getCptoolMemExec(cptool)
	executed := false
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == compileTestLanguage.RunScript {
			executed = true
		}
		return nil
	}

	solution := Solution{
		Name:        "sol",
		Language:    compileTestLanguage,
		Path:        "/sol.lang",
		LastUpdated: time.Now(),
	}
	err := cptool.Run(compileTestLanguage, solution, nil, nil, nil)
	if err != nil {
		t.Error(err)
	}
	if !executed {
		t.Error("run script should executed")
	}
}

func TestRunWithErrorCompilation(t *testing.T) {
	cptool := newTest()
	memexec := getCptoolMemExec(cptool)
	executed := false
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == compileTestLanguage.RunScript {
			executed = true
		}
		if m.GetPath() == compileTestLanguage.CompileScript {
			return errors.New("some error occured")
		}
		return nil
	}

	solution := Solution{
		Name:        "sol",
		Language:    compileTestLanguage,
		Path:        "/sol.lang",
		LastUpdated: time.Now(),
	}
	err := cptool.Run(compileTestLanguage, solution, nil, nil, nil)
	if err == nil {
		t.Error("Run should return an error")
	}
	if executed {
		t.Error("Run should not execute run script")
	}
}

func TestRunWithError(t *testing.T) {
	cptool := newTest()
	memexec := getCptoolMemExec(cptool)
	executed := false
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == compileTestLanguage.RunScript {
			executed = true
			return errors.New("some error occured")
		}
		return nil
	}

	solution := Solution{
		Name:        "sol",
		Language:    compileTestLanguage,
		Path:        "/sol.lang",
		LastUpdated: time.Now(),
	}
	err := cptool.Run(compileTestLanguage, solution, nil, nil, nil)
	if err == nil {
		t.Error("Run should return an error")
	}
	if !executed {
		t.Error("Run should execute run script")
	}
}

func TestRunByName(t *testing.T) {
	cptool := newTest()
	memexec := getCptoolMemExec(cptool)
	executed := false
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == compileTestLanguage.RunScript {
			executed = true
		}
		return nil
	}

	cptool.languages[compileTestLanguage.Name] = compileTestLanguage
	cptool.fs.Create(path.Join(cptool.workingDirectory, "solution.lang"))
	err := cptool.RunByName(compileTestLanguage.Name, "solution", nil, nil, nil)
	if err != nil {
		t.Error(err)
	}
	if !executed {
		t.Error("RunByName should execute run script")
	}
}

func TestRunByNameWithMissingSolution(t *testing.T) {
	cptool := newTest()
	memexec := getCptoolMemExec(cptool)
	executed := false
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == compileTestLanguage.RunScript {
			executed = true
		}
		return nil
	}

	cptool.languages[compileTestLanguage.Name] = compileTestLanguage
	err := cptool.RunByName(compileTestLanguage.Name, "solution", nil, nil, nil)
	if err == nil {
		t.Error("RunByName should return an error")
	}
	if err != ErrNoSuchSolution {
		t.Error("RunByName should return ErrNoSuchSolution error")
	}
	if executed {
		t.Error("RunByName should not execute run script")
	}
}

func TestRunByNameWithMissingLanguage(t *testing.T) {
	cptool := newTest()
	memexec := getCptoolMemExec(cptool)
	executed := false
	memexec.RunCallback = func(m *executioner.BaseCmd) error {
		if m.GetPath() == compileTestLanguage.RunScript {
			executed = true
		}
		return nil
	}

	cptool.fs.Create(path.Join(cptool.workingDirectory, "solution.lang"))
	err := cptool.RunByName(compileTestLanguage.Name, "solution", nil, nil, nil)
	if err == nil {
		t.Error("RunByName should return an error")
	}
	if err != ErrNoSuchLanguage {
		t.Error("RunByName should return ErrNoSuchLanguage error")
	}
	if executed {
		t.Error("RunByName should not execute run script")
	}
}
