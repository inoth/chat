package controller

import (
	"chat/apps/chat/middleware"
	"chat/apps/chat/service"
	"chat/toybox/res"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	uid := middleware.GetCurrentUserId(c)
	user, err := service.GetUserInfo(uid)
	if err != nil {
		res.Err(c, res.ParamErrorCode, err)
		return
	}
	res.OK(c, "ok", user)
}

func GetUserList(c *gin.Context) {}
