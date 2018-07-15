package core

import (
	"os"
	"path"
	"testing"
)

func TestGetSolution(t *testing.T) {
	cptool := newTest()
	language := Language{Extension: "c"}
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "a.c"))
	if err != nil {
		t.Error(err)
	}

	solution, err := cptool.GetSolution("a", language)
	if err != nil {
		t.Error(err)
	}
	if solution.Name != "a" {
		t.Error("solution name should be a")
	}
	if solution.Language != language {
		t.Error("language solution does not match")
	}
}

// test when solution is inside directory, not in current directory
func TestGetSolutionInsideDirectory(t *testing.T) {
	cptool := newTest()
	language := Language{Extension: "custom_extension"}
	_, err := cptool.fs.Create(path.Join(cptool.workingDirectory, "some/dir/a.custom_extension"))
	if err != nil {
		t.Error(err)
	}

	solution, err := cptool.GetSolution("some/dir/a", language)
	if err != nil {
		t.Error(err)
	}
	if solution.Name != "some/dir/a" {
		t.Error("solution name should be some/dir/a")
	}
	if solution.Language != language {
		t.Error("language solution does not match")
	}
}

func TestGetSolutionNotExists(t *testing.T) {
	cptool := newTest()
	language := Language{Extension: "c"}

	_, err := cptool.GetSolution("a", language)
	if err == nil {
		t.Error("should return error when solution does not exists")
	}
	if err != ErrNoSuchSolution {
		t.Error("should return no such solution error when solution does not exists")
	}
}

func TestGetSolutionIsDirectory(t *testing.T) {
	cptool := newTest()
	language := Language{Extension: "custom_extension"}
	err := cptool.fs.MkdirAll(path.Join(cptool.workingDirectory, "a.custom_extension"), os.ModePerm)
	if err != nil {
		t.Error(err)
	}

	_, err = cptool.GetSolution("a", language)
	if err == nil {
		t.Error("should return error when solution does not exists")
	}
	if err != ErrNoSuchSolution {
		t.Error("should return no such solution error when solution does not exists")
	}
}
