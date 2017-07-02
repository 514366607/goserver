package cg

import (
	"fmt"
	"gameserver/ipc"
)

type Player struct {
	Name  string
	Level int
	Exp   int
	mq    chan *ipc.Message //等待收取的消息
}

func NewPlayer() *Player {
	m := make(chan *ipc.Message, 1024)
	player := &Player{"", 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "received message :", msg.Content)
		}
	}(player)

	return player
}
