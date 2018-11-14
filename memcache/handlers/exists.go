package handlers

import (
	"golang_homework/memcache/common"
	"sync"
)

type ExistsCommandHandler struct {
	Cache *sync.Map
}

func (handler *ExistsCommandHandler) HandleCommand(command *common.Command) string {
	var val, ok = handler.Cache.Load(command.Args[0])
	if !ok || (val == nil) {
		return "0"
	}
	return "1"
}

func (handler *ExistsCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "exists" && len(command.Args) == 1
}
