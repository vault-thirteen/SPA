package cla

import (
	"fmt"
	"strings"
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
	ErrUnknownAction = "unknown action: %v"
)

type ActionId = byte

func NewActionId(actionStr string) (actionId ActionId, err error) {
	switch strings.ToLower(actionStr) {
	case ActionStrHash:
		return ActionIdHash, nil
	case ActionStrCheck:
		return ActionIdCheck, nil
	default:
		return 0, fmt.Errorf(ErrUnknownAction, actionStr)
	}
}
