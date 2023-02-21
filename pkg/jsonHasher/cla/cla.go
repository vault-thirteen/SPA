package cla

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const UsageHint = `Usage:
	tool.exe [Action] [File]

Notes:
	[Action] is action type: 'hash' or 'check',
	[File] is a path to a file to be processed,
	'tool.exe' is here just for reference.

Example:
	tool.exe hash my.json
	tool.exe check my.json`

const (
	ErrSyntax = "syntax error"
)

type CommandLineArguments struct {
	ActionId      ActionId
	FileToProcess string
}

func ReadCLA() (cla *CommandLineArguments, err error) {
	if len(os.Args) != 3 {
		return nil, errors.New(ErrSyntax)
	}

	cla = &CommandLineArguments{}

	cla.ActionId, err = NewActionId(strings.TrimSpace(os.Args[1]))
	if err != nil {
		return nil, err
	}

	cla.FileToProcess = strings.TrimSpace(os.Args[2])

	return cla, nil
}

func ShowUsage() {
	fmt.Println(UsageHint)
}
