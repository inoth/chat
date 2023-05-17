package test

import (
	"chat/components/redis"
	"chat/toybox"
	"fmt"
	"testing"
)

func TestNewToyBox(t *testing.T) {
	tb := toybox.New(
		toybox.WithCfgPath("config"),
		toybox.EnableComponents(&redis.RedisComponents{}),
	)

	fmt.Printf("%+v", tb)
}
