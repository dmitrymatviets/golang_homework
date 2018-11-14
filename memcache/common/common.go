package common

import (
	"fmt"
	"strings"
	"sync"
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

type Cache struct {
	sync.RWMutex
	Items map[string]string
}
