package controller

import (
	"chat/apps/chat/model/request"
	"chat/apps/chat/service"
	"chat/toybox/components/config"
	mid "chat/toybox/middleware"
	"chat/toybox/res"

	"github.com/gin-gonic/gin"
)

// 常规登录
func Login(c *gin.Context) {
	req, ok := mid.ParseJsonParam[request.LoginRequest](c)
	if !ok {
		return
	}
	user, err := service.CheckUserInfo(req.UserName, req.Passwd)
	if err != nil {
		res.Err(c, res.ParamErrorCode, err)
		return
	}
	token, err := mid.CreateToken(config.Cfg.GetString("jwt_key"), map[string]interface{}{
		"id":     user.Id,
		"name":   user.UserName,
		"avatar": user.Avatar,
	})
	if err != nil {
		res.Err(c, res.InternalErrorCode, err)
		return
	}
	c.Header("Authorization", token)
	c.SetCookie("Authorization", token, 86400, "", "", false, false)
	res.OK(c, "ok", user)
}

// 外部授权登录
func AuthLogin(c *gin.Context) {}
