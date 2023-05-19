package middleware

import (
	"chat/toybox/common"
	"chat/toybox/components/config"
	"chat/toybox/res"
	"fmt"
	"strings"

	mid "chat/toybox/middleware"

	"github.com/gin-gonic/gin"
)

func GetLoginStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "login") {
			c.Next()
			return
		}
		token := ""
		token, _ = c.Cookie("Authorization")
		if len(token) <= 0 {
			token = c.GetHeader("Authorization")
		}
		if len(token) <= 0 {
			res.Err(c, res.Unauthorized, fmt.Errorf("unable to get token"))
			c.Abort()
			return
		}
		customerInfo, err := mid.ParseToken(config.Cfg.GetString(""), token)
		if err != nil {
			res.Err(c, res.Unauthorized, fmt.Errorf("token parsing failure: %v", err))
			c.Abort()
			return
		}
		c.Set("user_info", customerInfo.UserInfo)
		c.Next()
	}
}

func GetCurrentUserId(c *gin.Context) string {
	user, ok := c.Get("user_info")
	if !ok {
		return ""
	}
	u, ok := user.(map[string]interface{})
	if !ok {
		return ""
	}
	id, ok := common.GetStringValue(u, "id")
	if !ok {
		return ""
	}
	return id
}

func GetCurrentUserName(c *gin.Context) string {
	user, ok := c.Get("user_info")
	if !ok {
		return ""
	}
	u, ok := user.(map[string]interface{})
	if !ok {
		return ""
	}
	id, ok := common.GetStringValue(u, "name")
	if !ok {
		return ""
	}
	return id
}
