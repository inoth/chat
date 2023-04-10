package middleware

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/res"
	"github.com/inoth/toybox/utils"

	jwtauth "github.com/inoth/toybox/utils/jwt_auth"
)

func JwtAuthToken(jwtSignKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := regexp.MustCompile("login")
		if r.MatchString(c.Request.URL.Path) {
			c.Next()
			return
		}

		token := strings.ReplaceAll(c.GetHeader("Authorization"), "Bearer ", "")
		if len(token) <= 0 {
			var err error
			token, err = c.Cookie("Authorization")
			if err != nil {
				res.ResultErr(c, res.InvalidRequestErrorCode, errors.New("session not found"))
				c.Abort()
				return
			}
		}
		if len(token) <= 0 {
			res.ResultErr(c, res.InvalidRequestErrorCode, errors.New("session not found"))
			c.Abort()
			return
		}
		if len(jwtSignKey) <= 0 {
			jwtSignKey = jwtauth.DEFAULT_SIGNKEY
		}
		customerInfo, err := jwtauth.ParseToken(jwtSignKey, token)
		if err != nil {
			res.ResultErr(c, res.InvalidRequestErrorCode, err)
			c.Abort()
			return
		}
		c.Set("user", customerInfo.UserInfo)
	}
}

func GetUserInfoBy(c *gin.Context, key string) string {
	userInfo, ok := c.Get("user")
	if !ok {
		return ""
	}
	user := userInfo.(map[string]interface{})
	val, ok := utils.GetStringValue(user, key)
	if !ok {
		return ""
	}
	return val
}
