package executioner

import osexec "os/exec"

// OsExec implements exec from os package
type OsExec struct{}

// OsCmd implements cmd from os package
type OsCmd struct {
	BaseCmd
}

// NewOSExec create new OsExec instance
func NewOSExec() Exec {
	return &OsExec{}
}

// Command returns the Cmd struct to execute the named program with the given arguments
func (*OsExec) Command(name string, arg ...string) Cmd {
	oscmd := osexec.Command(name, arg...)
	base := BaseCmd{oscmd}
	return &OsCmd{base}
}
