package handlers

import (
	"golang_homework/memcache/common"
	"sync"
)

type DelCommandHandler struct {
	Cache *sync.Map
	sync.Mutex
}

func (handler *DelCommandHandler) HandleCommand(command *common.Command) string {
	handler.Lock()
	defer handler.Unlock()

	var val, ok = handler.Cache.Load(command.Args[0])
	if !ok || (val == nil) {
		return "0"
	}
	handler.Cache.Delete(command.Args[0])
	return "1"
}

func (handler *DelCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "del" && len(command.Args) == 1
}
