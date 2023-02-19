package main

import (
	"errors"
	"os"
	"strings"
)

const ErrConfigurationFileIsNotSpecified = "path to a configuration file is not specified"

type CommandLineArguments struct {
	ConfigurationFilePath string
}

func readCLA() (cla *CommandLineArguments, err error) {
	if len(os.Args) != 2 {
		return nil, errors.New(ErrConfigurationFileIsNotSpecified)
	}

	cla = &CommandLineArguments{
		ConfigurationFilePath: strings.TrimSpace(os.Args[1]),
	}

	return cla, nil
}
