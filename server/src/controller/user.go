package controller

import (
	"chat/src/request"
	"errors"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/components/config"
	"github.com/inoth/toybox/middleware"
	"github.com/inoth/toybox/res"
	"github.com/inoth/toybox/utils/encrypt"
	jwtauth "github.com/inoth/toybox/utils/jwt_auth"
)

func getAvatar() string {
	avatars := config.Cfg.GetStringSlice("base.avatars")
	return avatars[rand.Int()%len(avatars)]
}

func Login(c *gin.Context) {
	req, ok := middleware.RequestJsonParamHandler[request.LoginRequest](c)
	if !ok {
		return
	}
	var users []request.LoginRequest
	err := config.Cfg.UnmarshalKey("base.users", &users)
	if err != nil {
		res.ResultErr(c, res.ParamErrorCode, err)
		return
	}
	for _, user := range users {
		if user.UserName == req.UserName && user.Passwd == req.Passwd {
			id := encrypt.Md5(c.ClientIP() + user.UserName)
			userInfo := map[string]interface{}{"id": id, "name": user.UserName, "avatar": getAvatar()}
			token, err := jwtauth.CreateToken(jwtauth.DEFAULT_SIGNKEY, userInfo)
			if err != nil {
				res.ResultErr(c, res.ParamErrorCode, err)
				return
			}
			c.Header("Authorization", token)
			c.SetCookie("Authorization", token, 60*60*24, "/", "localhost", false, true)
			res.ResultOk(c, res.SuccessCode, userInfo)
			return
		}
	}
	res.ResultErr(c, res.ParamErrorCode, errors.New("账号密码错误"))
}
