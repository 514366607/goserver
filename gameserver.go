package main

import (
	"fmt"
	"gameserver/cg"
	"gameserver/ipc"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *ipc.Response {
	return &ipc.Response{"OK", "ECHO: " + method + " - " + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func main() {
	fmt.Println("I")

	server := ipc.NewIpcServer(&EchoServer{})
	fmt.Println(*server)
}
