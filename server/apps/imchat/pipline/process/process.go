package process

import "chat/apps/imchat"

type Creator func() imchat.MessageProcess

var Process map[string]Creator = make(map[string]Creator)

func Add(name string, creator Creator) {
	Process[name] = creator
}
