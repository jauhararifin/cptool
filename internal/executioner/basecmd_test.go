package executioner

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func TestBaseCmdSetter(t *testing.T) {
	baseCmd := BaseCmd{&exec.Cmd{}}

	baseCmd.SetPath("somepathtotest")
	if baseCmd.Path != "somepathtotest" {
		t.Error("Path should contains somepathtotest")
	}

	baseCmd.SetArgs([]string{"arg1", "a", "..."})
	if len(baseCmd.Args) != 3 {
		t.Error("Args should contains 3 elements")
	}
	if baseCmd.Args[0] != "arg1" {
		t.Error("Args[0] should contains \"arg1\"")
	}
	if baseCmd.Args[1] != "a" {
		t.Error("Args[1] should contains \"a\"")
	}
	if baseCmd.Args[2] != "..." {
		t.Error("Args[2] should contains \"...\"")
	}

	baseCmd.SetEnv([]string{
		"key=value",
		"URL=example.com",
	})
	if len(baseCmd.Env) != 2 {
		t.Error("Env should contains 2 elements")
	}
	if baseCmd.Env[0] != "key=value" {
		t.Error("Env[0] should contains key=value")
	}
	if baseCmd.Env[1] != "URL=example.com" {
		t.Error("Env[1] should contains URL=example.com")
	}

	baseCmd.SetDir("/root")
	if baseCmd.Dir != "/root" {
		t.Error("Dir should contains /root")
	}

	reader := strings.NewReader("somereader")
	baseCmd.SetStdin(reader)
	if baseCmd.Stdin != reader {
		t.Error("Stdin mismatch")
	}

	writer := new(bytes.Buffer)
	baseCmd.SetStdout(writer)
	if baseCmd.Stdout != writer {
		t.Error("Stdout mismatch")
	}

	writerErr := new(bytes.Buffer)
	baseCmd.SetStderr(writerErr)
	if baseCmd.Stderr != writerErr {
		t.Error("Stderr mismatch")
	}
}

func TestBaseCmdGetter(t *testing.T) {
	baseCmd := BaseCmd{&exec.Cmd{}}

	baseCmd.Path = "somepathtotest"
	if baseCmd.GetPath() != "somepathtotest" {
		t.Error("GetPath should returns somepathtotest")
	}

	baseCmd.Args = []string{"arg1", "a", "..."}
	if len(baseCmd.GetArgs()) != 3 {
		t.Error("GetArgs should contains 3 elements")
	}
	if baseCmd.GetArgs()[0] != "arg1" {
		t.Error("GetArgs's first element should contains arg1")
	}
	if baseCmd.GetArgs()[1] != "a" {
		t.Error("GetArgs's second element should contains a")
	}
	if baseCmd.GetArgs()[2] != "..." {
		t.Error("GetArgs's third element should contains ...")
	}

	baseCmd.Env = []string{
		"key=value",
		"URL=example.com",
	}
	if len(baseCmd.GetEnv()) != 2 {
		t.Error("GetEnv should returns 2 elemnts")
	}
	if baseCmd.GetEnv()[0] != "key=value" {
		t.Error("GetEnv's first element should contains key=value")
	}
	if baseCmd.GetEnv()[1] != "URL=example.com" {
		t.Error("GetEnv's second element should contans URL=example.com")
	}

	baseCmd.Dir = "/root"
	if baseCmd.GetDir() != "/root" {
		t.Error("GetDir should returns /root")
	}

	reader := strings.NewReader("someinput")
	baseCmd.Stdin = reader
	if baseCmd.GetStdin() != reader {
		t.Error("GetStdin returns wrong reader")
	}

	writer := new(bytes.Buffer)
	baseCmd.Stdout = writer
	if baseCmd.GetStdout() != writer {
		t.Error("GetStdout returns wrong writer")
	}

	writerErr := new(bytes.Buffer)
	baseCmd.Stderr = writerErr
	if baseCmd.GetStderr() != writerErr {
		t.Error("GetStderr returns wrong writer")
	}
}
