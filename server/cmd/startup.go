package main

import (
	"chatroom/protocol/msg"
	"chatroom/server/common"
	"chatroom/server/module/client"
	"chatroom/server/module/gateway"
	"chatroom/server/module/login"
	"chatroom/server/module/room"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/actor/cmd"
	"github.com/wwj31/dogactor/tools"
)

func startup() *actor.System {
	system, _ := actor.NewSystem(actor.ProtoIndex(newProtoIndex()), actor.WithCMD(cmd.New()))

	if *appId == "client" {
		_ = system.Add(actor.New(*appId, client.New()))
	} else {
		_ = system.Add(actor.New(common.Gateway, gateway.New()))
		_ = system.Add(actor.New(common.Login, login.New()))
		_ = system.Add(actor.New(common.Room1, room.New()))
		_ = system.Add(actor.New(common.Room2, room.New()))
	}
	return system
}

func newProtoIndex() *tools.ProtoIndex {
	return tools.NewProtoIndex(func(name string) (v interface{}, ok bool) {
		return msg.Spawner(name)
	}, tools.EnumIdx{
		PackageName: "msg",
		Enum2Name:   msg.CHATROOM_name,
	})
}
