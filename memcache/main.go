package main

import (
	"bufio"
	"fmt"
	"golang_homework/memcache/common"
	"golang_homework/memcache/handlers"
	"net"
	"sync"
)

var cache = sync.Map{}

var supportedHandlers = []common.CommandHandler{
	&handlers.SetCommandHandler{&cache},
	&handlers.GetCommandHandler{&cache},
	&handlers.DelCommandHandler{&cache, sync.Mutex{}},
	&handlers.ExistsCommandHandler{&cache},
	&handlers.IncrbyCommandHandler{&cache, sync.Mutex{}},
	&handlers.DecrbyCommandHandler{&cache, sync.Mutex{}},
}

const Port = "9999"

func main() {
	ln, _ := net.Listen("tcp", ":"+Port)
	fmt.Println("Запущено на localhost:" + Port)
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("---")
		fmt.Print("[<] " + message)
		result := handleMessage(message)
		fmt.Println("[>] " + result)
		conn.Write([]byte(result + "\n"))
	}
}

func handleMessage(message string) string {
	cmd, err := common.ParseCommand(message)
	if err != nil {
		return err.Error()
	}
	for _, handler := range supportedHandlers {
		if handler.CanHandle(cmd) {
			return handler.HandleCommand(cmd)
		}
	}

	return "Неизвестная команда"
}
