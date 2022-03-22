package client

import (
	"chatroom/protocol/msg"
	"fmt"
)

func (s *Client) InitCmd() {
	s.RegistCmd("login", s.login, "login <name> <roomId> 登录并进入房间")
	s.RegistCmd("chat", s.chat, "chat <text> 发送聊天消息")
	s.RegistCmd("change", s.change, "change <roomId> 切换房间")

}

func (s *Client) login(arg ...string) {
	if len(arg) != 2 {
		fmt.Println("login 指令参数数量错误")
		return
	}

	s.name = arg[0]
	id := arg[1]

	logReq := &msg.LoginReq{
		Name: s.name,
		Room: "room" + id,
	}
	s.SendToServer(msg.CHATROOM_LOGIN_REQ.Int32(), logReq)
}

func (s *Client) chat(arg ...string) {
	if len(arg) == 0 {
		fmt.Println("chat 指令参数数量错误")
		return
	}

	var data string
	for _, c := range arg {
		data += c + " "
	}

	chatReq := &msg.ChatReq{
		ChatText: data,
	}
	s.SendToServer(msg.CHATROOM_CHAT_REQ.Int32(), chatReq)
}

func (s *Client) change(arg ...string) {
	if len(arg) != 1 {
		fmt.Println("chat 指令参数数量错误")
		return
	}

	roomId := arg[0]
	if roomId == s.room {
		return
	}
	s.login(s.name, roomId)
}
