package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{"OK", "ECHO: " + method + " - " + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	var client1 = NewIpcClient(server)
	var client2 = NewIpcClient(server)

	resp1, _ := client1.Call("foo", "From Client1")
	resp2, _ := client2.Call("foo", "From Client2")
	resp3, _ := client2.Call("foo2", "From Client2")
	if resp1.Body != "ECHO: foo - From Client1" || resp2.Body != "ECHO: foo - From Client2" || resp3.Body == "" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2", resp2)
	}

	client1.Close()
	client2.Close()

}
