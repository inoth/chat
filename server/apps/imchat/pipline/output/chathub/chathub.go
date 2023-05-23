package chathub

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/output"
	"fmt"
)

var (
	name = "chathub"
)

type ChatHubOutput struct{}

func init() {
	output.Add(name, func() imchat.MessageOutPut {
		return &ChatHubOutput{}
	})
}

func (cho *ChatHubOutput) Write(msg message.MessageBox) {
	fmt.Printf("这里准备往 chathub 服务直接写入消息: %v", string(msg.String()))
}
