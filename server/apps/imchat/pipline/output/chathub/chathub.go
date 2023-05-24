package chathub

import (
	"chat/apps/imchat"
	"chat/apps/imchat/core"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/output"
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

func (cho *ChatHubOutput) Name() string { return name }

func (cho *ChatHubOutput) Write(msg message.MessageBox) {
	core.SendMessage(msg)
}
