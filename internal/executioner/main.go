package executioner

import "io"

// Cmd represents an external command being prepared or run
type Cmd interface {
	CombinedOutput() ([]byte, error)
	Output() ([]byte, error)
	Run() error
	Start() error
	StderrPipe() (io.ReadCloser, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	Wait() error
}

// Exec runs external commands.
type Exec interface {
	Command(name string, arg ...string) Cmd
}
