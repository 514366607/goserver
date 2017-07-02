package cg

import (
	// "fmt"
	"gameserver/ipc"
	"testing"
)

var centerClient *CenterClient

func TestIpc(t *testing.T) {
	server := ipc.NewIpcServer(&CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &CenterClient{client}

	player := NewPlayer()
	player.Name = "test"
	player.Level = 1
	player.Exp = 1
	centerClient.AddPlayer(player)

	player2 := NewPlayer()
	player2.Name = "test2"
	player2.Level = 1
	player2.Exp = 1
	centerClient.AddPlayer(player2)

	i_Res, err := centerClient.ListPlayer("")
	if err != nil {
		t.Error("AddPlayer ERROR!", err)
	}

	err = centerClient.RemovePlayer("test")
	if err != nil {
		t.Error("Remove Player ERROR!", err)
	}

	i_Res, err = centerClient.ListPlayer("")
	if err != nil {
		t.Error("AddPlayer ERROR!", err)
	}

	if len(i_Res) != 1 {
		t.Error(" PlayerNumber != 1 ! ")
	}

	err = centerClient.RemovePlayer("test2")
	if err != nil {
		t.Error("Remove Player ERROR!", err)
	}

	// var client1 = NewIpcClient(server)
	// var client2 = NewIpcClient(server)

	// resp1, _ := client1.Call("foo", "From Client1")
	// resp2, _ := client2.Call("foo", "From Client2")
	// resp3, _ := client2.Call("foo2", "From Client2")
	// if resp1.Body != "ECHO: foo - From Client1" || resp2.Body != "ECHO: foo - From Client2" || resp3.Body == "" {
	// 	t.Error("IpcClient.Call failed. resp1:", resp1, "resp2", resp2)
	// }

	// client1.Close()
	// client2.Close()

}
