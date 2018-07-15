package core

import (
	"errors"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/spf13/afero"
)

// Language defines p5cm per secondrogramming language used for competitive programming
type Language struct {
	Name        string
	Extension   string
	VerboseName string

	CompileScript string
	RunScript     string
	DebugScript   string
	Debuggable    bool
}

// ErrInvalidLanguageDirectory indicates directory is not a valid language definition
var ErrInvalidLanguageDirectory = errors.New("Invalid language directory")

// ErrInvalidLanguageConfigurationFile indicates language configuration file has invalid format
var ErrInvalidLanguageConfigurationFile = errors.New("Invalid language configuration file")

// ErrNoSuchLanguage indicates no language found
var ErrNoSuchLanguage = errors.New("No such language")

type languageConfFile struct {
	VerboseName string `toml:"verbose_name"`
	Extension   string `toml:"extension"`
}

func (cptool *CPTool) getLanguagesPaths() []string {
	configurationPaths := cptool.GetConfigurationPaths()
	langPaths := make([]string, 0)
	for _, confPath := range configurationPaths {
		langPaths = append(langPaths, path.Join(confPath, "langs"))
	}
	return langPaths
}

func (cptool *CPTool) getLanguageFromDirectory(languagePath string) (*Language, error) {
	info, err := cptool.fs.Stat(languagePath)
	if err != nil || !info.IsDir() {
		return nil, ErrInvalidLanguageDirectory
	}

	language := &Language{}
	language.Name = info.Name()
	language.VerboseName = language.Name
	language.Extension = language.Name

	configPath := path.Join(languagePath, "lang.conf")
	info, err = cptool.fs.Stat(configPath)
	if err == nil && !info.IsDir() {
		configFile, _ := cptool.fs.Open(configPath)
		languageConf := languageConfFile{}
		if _, err = toml.DecodeReader(configFile, &languageConf); err != nil {
			return nil, ErrInvalidLanguageConfigurationFile
		}
		if len(languageConf.VerboseName) > 0 {
			language.VerboseName = languageConf.VerboseName
		}
		if len(languageConf.Extension) > 0 {
			language.Extension = languageConf.Extension
		}
	}

	language.CompileScript = path.Join(languagePath, "compile")
	info, err = cptool.fs.Stat(language.CompileScript)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, ErrInvalidLanguageDirectory
	}

	language.RunScript = path.Join(languagePath, "run")
	info, err = cptool.fs.Stat(language.RunScript)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, ErrInvalidLanguageDirectory
	}

	DebugScript := path.Join(languagePath, "debugcompile")
	if info, err = cptool.fs.Stat(DebugScript); err == nil && !info.IsDir() {
		language.Debuggable = true
		language.DebugScript = DebugScript
	}

	return language, nil
}

func (cptool *CPTool) loadAllLanguages() {
	langPaths := cptool.getLanguagesPaths()
	for _, path := range langPaths {
		info, err := cptool.fs.Stat(path)
		if err != nil || !info.IsDir() {
			continue
		}

		afero.Walk(cptool.fs, path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				if lang, err := cptool.getLanguageFromDirectory(path); err == nil {
					if _, ok := cptool.languages[lang.Name]; ok {
						return nil
					}
					cptool.languages[lang.Name] = *lang
				}
			}
			return nil
		})
	}
}

// GetAllLanguages returns all known language as array of Language
func (cptool *CPTool) GetAllLanguages() ([]Language, map[string]Language) {
	languages := make([]Language, 0)
	for _, value := range cptool.languages {
		languages = append(languages, value)
	}
	return languages, cptool.languages
}

// GetLanguageByName returns language that has specific name
func (cptool *CPTool) GetLanguageByName(name string) (*Language, error) {
	if lang, ok := cptool.languages[name]; ok {
		return &lang, nil
	}
	return nil, ErrNoSuchLanguage
}

// GetLanguageByExtension returns language that has specific extension
func (cptool *CPTool) GetLanguageByExtension(extension string) []Language {
	results := make([]Language, 0)
	for _, lang := range cptool.languages {
		if lang.Extension == extension {
			results = append(results, lang)
		}
	}
	return results
}

// GetDefaultLanguage returns default language
func (cptool *CPTool) GetDefaultLanguage() (Language, error) {
	languages, _ := cptool.GetAllLanguages()
	if len(languages) == 0 {
		return Language{}, ErrNoSuchLanguage
	}
	return languages[0], nil
}

// GetDefaultLanguageForExtension returns default language with specific extension
func (cptool *CPTool) GetDefaultLanguageForExtension(extension string) (*Language, error) {
	languages, _ := cptool.GetAllLanguages()
	for _, lang := range languages {
		if lang.Extension == extension {
			return &lang, nil
		}
	}
	return nil, ErrNoSuchLanguage
}
