package test

import (
	_ "chat/apps/all"

	"chat/toybox"
	"chat/toybox/components/config"
	"chat/toybox/components/mysql"
	"chat/toybox/components/redis"
	"fmt"
	"os"
	"testing"
)

func TestNewToyBox(t *testing.T) {
	os.Setenv("GORUNEVN", "dev")

	tb := toybox.New(
		toybox.WithCfgPath("../config"),
		toybox.EnableComponents(config.New(), redis.New(), mysql.New()),
	)

	tb.Run()

	fmt.Printf("%+v", tb)
}
