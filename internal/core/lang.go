package core

import (
	"errors"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/spf13/afero"
)

// Language defines programming language used for competitive programming like c, c++, java, etc. VerboseName contains string
// that will displayed to the user. Name contains string that unique for every language, this identifies the language. Extension
// is the extension of language, every solution that written using this language must have same extension as this (like .pas for
// Pascal language).
//
// Every programming language must have compile script and run script, these script contains how to compile and run the solution.
// Bash script can be used to compile the solution and run the compiled solution. CompileScript contains path to the compile script
// location. RunScript contains path to the script for running the compiled solution.
// Some programming languages are debuggable. When the language is debuggable, it Debuggable property will set to true and
// DebugScript contains path to compile script in debug mode.
//
// In CPTool, A language is defined using a folder, the folder name identifies the language name. Below are example of folder structure
// of C language
//
//     c/
//     ├── compile
//     ├── debugcompile
//     ├── lang.conf
//     └── run
//
// compile file contains script for compiling program. When the script executed, it will receive two arguments. The first one is the
// path to the solution's source code and the second one is the path where the compiled program should be located. Below is the example
// of compile script for c language
//
//     #!/bin/bash
//     SOURCE=$1
//     DEST=$2
//     gcc -x c -Wall -O2 -static -pipe -o "$DEST" "$SOURCE" -lm
//     exit $?
//
// File debugcompile contains script for compiling program in debug mode. It receive the same arguments as compile script when executed
// Below is the example of compile script for c language in debug mode.
//
//     #!/bin/bash
//     SOURCE=$1
//     DEST=$2
//     gcc -x c -Wall -O2 -static -pipe -o "$DEST" "$SOURCE" -lm -g
//     exit $?
//
// File lang.conf contains other information about language (VerboseName and Extension). Below the example for c language:
//
//     verbose_name=C
//     extension=c
//
// File run contains script for executing compiled program. When the script is executes, it will receive one argument that defines the
// location of compiled program. Below is example for c language:
//
//     #!/bin/bash
//     PROGRAM=$1
//     ./$PROGRAM
//     exit $?
//
type Language struct {
	Name        string
	Extension   string
	VerboseName string

	CompileScript string
	RunScript     string
	DebugScript   string
	Debuggable    bool
}

// ErrInvalidLanguageDirectory indicates directory is not a valid language definition.
var ErrInvalidLanguageDirectory = errors.New("Invalid language directory")

// ErrInvalidLanguageConfigurationFile indicates language configuration file has invalid format.
var ErrInvalidLanguageConfigurationFile = errors.New("Invalid language configuration file")

// ErrNoSuchLanguage indicates no language found.
var ErrNoSuchLanguage = errors.New("No such language")

// GetAllLanguages returns all known language as a pair of []Language, map[string]Language. The first element
// of pair contains array of all known languages. The second element of pair contains map of Language of string
// that map between language's name and itself.
func (cptool *CPTool) GetAllLanguages() ([]Language, map[string]Language) {
	languages := make([]Language, 0)
	for _, value := range cptool.languages {
		languages = append(languages, value)
	}
	return languages, cptool.languages
}

// GetLanguageByName returns language that has specific name. ErrNoSuchLanguage will returned when no such language exists.
func (cptool *CPTool) GetLanguageByName(name string) (Language, error) {
	if lang, ok := cptool.languages[name]; ok {
		return lang, nil
	}
	return Language{}, ErrNoSuchLanguage
}

// GetDefaultLanguage returns default language. The default language is defined in "config" file in some of configuration path.
// Can be "/etc/cptool/config", "~/.cptool/config", "$CPTOOL_HOME/config" or ".cptool/config" in your current working directory.
// Below is example of config file that defines default language as C Plus Plus ("cpp" is the language's name, while "C Plus Plus"
// is the language verbose name).
//
//     default_language=cpp
//
// When there is no valid config file, the default language is choosen between all known languages. When there is no known language,
// then ErrNoSuchLanguage error returned.
func (cptool *CPTool) GetDefaultLanguage() (Language, error) {
	configurationPaths := cptool.GetConfigurationPaths()
	for _, confPath := range configurationPaths {
		userConfigPath := path.Join(confPath, "config")
		info, err := cptool.fs.Stat(userConfigPath)
		if err != nil || info.IsDir() {
			continue
		}

		if userConfigFile, err := cptool.fs.Open(userConfigPath); err == nil {
			result := &struct {
				DefaultLanguage string `toml:"default_language"`
			}{}
			if _, err = toml.DecodeReader(userConfigFile, &result); err == nil {
				if len(result.DefaultLanguage) > 0 {
					if defaultLanguage, err := cptool.GetLanguageByName(result.DefaultLanguage); err == nil {
						return defaultLanguage, nil
					}
				}
			}
		}
	}

	languages, _ := cptool.GetAllLanguages()
	if len(languages) == 0 {
		return Language{}, ErrNoSuchLanguage
	}
	return languages[0], nil
}

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

func (cptool *CPTool) getLanguageFromDirectory(languagePath string) (Language, error) {
	info, err := cptool.fs.Stat(languagePath)
	if err != nil || !info.IsDir() {
		return Language{}, ErrInvalidLanguageDirectory
	}

	language := Language{}
	language.Name = info.Name()
	language.VerboseName = language.Name
	language.Extension = language.Name

	configPath := path.Join(languagePath, "lang.conf")
	info, err = cptool.fs.Stat(configPath)
	if err == nil && !info.IsDir() {
		configFile, _ := cptool.fs.Open(configPath)
		languageConf := languageConfFile{}
		if _, err = toml.DecodeReader(configFile, &languageConf); err != nil {
			return Language{}, ErrInvalidLanguageConfigurationFile
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
		return Language{}, err
	}
	if info.IsDir() {
		return Language{}, ErrInvalidLanguageDirectory
	}

	language.RunScript = path.Join(languagePath, "run")
	info, err = cptool.fs.Stat(language.RunScript)
	if err != nil {
		return Language{}, err
	}
	if info.IsDir() {
		return Language{}, ErrInvalidLanguageDirectory
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
					cptool.languages[lang.Name] = lang
				}
			}
			return nil
		})
	}
}
