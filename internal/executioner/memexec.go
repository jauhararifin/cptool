package executioner

import (
	"io"
)

// MemExec implements Exec using callbacks
type MemExec struct {
	CombinedOutputCallback func() ([]byte, error)
	OutputCallback         func() ([]byte, error)
	RunCallback            func() error
	StartCallback          func() error
	StderrPipeCallback     func() (io.ReadCloser, error)
	StdinPipeCallback      func() (io.WriteCloser, error)
	StdoutPipeCallback     func() (io.ReadCloser, error)
	WaitCallback           func() error
}

// MemCmd implements Cmd using memory
type MemCmd struct {
	MemExec
}

// NewMemExec not yet defined
func NewMemExec() Exec {
	return &MemExec{}
}

// Command returns the Cmd struct to execute the named program with the given arguments
func (m *MemExec) Command(name string, arg ...string) Cmd {
	return &MemCmd{MemExec: *m}
}

// CombinedOutput implements CombinedOutput of Cmd
func (m *MemCmd) CombinedOutput() ([]byte, error) {
	if m.CombinedOutputCallback != nil {
		return m.CombinedOutputCallback()
	}
	return make([]byte, 0), nil
}

// Output implements Output of Cmd
func (m *MemCmd) Output() ([]byte, error) {
	if m.OutputCallback != nil {
		return m.OutputCallback()
	}
	return make([]byte, 0), nil
}

// Run implements Run of Cmd
func (m *MemCmd) Run() error {
	if m.RunCallback != nil {
		return m.RunCallback()
	}
	return nil
}

// Start implements Start of Cmd
func (m *MemCmd) Start() error {
	if m.StartCallback != nil {
		return m.StartCallback()
	}
	return nil
}

// StderrPipe implements StderrPipe of Cmd
func (m *MemCmd) StderrPipe() (io.ReadCloser, error) {
	if m.StderrPipeCallback != nil {
		return m.StderrPipeCallback()
	}
	return nil, nil
}

// StdinPipe implements StdinPipe of Cmd
func (m *MemCmd) StdinPipe() (io.WriteCloser, error) {
	if m.StdinPipeCallback != nil {
		return m.StdinPipeCallback()
	}
	return nil, nil
}

// StdoutPipe implements StdoutPipe of Cmd
func (m *MemCmd) StdoutPipe() (io.ReadCloser, error) {
	if m.StdoutPipeCallback != nil {
		return m.StdoutPipeCallback()
	}
	return nil, nil
}

// Wait implements Wait of Cmd
func (m *MemCmd) Wait() error {
	if m.WaitCallback != nil {
		return m.WaitCallback()
	}
	return nil
}
