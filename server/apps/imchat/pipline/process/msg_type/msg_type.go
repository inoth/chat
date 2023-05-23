package msgtype

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/process"
	"fmt"
)

type MsgType struct {
}

func init() {
	process.Add("msgType", func() imchat.MessageProcess {
		return &MsgType{}
	})
}

func (jp *MsgType) Process(msg message.MessageBox) (message.MessageBox, error) {
	fmt.Println("这里处理消息中的类型")
	return msg, nil
}
