package core

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// Language defines programming language used for competitive programming
type Language struct {
	Name        string
	Extension   string
	VerboseName string

	compileScript string
	runScript     string
	debugScript   string
	debuggable    bool
}

// ErrInvalidLanguageDirectory indicates directory is not a valid language definition
var ErrInvalidLanguageDirectory = errors.New("Invalid language directory")

// ErrInvalidLanguageConfigurationFile indicates language configuration file has invalid format
var ErrInvalidLanguageConfigurationFile = errors.New("Invalid language configuration file")

// ErrNoSuchLanguage indicates no language found
var ErrNoSuchLanguage = errors.New("No such language")

type languageConfFile struct {
	verboseName string `toml:"verbose_name"`
	extension   string `toml:"extension"`
}

func getLanguagesPaths() []string {
	return []string{"/etc/cptool/langs", "/home/jauhararifin/.cptool/langs"}
}

func checkDirExists(path string) bool {
	if info, err := os.Stat(path); err != nil || !info.IsDir() {
		return false
	}
	return true
}

func checkFileExists(filepath string) bool {
	if info, err := os.Stat(filepath); err != nil || info.IsDir() {
		return false
	}
	return true
}

// GetLanguageFromDirectory extract language information from specific directory
func GetLanguageFromDirectory(languagePath string) (*Language, error) {
	info, err := os.Stat(languagePath)
	if err != nil || info.IsDir() {
		return nil, ErrInvalidLanguageDirectory
	}

	language := &Language{}
	language.Name = info.Name()
	language.VerboseName = language.Name
	language.Extension = language.Name

	configPath := path.Join(languagePath, "lang.conf")
	if checkFileExists(configPath) {
		languageConf := languageConfFile{}
		if _, err = toml.Decode(configPath, &languageConf); err != nil {
			return nil, ErrInvalidLanguageConfigurationFile
		}
		if len(languageConf.verboseName) == 0 {
			languageConf.verboseName = language.Name
		}
		if len(languageConf.extension) == 0 {
			languageConf.extension = language.Name
		}
	}

	language.compileScript = path.Join(languagePath, "compile")
	if !checkFileExists(language.compileScript) {
		return nil, ErrInvalidLanguageDirectory
	}

	language.runScript = path.Join(languagePath, "run")
	if !checkFileExists(language.runScript) {
		return nil, ErrInvalidLanguageDirectory
	}

	debugScript := path.Join(languagePath, "debugcompile")
	if checkFileExists(debugScript) {
		language.debuggable = true
		language.debugScript = debugScript
	}

	return language, nil
}

// GetAllLanguages returns all known language as array of Language
func GetAllLanguages() ([]Language, map[string]Language) {
	langPaths := getLanguagesPaths()
	langMap := make(map[string]Language)
	for _, path := range langPaths {
		if !checkDirExists(path) {
			continue
		}

		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if lang, err := GetLanguageFromDirectory(path); err == nil {
					langMap[lang.Name] = *lang
				}
			}
			return nil
		})
	}
	languages := make([]Language, 0)
	for _, value := range langMap {
		languages = append(languages, value)
	}
	return languages, langMap
}

// GetLanguageByName returns language that has specific name
func GetLanguageByName(name string) (*Language, error) {
	_, langMap := GetAllLanguages()
	if lang, ok := langMap[name]; ok {
		return &lang, nil
	}
	return nil, ErrNoSuchLanguage
}

// GetLanguageByExtension returns language that has specific extension
func GetLanguageByExtension(extension string) []Language {
	languages, _ := GetAllLanguages()
	results := make([]Language, 0)
	for _, lang := range languages {
		if lang.Extension == extension {
			results = append(results, lang)
		}
	}
	return results
}

// GetDefaultLanguage returns default language
func GetDefaultLanguage() (*Language, error) {
	languages, _ := GetAllLanguages()
	if len(languages) == 0 {
		return nil, ErrNoSuchLanguage
	}
	return &languages[0], nil
}

// GetDefaultLanguageForExtension returns default language with specific extension
func GetDefaultLanguageForExtension(extension string) (*Language, error) {
	languages, _ := GetAllLanguages()
	for _, lang := range languages {
		if lang.Extension == extension {
			return &lang, nil
		}
	}
	return nil, ErrNoSuchLanguage
}
