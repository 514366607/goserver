package cg

import (
	"fmt"
	"testing"
)

func TestRoom(t *testing.T) {

	player := NewPlayer()
	player.Name = "test"
	player.Level = 1
	player.Exp = 1
	// centerClient.AddPlayer(player)

	player2 := NewPlayer()
	player2.Name = "test2"
	player2.Level = 1
	player2.Exp = 1
	// centerClient.AddPlayer(player2)

	var room1 = NewRoom("房间1", "")
	var err = room1.Join(player, "")
	if err != nil {
		t.Error("玩家不应该在房间里")
	}

	err = room1.Join(player, "")
	if err == nil {
		t.Error("这里应该报玩家已经在房间里了")
	}

	err = room1.Join(player2, "")
	if err != nil {
		t.Error("玩家2进入房间不应该有任何的报错")
	}

	err = room1.Left(player2)
	if err != nil {
		t.Error("玩家2离开房间不应该有Error")
	}

	var roomPlayerName = room1.PlayerName()
	if roomPlayerName[0] != "test" {
		t.Error("test2离开后应该只有test在里面", roomPlayerName)
	}

	err = room1.Join(player2, "")
	if err != nil {
		t.Error("玩家2进入房间不应该有任何的报错")
	}

	err = room1.Left(player)
	if err != nil {
		t.Error("玩家离开房间不应该有Error")
	}

	roomPlayerName = room1.PlayerName()
	if roomPlayerName[0] != "test2" {
		t.Error("玩家1离开后应该只有test2玩家在里面", roomPlayerName)
	}

	err = room1.Join(player, "")
	if err != nil {
		t.Error("test2进入房间不应该有任何的报错")
	}

	player3 := NewPlayer()
	player3.Name = "test3"
	player3.Level = 1
	player3.Exp = 1
	err = room1.Join(player3, "")
	if err != nil {
		t.Error("test3进入房间不应该有任何的报错")
	}

	err = room1.Left(player)
	if err != nil {
		t.Error("玩家离开房间不应该有Error")
	}

	roomPlayerName = room1.PlayerName()
	if roomPlayerName[0] != "test2" || roomPlayerName[1] != "test3" {
		t.Error("test1离开后应该为[\"test2\" , \"test3\"]", roomPlayerName)
	}

	var room2 = NewRoom("房间2", "123123")
	err = room2.Join(player, "123123")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(Rooms)

}
