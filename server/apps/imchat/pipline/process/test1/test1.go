package test1

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/process"
	"fmt"
)

type Test1Process struct {
}

func init() {
	process.Add("process_test1", func() imchat.MessageProcess {
		return &Test1Process{}
	})
}

func (jp *Test1Process) Process(msg message.MessageBox) (message.MessageBox, error) {
	fmt.Println("1:process_test1")
	return msg, nil
}
