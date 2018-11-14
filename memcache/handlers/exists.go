package handlers

import (
	"golang_homework/memcache/common"
)

type ExistsCommandHandler struct {
	Cache *common.Cache
}

func (handler *ExistsCommandHandler) HandleCommand(command *common.Command) string {
	handler.Cache.RLock()
	defer handler.Cache.RUnlock()

	var _, ok = handler.Cache.Items[command.Args[0]]
	if !ok {
		return "0"
	}
	return "1"
}

func (handler *ExistsCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "exists" && len(command.Args) == 1
}
