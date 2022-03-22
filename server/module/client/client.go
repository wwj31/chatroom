package client

import (
	"chatroom/protocol/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
	"time"
)

const addr = "127.0.0.1:8888"

func New() *Client {
	return new(Client)
}

type Client struct {
	actor.Base
	cli  network.Client
	room string
	name string
}

func (s *Client) OnInit() {
	s.cli = network.NewTcpClient(addr, func() network.DecodeEncoder { return &network.StreamCode{} })
	s.cli.AddLast(func() network.NetSessionHandler { return &SessionHandler{client: s} })
	_ = s.cli.Start(false)
	s.InitCmd()
	// 心跳
	s.AddTimer(tools.XUID(), tools.NowTime()+int64(10*time.Second), func(dt int64) {
		s.SendToServer(msg.CHATROOM_PING.Int32(), &msg.Ping{ClientTime: time.Now().UnixMilli()})
	}, -1)

	fmt.Println("连接成功!输入 h 查看帮助命令")
}

func (s *Client) SendToServer(msgId int32, pb proto.Message) {
	bytes, err := proto.Marshal(pb)
	expect.Nil(err)

	data := network.CombineMsgWithId(msgId, bytes)
	err = s.cli.SendMsg(data)
	expect.Nil(err)
}

func (s *Client) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch resp := v.(type) {
	case *msg.Pong:
		//fmt.Println("aliving...")
	case *msg.LoginResp:
		fmt.Println("登录进入房间:", resp.String())
		s.room = resp.Room
	case *msg.BroadcastMsg:
		for _, chatmsg := range resp.Msgs {
			t := time.UnixMilli(chatmsg.CreateAt / int64(time.Millisecond))
			fmt.Println(fmt.Sprintf("%v [%v]:%v ", t.Format("2006-01-02 15:04:05"), chatmsg.Name, chatmsg.ChatText))
		}
	case *msg.ChatResp:
		//fmt.Println("聊天发送成功 code:", resp.String())
	}
}
