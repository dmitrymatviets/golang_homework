package common

import (
	"fmt"
	"strings"
)

type Command struct {
	Operator string
	Args     []string
}

func ParseCommand(s string) (*Command, error) {
	split := strings.Fields(s)
	if len(split) == 0 {
		return nil, fmt.Errorf("Пустая команда")
	}
	return &Command{strings.ToLower(split[0]), split[1:]}, nil
}

type CommandHandler interface {
	HandleCommand(command *Command) string
	CanHandle(command *Command) bool
}
