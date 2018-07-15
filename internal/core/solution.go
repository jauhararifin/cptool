package core

import (
	"errors"
	"path"
	"time"
)

// Solution store information about solution
type Solution struct {
	Name        string
	Language    Language
	Path        string
	LastUpdated time.Time
}

// ErrNoSuchSolution indicates that no solution exists
var ErrNoSuchSolution = errors.New("No such solution exists")

// GetSolution returns solution object
func (cptool *CPTool) GetSolution(name string, language Language) (Solution, error) {
	solutionPath := path.Join(cptool.workingDirectory, name+"."+language.Extension)
	info, err := cptool.fs.Stat(solutionPath)
	if err != nil || info.IsDir() {
		return Solution{}, ErrNoSuchSolution
	}

	return Solution{
		Name:        name,
		Language:    language,
		Path:        solutionPath,
		LastUpdated: info.ModTime(),
	}, nil
}
