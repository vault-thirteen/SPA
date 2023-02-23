package cla

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const ConfigurationFilePathDefault = "settings.txt"

const UsageHint = `Usage:
	tool.exe [Category] [SettingsFile]

Notes:
	[Category] is the path of the category,
	[SettingsFile] is a path to the settings file,
	'tool.exe' is here just for reference.
	If the second parameter is omitted, the default value is used.
	Default settings file is 'settings.txt'.
	The 'news' category is the main index of the website.

Example:
	tool.exe soft settings.txt
	tool.exe news settings.txt
	tool.exe tech`

const (
	ErrSyntax = "syntax error"
)

type CommandLineArguments struct {
	Category              string
	ConfigurationFilePath string
}

func ReadCLA() (cla *CommandLineArguments, err error) {
	if len(os.Args) < 2 {
		return nil, errors.New(ErrSyntax)
	}

	cla = &CommandLineArguments{}

	cla.Category = strings.TrimSpace(strings.ToLower(os.Args[1]))

	if len(os.Args) == 2 {
		cla.ConfigurationFilePath = ConfigurationFilePathDefault
		return cla, nil
	}

	cla.ConfigurationFilePath = strings.TrimSpace(os.Args[2])

	return cla, nil
}

// IsDefaultFile tells whether the default file path is used for the
// configuration file.
func (cla *CommandLineArguments) IsDefaultFile() (isDefaultFile bool) {
	return cla.ConfigurationFilePath == ConfigurationFilePathDefault
}

func ShowUsage() {
	fmt.Println(UsageHint)
}
