package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	gredis "github.com/go-redis/redis/v8"
)

var (
	// componentName = "redis"
	redisOnce sync.Once
	Redis     *RedisComponents
)

// func init() {
// 	components.Add(componentName, func() toybox.Component {
// 		return &RedisComponents{}
// 	})
// }

type RedisComponents struct {
	client *gredis.ClusterClient

	URLs        []string `toml:"urls" json:"urls" yaml:"urls"`
	Username    string   `toml:"username" json:"username" yaml:"username"`
	Password    string   `toml:"password" json:"password" yaml:"password"`
	PoolSize    int      `toml:"poolsize" json:"poolsize" yaml:"poolsize"`
	PoolTimeout int      `toml:"pooltimeout" json:"pooltimeout" yaml:"pooltimeout"`
}

func (rc *RedisComponents) Close() error { return rc.client.Close() }

func (rc *RedisComponents) Init() (err error) {
	redisOnce.Do(func() {
		client := gredis.NewClusterClient(&gredis.ClusterOptions{
			Addrs:       rc.URLs,
			Password:    rc.Password,
			Username:    rc.Username,
			PoolSize:    rc.PoolSize,
			PoolTimeout: time.Duration(rc.PoolTimeout),
		})
		_, err = client.Ping(context.Background()).Result()
		if err != nil {
			err = fmt.Errorf("failed to connect to redis: %v", err)
			return
		}
		rc.client = client
		Redis = rc
	})
	return
}

func (rc *RedisComponents) Get(key string) (string, error) {
	res, err := rc.client.Get(context.Background(), key).Result()
	if err == gredis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return res, nil
}

func (rc *RedisComponents) Set(key string, val interface{}, expiration ...time.Duration) error {
	expir := time.Duration(0)
	if len(expiration) > 0 {
		expir = expiration[0]
	}
	return rc.client.Set(context.Background(), key, val, expir).Err()
}

func (rc *RedisComponents) Redis() *gredis.ClusterClient {
	return rc.client
}
