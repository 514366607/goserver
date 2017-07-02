package cg

import (
	"errors"
	"fmt"
	"gameserver/ipc"
)

type Room struct {
	RoomName string
	password string
	players  []*Player
	mq       chan *ipc.Message //聊天内容
}

func (room *Room) PlayerList() []*Player {
	return room.players
}

func (room *Room) PlayerName() []string {
	var slice_PlayerName = make([]string, 0)

	for _, p := range room.players {
		slice_PlayerName = append(slice_PlayerName, p.Name)
	}

	return slice_PlayerName
}

func (room *Room) Join(player *Player, s_password string) error {
	if room.password != "" && s_password != room.password {
		return errors.New("房间密码错误")
	}

	for _, p := range room.players {
		if p == player {
			return errors.New("玩家已在房间中！不进行任何操作")
		}
	}
	room.players = append(room.players, player)
	return nil
}

func (room *Room) Left(player *Player) error {
	for i_key, slice_RoomPlayer := range room.players {
		if slice_RoomPlayer.Name == player.Name {
			if len(room.players) == 1 {
				//只有一条直接清掉
				room.players = make([]*Player, 0)
			} else if i_key == len(room.players)-1 {
				//最后一条的情况
				room.players = room.players[:i_key]
			} else if i_key == 0 {
				//去掉第一条
				room.players = room.players[1:]
			} else {
				//中间
				room.players = append(room.players[:i_key], room.players[i_key+1:]...)
			}
			return nil
		}
	}
	return errors.New("玩家不在房间中")
}

func NewRoom(s_roomName, s_password string) *Room {
	chan_Chat := make(chan *ipc.Message, 1024)
	var room = &Room{s_roomName, s_password, make([]*Player, 0), chan_Chat}

	go func(r *Room) {
		for {
			msg := <-r.mq
			fmt.Println("Room-", r.RoomName, "|", msg.From, "received message :", msg.Content)
		}
	}(room)

	Rooms = append(Rooms, room)

	return room
}

var Rooms []*Room
var GlobalRoom = NewRoom("global", "")
