package executioner

import (
	"testing"
)

func TestMemExecCommand(t *testing.T) {
	e := NewMemExec()
	executed := false
	e.RunCallback = func(cmd *MemCmd) error {
		executed = true
		if cmd.Context != nil {
			t.Error("context should be nil")
		}
		if cmd.GetPath() != "commandname" {
			t.Error("path should contains commandname")
		}
		args := cmd.GetArgs()
		if len(args) != 4 {
			t.Error("args should contains 3 elements")
		}
		if args[0] != "commandname" {
			t.Error("first arg should contains commandname")
		}
		if args[1] != "Arg1" {
			t.Error("second arg should contains Arg1")
		}
		if args[2] != "Arg2" {
			t.Error("third arg should contains Arg2")
		}
		if args[3] != "Arg3" {
			t.Error("forth arg should contains Arg3")
		}
		return nil
	}
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	err := cmd.Run()
	if !executed {
		t.Error("Run should executed")
	}
	if err != nil {
		t.Error("RUn shouldn't returns any error")
	}
}
