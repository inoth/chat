package chat

import (
	"chat/apps"
	"chat/toybox"

	"github.com/gin-gonic/gin"
)

var (
	serverName = "chat"
)

type ChatServer struct {
	ServerName          string
	Port                string
	DependentComponents []string

	// controllers []func(ctx *gin.Context)
}

func init() {
	apps.Add(serverName, func() toybox.Server {
		return &ChatServer{
			Port:                ":9978",
			DependentComponents: []string{"config", "redis", "mysql"},
		}
	})
}

func (cs *ChatServer) Init() error {
	return nil
}

func (cs *ChatServer) Start() error {
	r := gin.New()
	chat := r.Group(cs.ServerName)
	{ // heartbeat
		chat.GET("heartbeat", func(ctx *gin.Context) { ctx.String(200, "") })
	}
	{ // business

	}
	r.Run(cs.Port)
	return nil
}
