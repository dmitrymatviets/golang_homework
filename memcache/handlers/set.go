package handlers

import (
	"golang_homework/memcache/common"
)

type SetCommandHandler struct {
	Cache *common.Cache
}

func (handler *SetCommandHandler) HandleCommand(command *common.Command) string {
	handler.Cache.Lock()
	defer handler.Cache.Unlock()
	handler.Cache.Items[command.Args[0]] = command.Args[1]
	return "OK"
}

func (handler *SetCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "set" && len(command.Args) == 2
}
