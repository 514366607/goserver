package cg

import (
	"encoding/json"
	"errors"
	"gameserver/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}

	resp, err := client.Call("addplayerplayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("RemovePlayer", name)
	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

func (client *CenterClient) listPlayer(params string) (ps []*Player, err error) {
	resp, _ := client.Call("listPlayer", params)
	if resp.Code != "200" {
		err = errors.New(resp.Code)
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)

	return
}

func (client *CenterClient) broadcast(message string) error {
	m := &ipc.Message{Content: message}

	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, _ := client.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}

	return errors.New(resp.Code)
}
