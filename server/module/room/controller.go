package room

import (
	"chatroom/protocol/msg"
	"chatroom/server/common"
	"chatroom/server/common/log"
	"chatroom/server/module/room/dirty"
	"time"
)

// ChatReq 用户请求聊天的消息及处理
func (s *Room) ChatReq(req *msg.ChatReq, session string) *msg.ChatResp {
	if _, found := s.users[session]; !found {
		log.Errorw("invaid session ", "session", session)
		return &msg.ChatResp{Err: msg.ERROR_UNKNOWN_ERROR}
	}

	chatmsg := &msg.ChatMsg{
		Name:     s.users[session],
		ChatText: dirty.CheckDirty(req.ChatText),
		CreateAt: time.Now().UnixNano(),
	}

	// 保留历史50条
	s.addHistory(chatmsg)

	// 追加聊天记录
	s.addSeqMsg(chatmsg)

	// 广播聊天消息
	for session, _ := range s.users {
		pb := common.NewGateWrapperByPb(&msg.BroadcastMsg{Msgs: []*msg.ChatMsg{chatmsg}}, session)
		_ = s.Send(common.Gateway, pb)
	}
	return &msg.ChatResp{}
}
