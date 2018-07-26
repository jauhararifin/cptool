package executioner

import (
	"io"
	"os"
	"os/exec"
	"syscall"
)

// BaseCmd implements basic functions of cmd
type BaseCmd struct {
	*exec.Cmd
}

// SetPath set Command's Path
func (b *BaseCmd) SetPath(value string) {
	b.Path = value
}

// SetArgs set Command's Args
func (b *BaseCmd) SetArgs(value []string) {
	b.Args = value
}

// SetEnv set Command's Env
func (b *BaseCmd) SetEnv(value []string) {
	b.Env = value
}

// SetDir set Command's Dir
func (b *BaseCmd) SetDir(value string) {
	b.Dir = value
}

// SetStdin set Command's Stdin
func (b *BaseCmd) SetStdin(value io.Reader) {
	b.Stdin = value
}

// SetStdout set Command's Stdout
func (b *BaseCmd) SetStdout(value io.Writer) {
	b.Stdout = value
}

// SetStderr set Command's Stderr
func (b *BaseCmd) SetStderr(value io.Writer) {
	b.Stderr = value
}

// SetExtraFiles set Command's ExtraFiles
func (b *BaseCmd) SetExtraFiles(value []*os.File) {
	b.ExtraFiles = value
}

// SetSysProcAttr set Command's SysProcAttr
func (b *BaseCmd) SetSysProcAttr(value *syscall.SysProcAttr) {
	b.SysProcAttr = value
}

// GetPath return the Cmd's Path
func (b *BaseCmd) GetPath() string {
	return b.Path
}

// GetArgs return the Cmd's Args
func (b *BaseCmd) GetArgs() []string {
	return b.Args
}

// GetEnv return the Cmd's Env
func (b *BaseCmd) GetEnv() []string {
	return b.Env
}

// GetDir return the Cmd's Dir
func (b *BaseCmd) GetDir() string {
	return b.Dir
}

// GetStdin return the Cmd's Stdin
func (b *BaseCmd) GetStdin() io.Reader {
	return b.Stdin
}

// GetStdout return the Cmd's Stdout
func (b *BaseCmd) GetStdout() io.Writer {
	return b.Stdout
}

// GetStderr return the Cmd's Stderr
func (b *BaseCmd) GetStderr() io.Writer {
	return b.Stderr
}

// GetExtraFiles return the Cmd's ExtraFiles
func (b *BaseCmd) GetExtraFiles() []*os.File {
	return b.ExtraFiles
}

// GetSysProcAttr return the Cmd's SysProcAttr
func (b *BaseCmd) GetSysProcAttr() *syscall.SysProcAttr {
	return b.SysProcAttr
}

// GetProcess return the Cmd's Process
func (b *BaseCmd) GetProcess() *os.Process {
	return b.Process
}

// GetProcessState return the Cmd's ProcessState
func (b *BaseCmd) GetProcessState() *os.ProcessState {
	return b.ProcessState
}
