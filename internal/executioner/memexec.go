package executioner

import (
	"context"
	"io"
	"os/exec"
)

// MemExec implements Exec using callbacks
type MemExec struct {
	CombinedOutputCallback func(*MemCmd) ([]byte, error)
	OutputCallback         func(*MemCmd) ([]byte, error)
	RunCallback            func(*MemCmd) error
	StartCallback          func(*MemCmd) error
	StderrPipeCallback     func(*MemCmd) (io.ReadCloser, error)
	StdinPipeCallback      func(*MemCmd) (io.WriteCloser, error)
	StdoutPipeCallback     func(*MemCmd) (io.ReadCloser, error)
	WaitCallback           func(*MemCmd) error
}

// MemCmd implements Cmd using memory
type MemCmd struct {
	*BaseCmd
	MemExec
	Context context.Context
}

// NewMemExec not yet defined
func NewMemExec() *MemExec {
	return &MemExec{}
}

// Command returns the Cmd struct to execute the named program with the given arguments
func (m *MemExec) Command(name string, arg ...string) Cmd {
	oscmd := exec.Command(name, arg...)
	basecmd := &BaseCmd{Cmd: oscmd}
	return &MemCmd{MemExec: *m, BaseCmd: basecmd, Context: nil}
}

// CommandContext create command with context
func (m *MemExec) CommandContext(ctx context.Context, name string, arg ...string) Cmd {
	oscmd := exec.CommandContext(ctx, name, arg...)
	basecmd := &BaseCmd{Cmd: oscmd}
	return &MemCmd{MemExec: *m, BaseCmd: basecmd, Context: ctx}
}

// CombinedOutput implements CombinedOutput of Cmd
func (m *MemCmd) CombinedOutput() ([]byte, error) {
	if m.CombinedOutputCallback != nil {
		return m.CombinedOutputCallback(m)
	}
	return make([]byte, 0), nil
}

// Output implements Output of Cmd
func (m *MemCmd) Output() ([]byte, error) {
	if m.OutputCallback != nil {
		return m.OutputCallback(m)
	}
	return make([]byte, 0), nil
}

// Run implements Run of Cmd
func (m *MemCmd) Run() error {
	if m.RunCallback != nil {
		return m.RunCallback(m)
	}
	return nil
}

// Start implements Start of Cmd
func (m *MemCmd) Start() error {
	if m.StartCallback != nil {
		return m.StartCallback(m)
	}
	return nil
}

// StderrPipe implements StderrPipe of Cmd
func (m *MemCmd) StderrPipe() (io.ReadCloser, error) {
	if m.StderrPipeCallback != nil {
		return m.StderrPipeCallback(m)
	}
	return nil, nil
}

// StdinPipe implements StdinPipe of Cmd
func (m *MemCmd) StdinPipe() (io.WriteCloser, error) {
	if m.StdinPipeCallback != nil {
		return m.StdinPipeCallback(m)
	}
	return nil, nil
}

// StdoutPipe implements StdoutPipe of Cmd
func (m *MemCmd) StdoutPipe() (io.ReadCloser, error) {
	if m.StdoutPipeCallback != nil {
		return m.StdoutPipeCallback(m)
	}
	return nil, nil
}

// Wait implements Wait of Cmd
func (m *MemCmd) Wait() error {
	if m.WaitCallback != nil {
		return m.WaitCallback(m)
	}
	return nil
}
