package executioner

import (
	"context"
	"io"
	"os"
	"syscall"
)

// Cmd represents an external command being prepared or run
type Cmd interface {
	SetPath(string)
	SetArgs([]string)
	SetEnv([]string)
	SetDir(string)
	SetStdin(io.Reader)
	SetStdout(io.Writer)
	SetStderr(io.Writer)
	SetExtraFiles([]*os.File)
	SetSysProcAttr(*syscall.SysProcAttr)
	SetProcess(*os.Process)
	SetProcessState(*os.ProcessState)

	GetPath() string
	GetArgs() []string
	GetEnv() []string
	GetDir() string
	GetStdin() io.Reader
	GetStdout() io.Writer
	GetStderr() io.Writer
	GetExtraFiles() []*os.File
	GetSysProcAttr() *syscall.SysProcAttr
	GetProcess() *os.Process
	GetProcessState() *os.ProcessState

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
	CommandContext(ctx context.Context, name string, arg ...string) Cmd
}
