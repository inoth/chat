package main

import (
	"chat/toybox"
	"chat/toybox/components/config"
	"chat/toybox/components/logger/zap"
	"chat/toybox/components/mysql"

	_ "chat/apps/all"
)

func main() {
	toybox.New(
		toybox.WithComponentCfgPath("../config"),
		toybox.EnableComponents(
			zap.New(),
			config.New(),
			// redis.New(),
			mysql.New(),
		),
	).Run()
}
