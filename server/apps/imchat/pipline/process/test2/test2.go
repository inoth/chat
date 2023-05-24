package test2

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/process"
	"fmt"
)

type Test2Process struct {
}

func init() {
	process.Add("process_test2", func() imchat.MessageProcess {
		return &Test2Process{}
	})
}

func (jp *Test2Process) Process(msg message.MessageBox) (message.MessageBox, error) {
	fmt.Println("2:process_test2")
	return msg, nil
}
