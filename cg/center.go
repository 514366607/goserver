package cg

import (
	"encoding/json"
	"errors"
	"gameserver/ipc"
	"sync"
)

//判断下是不是实现了CenterServer的接口
var _ ipc.Server = &CenterServer{}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

//添加用户
func (server *CenterServer) addPlayer(params string) error {
	Player := NewPlayer()

	err := json.Unmarshal([]byte(params), &Player)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.players = append(server.players, Player) //没有判断重复登陆的情况
	return nil
}

//删除用户
func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i, v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 {
				server.players = make([]*Player, 0)
			} else if i == len(server.players)-1 {
				server.players = server.players[:i]
			} else if i == 0 {
				server.players = server.players[i:]
			} else {
				server.players = append(server.players[:i-1], server.players[:i+1]...)
			}
			return nil
		}
	}
	return errors.New("Plsyer not found.")
}

//用户列表
func (server *CenterServer) listPlayer(params string) (players string, err error) {
	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("No player online.")
	}
	return
}

//广播
func (server *CenterServer) broadcast(params string) error {
	var message ipc.Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}

	if len(server.players) > 0 {
		for _, player := range server.players {
			player.mq <- &message
		}
	} else {
		err = errors.New("No player online.")
	}

	return err
}

func (server *CenterServer) Handle(method, params string) *ipc.Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{"200", players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200"}
	default:
		return &ipc.Response{Code: "404", Body: method + ":" + params}
	}
	return &ipc.Response{Code: "200"}
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)
	return &CenterServer{servers: servers, players: players}
}
