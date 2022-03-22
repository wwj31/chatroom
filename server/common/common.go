package common

import (
	"chatroom/protocol/msg"
	"chatroom/server/common/log"
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

func GateSession(gateId string, sessionId uint32) string {
	return fmt.Sprintf("%v:%v", gateId, sessionId)
}

// NewGateWrapperByBytes 封装网关透传消息
func NewGateWrapperByBytes(data []byte, msgName string, gateSession string) *msg.GateMsgWrapper {
	return &msg.GateMsgWrapper{GateSession: gateSession, MsgName: msgName, Data: data}
}

// UnwrapperGateMsg 解包网关消息
func UnwrapperGateMsg(v interface{}) (pb interface{}, msgName string, session string, err error) {
	wrapper, is := v.(*msg.GateMsgWrapper)
	if !is {
		return v, "", "", nil
	}

	pbMsg, ok := msg.Spawner(wrapper.MsgName, true)
	if !ok {
		return nil, "", wrapper.GateSession, fmt.Errorf("msg not found in outer Spawner %v", wrapper.MsgName)
	}
	actMsg := pbMsg.(proto.Message)

	err = proto.Unmarshal(wrapper.Data, actMsg)
	if err != nil {
		return nil, wrapper.MsgName, wrapper.GateSession, err
	}

	return actMsg, wrapper.MsgName, wrapper.GateSession, nil
}

// NewGateWrapperByPb 封装通过网关发送给客户端消息
func NewGateWrapperByPb(pb proto.Message, gateSession string) *msg.GateMsgWrapper {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorw("marshal pb failed", "err", err)
		return nil
	}

	name := reflect.TypeOf(pb).String()
	if name[0] == '*' {
		name = name[1:]
	}
	return &msg.GateMsgWrapper{GateSession: gateSession, MsgName: name, Data: data}
}
