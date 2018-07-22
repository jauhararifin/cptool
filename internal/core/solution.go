package core

import (
	"errors"
	"path"
	"time"
)

// Solution store information about solution includes solution's name, language, location, and last updated.
type Solution struct {
	Name        string
	Language    Language
	Path        string
	LastUpdated time.Time
}

// ErrNoSuchSolution indicates that no solution exists
var ErrNoSuchSolution = errors.New("No such solution exists")

// GetSolution returns solution object with specific name and language. Solution is a single file that contains
// your source code that written in known programming language. A file named "example.cpp" in current working
// directory can be considered as a solution named "example", with language "cpp" (because cpp's extension is "cpp").
// If no solution with the specified name and language exists, then this method will return an ErrNoSuchSolution error.
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
