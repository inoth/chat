package controller

import (
	"chat/static"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/ginserver"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (ic *IndexController) Prefix() string {
	return "/api"
}

func (ic *IndexController) Middlewares() []gin.HandlerFunc {
	return nil
}

func (ic *IndexController) Routers() []ginserver.Router {
	return []ginserver.Router{
		{Method: "GET", Path: "/chat", Handle: []gin.HandlerFunc{ic.ChatHome}},
	}
}

func (ic *IndexController) ChatHome(c *gin.Context) {
	buf, err := static.StaticIndex.ReadFile("index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", buf)
}
