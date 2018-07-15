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
	return MemExec{}
}

// Command returns the Cmd struct to execute the named program with the given arguments
func (MemExec) Command(name string, arg ...string) Cmd {
	return &MemCmd{}
}

// CombinedOutput implements CombinedOutput of Cmd
func (m *MemCmd) CombinedOutput() ([]byte, error) {
	if m.CombinedOutputCallback != nil {
		m.CombinedOutputCallback()
	}
	return make([]byte, 0), nil
}

// Output implements Output of Cmd
func (m *MemCmd) Output() ([]byte, error) {
	if m.OutputCallback != nil {
		m.OutputCallback()
	}
	return make([]byte, 0), nil
}

// Run implements Run of Cmd
func (m *MemCmd) Run() error {
	if m.RunCallback != nil {
		m.RunCallback()
	}
	return nil
}

// Start implements Start of Cmd
func (m *MemCmd) Start() error {
	if m.StartCallback != nil {
		m.StartCallback()
	}
	return nil
}

// StderrPipe implements StderrPipe of Cmd
func (m *MemCmd) StderrPipe() (io.ReadCloser, error) {
	if m.StderrPipeCallback != nil {
		m.StderrPipeCallback()
	}
	return nil, nil
}

// StdinPipe implements StdinPipe of Cmd
func (m *MemCmd) StdinPipe() (io.WriteCloser, error) {
	if m.StdinPipeCallback != nil {
		m.StdinPipeCallback()
	}
	return nil, nil
}

// StdoutPipe implements StdoutPipe of Cmd
func (m *MemCmd) StdoutPipe() (io.ReadCloser, error) {
	if m.StdoutPipeCallback != nil {
		m.StdoutPipeCallback()
	}
	return nil, nil
}

// Wait implements Wait of Cmd
func (m *MemCmd) Wait() error {
	if m.WaitCallback != nil {
		m.WaitCallback()
	}
	return nil
}
