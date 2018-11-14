package handlers

import (
	"golang_homework/memcache/common"
)

type GetCommandHandler struct {
	Cache *common.Cache
}

func (handler *GetCommandHandler) HandleCommand(command *common.Command) string {
	handler.Cache.RLock()
	defer handler.Cache.RUnlock()

	item, ok := handler.Cache.Items[command.Args[0]]
	if !ok {
		return "0"
	}
	return "1\n" + item
}

func (handler *GetCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "get" && len(command.Args) == 1
}
