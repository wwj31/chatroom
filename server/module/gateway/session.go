package gateway

import (
	"chatroom/protocol/msg"
	"chatroom/server/common"
	"chatroom/server/common/log"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
	"time"
)

type UserSession struct {
	network.NetSession
	gateway *GateWay

	// 玩家所在房间
	room string

	// 链接保活时间
	LeaseTime int64
}

// OnSessionCreated 建立链接
func (s *UserSession) OnSessionCreated(sess network.NetSession) {
	s.NetSession = sess
	s.LeaseTime = time.Now().UnixMilli()

	// 这里只做session映射，等待客户端请求登录
	_ = s.gateway.Send(s.gateway.ID(), func() {
		s.gateway.sessions[s.Id()] = s
	})
}

// OnSessionClosed 链接关闭
func (s *UserSession) OnSessionClosed() {
	if s.room != "" {
		// 连接断开，通知房间
		session := common.GateSession(s.gateway.ID(), s.Id())
		brokenNotify := &msg.SessionBroken{Session: session}
		_ = s.gateway.Send(s.room, brokenNotify)
		_ = s.gateway.Send(common.Login, brokenNotify)
	}

	_ = s.gateway.Send(s.gateway.ID(), func() {
		delete(s.gateway.sessions, s.Id())
	})
}

func (s *UserSession) OnRecv(data []byte) {
	if len(data) < 4 {
		log.Warnw("invalid data len", "len(data)", len(data), "session", s.Id())
	}

	msgId := int32(network.Byte4ToUint32(data[:4]))

	var err error
	defer func() {
		if err != nil {
			log.Errorw("OnRecv error", "err", err, "msgId", msgId)
		}
	}()

	protoIndex := s.gateway.System().ProtoIndex()
	// 心跳
	if msgId == msg.CHATROOM_PING.Int32() {
		ping := network.NewBytesMessageParse(data, protoIndex).Proto().(*msg.Ping)

		pong := network.NewPbMessage(&msg.Pong{
			ClientTime: ping.ClientTime,
			ServerTime: tools.Milliseconds(),
		}, msg.CHATROOM_PONG.Int32())

		err = s.SendMsg(pong.Buffer())
		s.LeaseTime = time.Now().UnixMilli()
		return
	}

	msgName, ok := protoIndex.MsgIdToName(msgId)
	if !ok {
		log.Errorw("proto not find struct", "msgId", msgId)
		return
	}
	gSession := common.GateSession(s.gateway.ID(), s.Id())
	wrapperMsg := common.NewGateWrapperByBytes(data[4:], msgName, gSession)

	// 消息派发至不同服务组件
	if msg.CHATROOM_LOGIN_REQ.Int32() == msgId {
		err = s.gateway.Send(common.Login, wrapperMsg)
	} else {
		err = s.gateway.Send(s.room, wrapperMsg)
	}
}
