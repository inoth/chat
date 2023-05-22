package msgtype

import (
	"chat/apps/imchat"
	"chat/apps/imchat/pipline/process"
)

type MsgType struct {
}

func init() {
	process.Add("msgType", func() imchat.MessageProcess {
		return &MsgType{}
	})
}

func (jp *MsgType) Process() error {
	return nil
}
