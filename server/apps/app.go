package apps

import "chat/toybox"

type Creator func() toybox.Server

var Servers map[string]Creator = make(map[string]Creator)

func Add(name string, creator Creator) {
	Servers[name] = creator
}
