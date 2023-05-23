package imchat

import "chat/apps/imchat/message"

type MessageProcess interface {
	Process(msg message.MessageBox) (message.MessageBox, error)
}
