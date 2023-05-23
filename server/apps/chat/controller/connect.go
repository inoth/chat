package controller

import (
	"chat/apps/imchat/core"
	"chat/toybox/common"

	"github.com/gin-gonic/gin"
)

func ConnectUpgrade(c *gin.Context) {
	uid := common.RandString(12)
	client, err := core.NewClient(uid, core.GetChatHub(), c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	client.Run()
}
