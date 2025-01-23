package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/ginserver"
	"github.com/inoth/toybox/wsserver"
)

type ChatController struct {
	hub *wsserver.WebsocketServer
}

func NewChatController(
	hub *wsserver.WebsocketServer,
) *ChatController {
	return &ChatController{
		hub: hub,
	}
}

func (cc *ChatController) Prefix() string {
	return "/api"
}

func (cc *ChatController) Middlewares() []gin.HandlerFunc {
	return nil
}

func (cc *ChatController) Routers() []ginserver.Router {
	return []ginserver.Router{
		{Method: "GET", Path: "/ws", Handle: []gin.HandlerFunc{cc.Connect}},
	}
}

func (cc *ChatController) Connect(c *gin.Context) {
	//TODO: 客户端添加连接与断开函数调用事件
	// AfterConnection func()
	// AfterClose      func()
	clientID, err := wsserver.NewClient(cc.hub, c.Writer, c.Request)
	if err != nil {
		c.String(200, err.Error())
		return
	}
	clientIds = append(clientIds, clientID)
	fmt.Println(clientID)
}
