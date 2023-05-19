package chat

import (
	"chat/apps/chat/controller"
	"chat/apps/chat/middleware"
	"chat/toybox/res"
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
			RequiredComponents: []string{},
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
	chat.Use(
		middleware.GetLoginStatus(),
	)
	{ // heartbeat
		chat.GET("heartbeat", func(ctx *gin.Context) { ctx.String(200, "") })
	}
	{ // business
		chat.POST("login", controller.Login)

		chat.GET("user", func(c *gin.Context) {
			res.OK(c, "user api")
		})
	}
	r.Run(cs.Port)
	return nil
}
