package handlers

import (
	"fmt"
	"golang_homework/memcache/common"
	"strconv"
	"sync"
)

type IncrbyCommandHandler struct {
	Cache *sync.Map
	sync.Mutex
}

func (handler *IncrbyCommandHandler) HandleCommand(command *common.Command) string {
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

	num++

	numStr := fmt.Sprintf("%v", num)
	handler.Cache.Store(command.Args[0], numStr)
	return "1\r" + numStr
}

func (handler *IncrbyCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "incrby" && len(command.Args) == 1
}
