package parser

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/process"
	"fmt"
)

type JsonParser struct {
}

func init() {
	process.Add("jsonParser", func() imchat.MessageProcess {
		return &JsonParser{}
	})
}

func (jp *JsonParser) Process(msg message.MessageBox) (message.MessageBox, error) {
	fmt.Println("这里来解析消息bytes")
	return msg, nil
}
