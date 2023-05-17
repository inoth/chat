package redis

import (
	"fmt"
	"testing"
)

func TestConnectToRedis(t *testing.T) {
	client := &RedisComponents{
		URLs:     []string{""},
		Password: "",
	}
	client.Set("testkey", 123)
	res, _ := client.Get("testkey")
	fmt.Println(res)
}
