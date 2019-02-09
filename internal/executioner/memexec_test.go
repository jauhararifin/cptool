package executioner

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestMemExecCommandCombinedOutput(t *testing.T) {
	e := NewMemExec()
	executed := false
	byteArr := []byte{1, 2}
	e.CombinedOutputCallback = func(cmd *MemCmd) ([]byte, error) {
		executed = true
		return byteArr, nil
	}
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	r, err := cmd.CombinedOutput()
	if !executed {
		t.Error("CombinedOutputCallback should executed")
	}
	if err != nil {
		t.Error("CombinedOutput shouldn't returns any error")
	}
	if len(r) != len(byteArr) || r[0] != byteArr[0] || r[1] != byteArr[1] {
		t.Error("Bytes differ")
	}
}

func TestMemExecCommandCombinedOutputNil(t *testing.T) {
	e := NewMemExec()
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	r, err := cmd.CombinedOutput()
	if err != nil {
		t.Error("CombinedOutput shouldn't returns any error")
	}
	if len(r) != 0 {
		t.Error("CombinedOutputCallback should returns zero byte")
	}
}

func TestMemExecCommandOutput(t *testing.T) {
	e := NewMemExec()
	executed := false
	byteArr := []byte{1, 2}
	e.OutputCallback = func(cmd *MemCmd) ([]byte, error) {
		executed = true
		return byteArr, nil
	}
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	r, err := cmd.Output()
	if !executed {
		t.Error("OutputCallback should executed")
	}
	if err != nil {
		t.Error("Output shouldn't returns any error")
	}
	if len(r) != len(byteArr) || r[0] != byteArr[0] || r[1] != byteArr[1] {
		t.Error("Bytes differ")
	}
}

func TestMemExecCommandOutputNil(t *testing.T) {
	e := NewMemExec()
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	r, err := cmd.Output()
	if err != nil {
		t.Error("Output shouldn't returns any error")
	}
	if len(r) != 0 {
		t.Error("OutputCallback should returns zero byte")
	}
}

func TestMemExecCommandRun(t *testing.T) {
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

func TestMemExecCommandRunNil(t *testing.T) {
	e := NewMemExec()
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	err := cmd.Run()
	if err != nil {
		t.Error("RUn shouldn't returns any error")
	}
}

func TestMemExecCommandContextRun(t *testing.T) {
	e := NewMemExec()
	executed := false
	ctx := context.Background()
	e.RunCallback = func(cmd *MemCmd) error {
		executed = true
		if cmd.GetPath() != "commandname" {
			t.Error("path should contains commandname")
		}
		if cmd.Context != ctx {
			t.Error("context differ")
		}
		return nil
	}
	cmd := e.CommandContext(ctx, "commandname")
	err := cmd.Run()
	if !executed {
		t.Error("Run should executed")
	}
	if err != nil {
		t.Error("RUn shouldn't returns any error")
	}
}

func TestMemExecCommandStart(t *testing.T) {
	e := NewMemExec()
	executed := false
	e.StartCallback = func(cmd *MemCmd) error {
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
	err := cmd.Start()
	if !executed {
		t.Error("StartCallback should executed")
	}
	if err != nil {
		t.Error("StartCallback shouldn't returns any error")
	}
}

func TestMemExecCommandStartNil(t *testing.T) {
	e := NewMemExec()
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	err := cmd.Start()
	if err != nil {
		t.Error("StartCallback shouldn't returns any error")
	}
}

func TestMemExecCommandStderrPipe(t *testing.T) {
	e := NewMemExec()
	executed := false
	reader := ioutil.NopCloser(bytes.NewReader([]byte("foo")))
	e.StderrPipeCallback = func(cmd *MemCmd) (io.ReadCloser, error) {
		executed = true
		return reader, nil
	}
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	io, err := cmd.StderrPipe()
	if !executed {
		t.Error("StderrPipeCallback should executed")
	}
	if err != nil {
		t.Error("StderrPipeCallback shouldn't returns any error")
	}
	if io != reader {
		t.Error("Reader differ")
	}
}

func TestMemExecCommandStderrPipeNil(t *testing.T) {
	e := NewMemExec()
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	reader, err := cmd.StderrPipe()
	if err != nil {
		t.Error("StderrPipeCallback shouldn't returns any error")
	}
	if reader == nil {
		t.Error("StderrPipeCallback shouldn't returns nil reader")
	}
}

func TestMemExecCommandStdinPipe(t *testing.T) {
	e := NewMemExec()
	executed := false
	e.StdinPipeCallback = func(cmd *MemCmd) (io.WriteCloser, error) {
		executed = true
		return os.Stdin, nil
	}
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	io, err := cmd.StdinPipe()
	if !executed {
		t.Error("StdinPipeCallback should executed")
	}
	if err != nil {
		t.Error("StdinPipeCallback shouldn't returns any error")
	}
	if io != os.Stdin {
		t.Error("Reader differ")
	}
}

func TestMemExecCommandStdinPipeNil(t *testing.T) {
	e := NewMemExec()
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	writer, err := cmd.StdinPipe()
	if err != nil {
		t.Error("StdinPipeCallback shouldn't returns any error")
	}
	if writer == nil {
		t.Error("StdinPipeCallback shouldn't returns nil writer")
	}
}

func TestMemExecCommandStdoutPipe(t *testing.T) {
	e := NewMemExec()
	executed := false
	reader := ioutil.NopCloser(bytes.NewReader([]byte("bar")))
	e.StdoutPipeCallback = func(cmd *MemCmd) (io.ReadCloser, error) {
		executed = true
		return reader, nil
	}
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	io, err := cmd.StdoutPipe()
	if !executed {
		t.Error("StdoutPipeCallback should executed")
	}
	if err != nil {
		t.Error("StdoutPipeCallback shouldn't returns any error")
	}
	if io != reader {
		t.Error("Reader differ")
	}
}

func TestMemExecCommandStdoutPipeNil(t *testing.T) {
	e := NewMemExec()
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	reader, err := cmd.StdoutPipe()
	if err != nil {
		t.Error("StdoutPipeCallback shouldn't returns any error")
	}
	if reader == nil {
		t.Error("StdoutPipeCallback shouldn't returns nil reader")
	}
}

func TestMemExecCommandWait(t *testing.T) {
	e := NewMemExec()
	executed := false
	e.WaitCallback = func(cmd *MemCmd) error {
		executed = true
		return nil
	}
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	err := cmd.Wait()
	if !executed {
		t.Error("WaitCallback should executed")
	}
	if err != nil {
		t.Error("WaitCallback shouldn't returns any error")
	}
}

func TestMemExecCommandWaitNil(t *testing.T) {
	e := NewMemExec()
	cmd := e.Command("commandname", "Arg1", "Arg2", "Arg3")
	err := cmd.Wait()
	if err != nil {
		t.Error("WaitCallback shouldn't returns any error")
	}
}
