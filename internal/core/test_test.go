package core

import (
	"context"
	"path"
	"testing"
	"time"
)

func TestTest(t *testing.T) {
	cptool := newTest()
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.1.out"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.in"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "test.2.out"))
	cptool.fs.Create(path.Join(cptool.workingDirectory, "solution.lang"))
	solution := Solution{
		Name:        "solution",
		Language:    compileTestLanguage,
		Path:        path.Join(cptool.workingDirectory, "solution.lang"),
		LastUpdated: time.Now(),
	}

	err := cptool.Test(context.Background(), compileTestLanguage, solution, "test")
	if err != nil {
		t.Error(err)
	}
}
