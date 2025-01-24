package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inoth/toybox/ginserver"
	"github.com/inoth/toybox/util/file"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (ic *IndexController) Prefix() string {
	return ""
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
	buf, err := file.ReadFile("static/index.html")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", buf)
}
