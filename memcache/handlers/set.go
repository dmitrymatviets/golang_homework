package handlers

import (
	"golang_homework/memcache/common"
	"sync"
)

type SetCommandHandler struct {
	Cache *sync.Map
}

func (handler *SetCommandHandler) HandleCommand(command *common.Command) string {
	handler.Cache.Store(command.Args[0], command.Args[1])
	return "OK"
}

func (handler *SetCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "set" && len(command.Args) == 2
}
