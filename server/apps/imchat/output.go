package imchat

import "chat/apps/imchat/message"

type MessageOutPut interface {
	Write(msg message.MessageBox)
}
