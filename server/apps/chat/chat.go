package chat

import (
	"chat/apps/chat/controller"
	"chat/toybox/server"

	"github.com/gin-gonic/gin"
)

var (
	serverName = "chat"
	ver        = "v1"
)

type ChatServer struct {
	ServerName string
	Port       string
	// 依赖组件
	RequiredComponents []string
}

func init() {
	server.Add(serverName, func() server.Server {
		return &ChatServer{
			ServerName:         serverName,
			Port:               ":9978",
			RequiredComponents: []string{"config", "redis", "mysql"},
		}
	})
}

func (cs *ChatServer) Name() string {
	return serverName
}

func (cs *ChatServer) RequiredComponent() []string {
	return cs.RequiredComponents
}

func (cs *ChatServer) Start() error {
	r := gin.New()
	chat := r.Group(cs.ServerName + "/" + ver)
	{ // heartbeat
		chat.GET("heartbeat", func(ctx *gin.Context) { ctx.String(200, "") })
	}
	{ // business
		chat.POST("login", controller.Login)
	}
	r.Run(cs.Port)
	return nil
}
