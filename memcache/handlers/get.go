package handlers

import (
	"golang_homework/memcache/common"
	"sync"
)

type GetCommandHandler struct {
	Cache *sync.Map
}

func (handler *GetCommandHandler) HandleCommand(command *common.Command) string {
	item, ok := handler.Cache.Load(command.Args[0])
	if !ok || (item == nil) {
		return "0"
	}
	return "1\n" + item.(string)
}

func (handler *GetCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "get" && len(command.Args) == 1
}
