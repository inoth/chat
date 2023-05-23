package chat

import (
	"chat/apps/chat/controller"
	"chat/apps/chat/middleware"
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
			RequiredComponents: []string{"config", "logger"},
		}
	})
}

func (cs *ChatServer) Name() string {
	return serverName
}

func (cs *ChatServer) RequiredComponent() []string {
	return cs.RequiredComponents
}

func (cs *ChatServer) Close() error { return nil }

func (cs *ChatServer) Start() error {
	r := gin.New()
	{ // heartbeat
		r.GET("heartbeat", func(ctx *gin.Context) { ctx.String(200, "") })
	}

	chat := r.Group(cs.ServerName + "/" + ver)
	{ // business
		chat.Use(
			middleware.GetLoginStatus(),
		)
		chat.POST("login", controller.Login)
	}
	user := chat.Group("user")
	{
		user.GET("", controller.GetUser)
	}
	room := chat.Group("room")
	{
		room.POST("", controller.CreateRoom)
		room.PATCH("", controller.UpdateRoom)
		room.GET("/:roomId", controller.GetRoom)
		room.GET("", controller.GetRoomList)

		room.POST("/join", controller.JoinRoom)
		room.POST("/exit", controller.ExitRoom)
	}
	ws := r.Group("ws")
	{
		// 直接加入ws连接，同时分配默认room
		ws.GET("", controller.ConnectUpgrade)
	}
	r.Run(cs.Port)
	return nil
}
