package test

import (
	_ "chat/apps/all"

	"chat/components/redis"
	"chat/toybox"
	"fmt"
	"os"
	"testing"
)

func TestNewToyBox(t *testing.T) {
	os.Setenv("GORUNEVN", "dev")

	tb := toybox.New(
		toybox.WithCfgPath("../config"),
		toybox.EnableComponents(redis.New()),
	)

	tb.Run()

	fmt.Printf("%+v", tb)
}
