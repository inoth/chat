package controller

import (
	"chat/internal/service"
	"fmt"

	"github.com/inoth/toybox/wsserver"
)

type body struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

type MessageController struct {
	svr *service.ChatService
}

func NewMessageController(svr *service.ChatService) *MessageController {
	return &MessageController{
		svr: svr,
	}
}

func (m *MessageController) Handler() wsserver.HandlerFunc {
	return func(c *wsserver.Context) {
		// fmt.Printf("%v\n", string(c.Body()))
		var data body
		err := c.BindJson(&data)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
		if data.ID == "all" {
			for _, id := range m.svr.GetClientId() {
				c.String(id, data.Body)
			}
		} else {
			c.String(data.ID, data.Body)
		}
	}
}
