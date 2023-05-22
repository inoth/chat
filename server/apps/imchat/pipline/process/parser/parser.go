package parser

import (
	"chat/apps/imchat"
	"chat/apps/imchat/pipline/process"
)

type JsonParser struct {
}

func init() {
	process.Add("jsonParser", func() imchat.MessageProcess {
		return &JsonParser{}
	})
}

func (jp *JsonParser) Process() error {
	return nil
}
