package controller

import (
	"github.com/gin-gonic/gin"

	"chat/apps/chat/middleware"
	"chat/apps/chat/model/request"
	"chat/apps/chat/service"
	mid "chat/toybox/middleware"
	"chat/toybox/res"
)

func GetRoom(c *gin.Context) {
	req, ok := mid.ParseQueryParam[request.RoomQueryRequest](c)
	if !ok {
		return
	}
	room, err := service.QueryRoomById(req.RoomId)
	if err != nil {
		res.Err(c, res.ParamErrorCode, err)
		return
	}
	res.OK(c, "ok", room)
}

func GetRoomList(c *gin.Context) {}

func CreateRoom(c *gin.Context) {
	req, ok := mid.ParseJsonParam[request.RoomRequest](c)
	if !ok {
		return
	}
	uid := middleware.GetCurrentUserId(c)
	name := middleware.GetCurrentUserName(c)
	avatar := middleware.GetCurrentAvatar(c)
	room, err := service.CreateRoom(uid, name, avatar, req)
	if err != nil {
		res.Err(c, res.InternalErrorCode, err)
		return
	}
	res.OK(c, "ok", room)
}

func DelRoom(c *gin.Context)    {}
func UpdateRoom(c *gin.Context) {}

func JoinRoom(c *gin.Context) {}
func ExitRoom(c *gin.Context) {}
