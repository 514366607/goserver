package ipc

import (
	"encoding/json"
	"fmt"
)

//服务器接口
type Server interface {
	Name() string
	Handle(method, params string) *Response
}

//服务器结构体，继承接口
type IpcServer struct {
	Server
}

//初始化Ipc服务器
func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

//连接方法
func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			s_Request := <-c
			if s_Request == "CLOSE" {
				//关闭连接
				break
			}

			var req Request
			var err = json.Unmarshal([]byte(s_Request), &req)
			if err != nil {
				fmt.Println("Invalid request format : ", s_Request)
			}
			resp := server.Handle(req.Method, req.Params)

			b, err := json.Marshal(resp)

			c <- string(b)
		}
		fmt.Println("Session Closed.")
	}(session)

	fmt.Println("A New session has been create ed successfully.")
	return session
}
