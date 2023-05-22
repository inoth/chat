package output

import "chat/apps/imchat"

type Creator func() imchat.MessageOutPut

var Outputs map[string]Creator = make(map[string]Creator)

func Add(name string, creator Creator) {
	Outputs[name] = creator
}
