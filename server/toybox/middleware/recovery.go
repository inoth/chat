package middleware

import (
	"chat/toybox/components/logger"
	"chat/toybox/res"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Error(err)
				switch e := err.(type) {
				case error:
					res.Err(c, res.StatusInternalServerError, e)
				default:
					res.Err(c, res.StatusInternalServerError, fmt.Errorf("unknown error %v", err))
				}
			}
		}()
		c.Next()
	}
}
