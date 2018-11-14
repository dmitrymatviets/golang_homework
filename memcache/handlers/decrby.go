package handlers

import (
	"fmt"
	"golang_homework/memcache/common"
	"strconv"
	"sync"
)

type DecrbyCommandHandler struct {
	Cache *sync.Map
	sync.Mutex
}

func (handler *DecrbyCommandHandler) HandleCommand(command *common.Command) string {
	handler.Lock()
	defer handler.Unlock()

	val, ok := handler.Cache.Load(command.Args[0])

	if val == nil || !ok {
		return "0"
	}

	num, err := strconv.Atoi(val.(string))

	if err != nil {
		return "0"
	}

	num--

	numStr := fmt.Sprintf("%v", num)
	handler.Cache.Store(command.Args[0], numStr)
	return "1\r" + numStr
}

func (handler *DecrbyCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "decrby" && len(command.Args) == 1
}
