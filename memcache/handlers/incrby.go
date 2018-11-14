package handlers

import (
	"fmt"
	"golang_homework/memcache/common"
	"strconv"
)

type IncrbyCommandHandler struct {
	Cache *common.Cache
}

func (handler *IncrbyCommandHandler) HandleCommand(command *common.Command) string {
	handler.Cache.Lock()
	defer handler.Cache.Unlock()

	val, ok := handler.Cache.Items[command.Args[0]]

	if !ok {
		return "0"
	}

	num, err := strconv.Atoi(val)

	if err != nil {
		return "0"
	}

	num++

	numStr := fmt.Sprintf("%v", num)
	handler.Cache.Items[command.Args[0]] = numStr
	return "1\r" + numStr
}

func (handler *IncrbyCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "incrby" && len(command.Args) == 1
}
