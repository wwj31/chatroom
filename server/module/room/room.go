package room

import (
	"chatroom/protocol/msg"
	"chatroom/server/common"
	"chatroom/server/common/log"
	"chatroom/server/module/room/dirty"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/wwj31/dogactor/actor"
	"reflect"
	"time"
)

// 抽象成为actor房间，每个房间独立管理自己的用户

const HISTORY_MAX = 50

func New() *Room {
	return new(Room)
}

type (
	Room struct {
		id string
		actor.Base
		users map[string]string // 房间内所有用户

		historyMsg []*msg.ChatMsg // 环形数组保存历史记录
		headPtr    int            // historyMsg偏移指针

		seqMsg []*msg.ChatMsg
	}
)

func (s *Room) OnInit() {
	s.users = make(map[string]string, 2)
	s.historyMsg = make([]*msg.ChatMsg, 0, HISTORY_MAX)
	s.seqMsg = make([]*msg.ChatMsg, 0, 2)
	s.headPtr = 0

	// 脏词表初始化
	dirty.Init()
}

// OnHandleMessage 处理消息
func (s *Room) OnHandleMessage(sourceId, targetId string, data interface{}) {
	v, _, session, err := common.UnwrapperGateMsg(data)
	if err != nil {
		log.Errorw("login unwrapper failed", "err", err)
		return
	}

	var resp proto.Message
	switch req := v.(type) {
	case *msg.ChatReq:
		resp = s.ChatReq(req, session)

	case *msg.SessionBroken:
		// 链接失效处理
		delete(s.users, req.Session)

	case string:
		s.printHF()
	}

	if resp != nil {
		_ = s.Send(common.Gateway, common.NewGateWrapperByPb(resp, session))
	}
}

// OnHandleRequest 处理请求
func (s *Room) OnHandleRequest(sourceId, targetId, requestId string, v interface{}) (respErr error) {
	switch req := v.(type) {
	case *msg.UserEnterRoomReq:
		_ = s.Response(requestId, s.enterRoom(req))
	default:
		return fmt.Errorf("unknown msg type %v", reflect.TypeOf(req).String())
	}
	return nil
}

// 进入房间逻辑
func (s *Room) enterRoom(req *msg.UserEnterRoomReq) *msg.UserEnterRoomResp {
	s.users[req.Session] = req.Name
	// todo... 如需展示房间内用户，此处做广播消息即可

	return &msg.UserEnterRoomResp{
		History: s.historyMsg,
	}
}

// 追加房间历史聊天记录
func (s *Room) addHistory(chatmsg *msg.ChatMsg) {
	if len(s.historyMsg) < HISTORY_MAX {
		s.historyMsg = append(s.historyMsg, chatmsg)
		return
	}

	nextHead := circleOffset(s.headPtr, HISTORY_MAX-1)
	s.historyMsg[s.headPtr] = chatmsg
	s.headPtr = nextHead
}

// 环形数组索引偏移
func circleOffset(c int, max int) (next int) {
	if c == max {
		return 0
	}
	return c + 1
}

// 追加聊天记录
func (s *Room) addSeqMsg(chatmsg *msg.ChatMsg) {
	if len(s.seqMsg) > 0 {
		idx := 0
		for i, c := range s.seqMsg {
			front := c
			if time.Now().UnixNano()-front.CreateAt >= int64(10*time.Minute) {
				idx = i
			} else {
				break
			}
		}
		s.seqMsg = s.seqMsg[idx:]
	}
	s.seqMsg = append(s.seqMsg, chatmsg)
}

// 打印出最近10分钟内发送频率最⾼的词
func (s *Room) printHF() {
	buckets := make(map[string]int, 2)
	var (
		maxCount int
		maxText  string
	)

	for _, v := range s.seqMsg {
		buckets[v.ChatText] = buckets[v.ChatText] + 1
		if buckets[v.ChatText] > maxCount {
			maxCount = buckets[v.ChatText]
			maxText = v.ChatText
		}
	}

	fmt.Println("近10分钟内发送频率最⾼的词:", maxText)
}
