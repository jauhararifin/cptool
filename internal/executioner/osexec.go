package executioner

import osexec "os/exec"

// OsExec implemens exec from os package
type OsExec struct{}

// NewOSExec create new OsExec instance
func NewOSExec() Exec {
	return &OsExec{}
}

// Command returns the Cmd struct to execute the named program with the given arguments
func (*OsExec) Command(name string, arg ...string) Cmd {
	return osexec.Command(name, arg...)
}
