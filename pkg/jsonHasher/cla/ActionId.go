package cla

import (
	"strings"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	ActionIdHash  = 1
	ActionIdCheck = 2
)

const (
	ActionStrHash  = "hash"
	ActionStrCheck = "check"
)

const (
	Errf_UnknownAction = "unknown action: %v"
)

type ActionId = byte

func NewActionId(actionStr string) (actionId ActionId, err error) {
	switch strings.ToLower(actionStr) {
	case ActionStrHash:
		return ActionIdHash, nil
	case ActionStrCheck:
		return ActionIdCheck, nil
	default:
		return 0, errors.ErrorFromFormattedText(Errf_UnknownAction, actionStr)
	}
}
