package core

import (
	"os"
	"os/user"

	"github.com/jauhararifin/cptool/internal/executioner"
	"github.com/jauhararifin/cptool/internal/logger"
	"github.com/spf13/afero"
)

// CPTool stores information about this tool. The information includes version, all known languages, and configuration directories.
type CPTool struct {
	languages map[string]Language

	exec executioner.Exec

	fs                  afero.Fs
	workingDirectory    string
	cptoolHomeDirectory string
	homeDirectory       string

	logger *logger.Logger
}

// New create new cptool instance. This instance contains working directory, cptool home directory, user home directory, and logger
func New(exec executioner.Exec, log *logger.Logger) (*CPTool, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	if exec == nil {
		exec = executioner.NewOSExec()
	}

	if log == nil {
		log = logger.New(os.Stderr, logger.INFO)
	}

	return &CPTool{
		languages: make(map[string]Language),

		exec: executioner.NewOSExec(),

		fs:                  afero.NewOsFs(),
		workingDirectory:    cwd,
		cptoolHomeDirectory: os.Getenv("CPTOOL_HOME"),
		homeDirectory:       user.HomeDir,

		logger: logger.New(os.Stderr, logger.INFO),
	}, nil
}

// Bootstrap will bootstrap cptool. The bootstrap process will load all language from known directories.
func (cptool *CPTool) Bootstrap() error {
	cptool.loadAllLanguages()
	return nil
}
