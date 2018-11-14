package handlers

import (
	"golang_homework/memcache/common"
)

type DelCommandHandler struct {
	Cache *common.Cache
}

func (handler *DelCommandHandler) HandleCommand(command *common.Command) string {
	handler.Cache.Lock()
	defer handler.Cache.Unlock()

	var _, ok = handler.Cache.Items[command.Args[0]]
	if !ok {
		return "0"
	}
	delete(handler.Cache.Items, command.Args[0])
	return "1"
}

func (handler *DelCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "del" && len(command.Args) == 1
}
