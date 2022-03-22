package login

import (
	"fmt"
	"reflect"
	"time"

	"chatroom/protocol/msg"
	"chatroom/server/common"
	"chatroom/server/common/log"
	"github.com/wwj31/dogactor/actor"
)

func New() *Login {
	return new(Login)
}

type Login struct {
	actor.Base
	users map[string]*userInfo // 所有玩家 map[name]session
}

type userInfo struct {
	session string // 网络链接
	roomId  string // 所在房间
	loginAt int64  // 登录时间
}

func (s *Login) OnInit() {
	s.users = make(map[string]*userInfo, 10)
	s.InitCmd()
}

// OnHandleMessage 处理消息
func (s *Login) OnHandleMessage(sourceId, targetId string, data interface{}) {
	v, _, session, err := common.UnwrapperGateMsg(data)
	if err != nil {
		log.Errorw("login unwrapper failed", "err", err)
		return
	}

	var resp *msg.LoginResp
	defer func() {
		if resp != nil {
			_ = s.Send(common.Gateway, common.NewGateWrapperByPb(resp, session))
		}
	}()

	switch req := v.(type) {
	case *msg.LoginReq:
		resp = s.login(req, session)
	case *msg.SessionBroken:
		for name, info := range s.users {
			if info.session == req.Session {
				delete(s.users, name)
			}
		}
	default:
		panic(fmt.Sprintf("unknown msg type %v", reflect.TypeOf(req).String()))
	}
}

func (s *Login) login(req *msg.LoginReq, session string) *msg.LoginResp {
	if info, found := s.users[req.Name]; found {
		if info.session != session {
			return &msg.LoginResp{ErrCode: msg.ERROR_REPEATED_NAME}
		}
		_ = s.Send(info.roomId, &msg.SessionBroken{Session: session})
	}

	// 考虑功能完整性和代码可读性，这里使用同步的通信机制，
	// 真实环境应改用异步的通信机制，提高login服务吞吐量
	result, erro := s.RequestWait(req.Room, &msg.UserEnterRoomReq{
		Name:    req.Name,
		Session: session,
	})

	if erro != nil {
		log.Errorw("request user enter room failed", "err", erro)
		return &msg.LoginResp{ErrCode: msg.ERROR_UNKNOWN_ERROR}
	}

	enterRoomResp, ok := result.(*msg.UserEnterRoomResp)
	if !ok {
		log.Errorw("assert UserEnterRoomResp failed", "type", reflect.TypeOf(result).String())
		return &msg.LoginResp{ErrCode: msg.ERROR_UNKNOWN_ERROR}
	}

	if _, ok := s.users[req.Name]; !ok {
		s.users[req.Name] = &userInfo{
			session: session,
			roomId:  req.Room,
		}
	}
	info := s.users[req.Name]
	info.loginAt = time.Now().UnixMilli()

	_ = s.Send(common.Gateway, &msg.EnterRoomNotifyGateway{Room: req.Room, Session: session})

	return &msg.LoginResp{
		Room:    req.Room,
		History: enterRoomResp.History,
	}
}
