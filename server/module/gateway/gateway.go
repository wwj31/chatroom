package gateway

import (
	"chatroom/protocol/msg"
	"chatroom/server/common/log"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/network"
	"github.com/wwj31/dogactor/tools"
	"strconv"
	"strings"
	"time"
)

func New() *GateWay {
	return new(GateWay)
}

type GateWay struct {
	actor.Base

	// 管理所有对外的玩家tcp连接
	listener network.Listener
	sessions map[uint32]*UserSession

	// 消息映射表
	protoIndex *tools.ProtoIndex
}

func (s *GateWay) OnInit() {
	s.sessions = make(map[uint32]*UserSession)

	s.listener = network.StartTcpListen(
		Addr,
		func() network.DecodeEncoder { return &network.StreamCode{MaxDecode: 10 * 1024} },
		func() network.NetSessionHandler { return &UserSession{gateway: s} },
	)

	s.AddTimer(tools.XUID(), tools.NowTime()+int64(time.Hour), s.checkDeadSession, -1)

	if err := s.listener.Start(); err != nil {
		log.Errorw("gateway listener start failed", "err", err, "addr", Addr)
		return
	}
	log.Debugf("gateway OnInit")
}

// 定期检查并清理非活跃链接
func (s *GateWay) checkDeadSession(dt int64) {
	for id, session := range s.sessions {
		if time.Now().UnixMilli()-session.LeaseTime > int64(time.Hour) {
			session.Stop()
			delete(s.sessions, id)
			log.Warnw(" find dead session", "session", id)
		}
	}
}

// 所有消息，直接转发给用户
func (s *GateWay) OnHandleMessage(sourceId, targetId string, v interface{}) {
	switch message := v.(type) {
	case *msg.GateMsgWrapper:
		// 用户消息直接转发前
		sessionId := split(message.GateSession)
		userSessionHandler := s.sessions[sessionId]
		if userSessionHandler == nil {
			log.Warnw("cannot find sessionId", "sessionId", sessionId)
			return
		}
		msgId, _ := s.System().ProtoIndex().MsgNameToId(message.GetMsgName())
		_ = userSessionHandler.SendMsg(network.CombineMsgWithId(msgId, message.Data))

	case *msg.EnterRoomNotifyGateway:
		sessionId := split(message.Session)
		userSessionHandler := s.sessions[sessionId]
		if userSessionHandler == nil {
			log.Warnw("cannot find sessionId", "sessionId", sessionId)
			return
		}
		if user, found := s.sessions[sessionId]; found {
			user.room = message.Room
		}
	}
}

func split(gateSession string) uint32 {
	strs := strings.Split(gateSession, ":")
	sessionId, _ := strconv.Atoi(strs[1])
	return uint32(sessionId)
}
