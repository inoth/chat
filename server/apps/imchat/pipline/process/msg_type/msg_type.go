package msgtype

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/process"
	"chat/toybox/components/logger"
	"fmt"
)

type MsgTypeProcess struct {
}

func init() {
	process.Add("msg_type", func() imchat.MessageProcess {
		return &MsgTypeProcess{}
	})
}

func (jp *MsgTypeProcess) Process(msg message.MessageBox) (message.MessageBox, error) {
	switch msg.MsgType() {
	case message.AuthMsg:
		m, _ := message.NewMsgBoxWithString("auth type is not yet supported", message.SysMsg, message.SysMsg, []string{msg.MsgSouce()})
		return m, nil
	case message.SysMsg:
		return handlerSysMsg(msg)
	case message.TextMsg:
		return handlerTextMsg(msg)
	default:
		errMsg := fmt.Sprintf("unknown message type %v", msg.MsgType())
		logger.Log.Errorf(errMsg)
		m, _ := message.NewMsgBoxWithString(errMsg, message.SysMsg, message.SysMsg, []string{msg.MsgSouce()})
		return m, nil
	}
}

func handlerTextMsg(msg message.MessageBox) (message.MessageBox, error) {
	return msg, nil
}

func handlerSysMsg(msg message.MessageBox) (message.MessageBox, error) {
	return msg, nil
}
