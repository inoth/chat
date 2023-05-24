package imchat

import "chat/apps/imchat/message"

type MessageOutPut interface {
	Name() string
	Write(msg message.MessageBox)
}
