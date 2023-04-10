package router

import (
	col "chat/src/controller"

	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/res"
)

type BaseRouter struct{}

func (BaseRouter) Prefix() string {
	return "chat/v1"
}
func (BaseRouter) LoadRouter(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		res.ResultOk(c, 200, "ok")
		// c.String(200, "ok")
	})

	r.POST("/login", col.Login)
}

type WsRouter struct{}

func (WsRouter) Prefix() string {
	return "ws/v1"
}
func (WsRouter) LoadRouter(r *gin.RouterGroup) {

}
