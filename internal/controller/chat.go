package controller

import (
	"chat/internal/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/ginserver"
	"github.com/inoth/toybox/ginserver/res"
	"github.com/inoth/toybox/wsserver"
)

type ChatController struct {
	hub *wsserver.WebsocketServer
	svr *service.ChatService
}

func NewChatController(
	hub *wsserver.WebsocketServer,
	svr *service.ChatService,
) *ChatController {
	return &ChatController{
		hub: hub,
		svr: svr,
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
		{Method: "GET", Path: "/clients", Handle: []gin.HandlerFunc{cc.AllClients}},
	}
}

func (cc *ChatController) Connect(c *gin.Context) {
	//TODO: 客户端添加连接与断开函数调用事件
	// AfterConnection func()
	// AfterClose      func()
	clientID, err := wsserver.NewClientWithEvent(cc.hub, c.Writer, c.Request,
		func(c *wsserver.Client) {
			cc.svr.AddClientId(c.ID)
		},
		func(c *wsserver.Client) {
			cc.svr.RemoveClientId(c.ID)
		},
	)
	if err != nil {
		c.String(200, err.Error())
		return
	}
	fmt.Println(clientID)
}

func (cc *ChatController) AllClients(c *gin.Context) {
	res.Ok(c, "", cc.svr.GetClientId())
}
