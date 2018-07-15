package executioner

import osexec "os/exec"

type OsExec struct{}

// NewOSExec not yet defined
func NewOSExec() Exec {
	return OsExec{}
}

// Command returns the Cmd struct to execute the named program with the given arguments
func (OsExec) Command(name string, arg ...string) Cmd {
	return osexec.Command(name, arg...)
}
