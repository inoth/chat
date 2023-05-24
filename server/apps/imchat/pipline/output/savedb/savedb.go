package savedb

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/output"
	"fmt"
)

var (
	name = "savedb"
)

type SavedbOutput struct{}

func init() {
	output.Add(name, func() imchat.MessageOutPut {
		return &SavedbOutput{}
	})
}

func (cho *SavedbOutput) Name() string { return name }

func (cho *SavedbOutput) Write(msg message.MessageBox) {
	fmt.Printf("这里准备往 mysql 数据库写入消息: %v\n", string(msg.String()))
}
