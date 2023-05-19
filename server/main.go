package main

import (
	"chat/toybox"

	_ "chat/apps/all"
)

func main() {
	tb := toybox.New(
		toybox.WithCfgPath("config"),
		toybox.EnableComponents(
		// config.New(),
		// redis.New(),
		// mysql.New(),
		),
	)
	tb.Run()
}
