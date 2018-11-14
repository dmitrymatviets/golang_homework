package handlers

import (
	"fmt"
	"golang_homework/memcache/common"
	"strconv"
)

type DecrbyCommandHandler struct {
	Cache *common.Cache
}

func (handler *DecrbyCommandHandler) HandleCommand(command *common.Command) string {
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

	num--

	numStr := fmt.Sprintf("%v", num)
	handler.Cache.Items[command.Args[0]] = numStr
	return "1\r" + numStr
}

func (handler *DecrbyCommandHandler) CanHandle(command *common.Command) bool {
	return command.Operator == "decrby" && len(command.Args) == 1
}
