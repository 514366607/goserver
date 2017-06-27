package cg

import (
// "sync"
)

type Room struct {
	RoomsId int
	Players []*Player
}

var Rooms []*Room

// func NewRoom(player *Player) Rooms {
// var RoomId int

// sync.RWMutex.Lock()

// RoomId = count(Rooms)
// Rooms = append(Rooms, &Room{RoomId, []int{player}})

// defer sync.RWMutex.Unlock()
// }
