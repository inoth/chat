package service

import (
	"chat/apps/chat/model"
	"chat/apps/chat/model/request"
	"chat/toybox/common"
	"chat/toybox/common/dbutil"
	"chat/toybox/components/mysql"
	"chat/toybox/components/redis"
	"context"
	"encoding/json"
	"time"
)

const (
	roomPrefix = "P5Chat:Room:"
)

func checkRoomByUser(uid string) bool {
	c, _ := redis.Rdc.Redis().Exists(context.Background(), roomPrefix+uid).Result()
	return c > 0
}

func CreateRoom(uid, nickName, avatar string, req request.RoomRequest) (*model.Room, error) {
	if checkRoomByUser(uid) {
		room, err := redis.Rdc.Get(roomPrefix + uid)
		if err != nil {
			return nil, err
		}
		var res model.Room
		err = json.Unmarshal([]byte(room), &res)
		if err != nil {
			return nil, err
		}
		return &res, nil
	}
	addRoom := &model.Room{
		Owner: model.OwnerInfo{
			Id:       uid,
			NickName: nickName,
			Avatar:   avatar,
		},
		Uid:    uid,
		Name:   req.Name,
		Desc:   req.Desc,
		Passwd: req.Passwd,
		Limit:  req.Limit,
		Tags:   req.Tags,
	}
	addRoom.Id = common.RandString(16)
	addRoom.CreatedTime = time.Now()

	err := addRoom.Add()
	if err != nil {
		return nil, err
	}

	err = redis.Rdc.Set(roomPrefix+uid, addRoom.String())
	if err != nil {
		return nil, err
	}
	return addRoom, nil
}

func QueryRoomById(rid string) (*model.Room, error) {
	var room model.Room
	err := mysql.Mysql.DB().Model(model.Room{}).Where(model.Room{Id: rid}).First(&room).Error
	if err != nil {
		return nil, err
	}
	roomStr, err := redis.Rdc.Get(roomPrefix + room.Uid)
	if err != nil {
		return nil, err
	}

	var res model.Room
	err = json.Unmarshal([]byte(roomStr), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func QueryRoomList(req request.RoomQueryRequest) ([]model.Room, int64, error) {
	query := model.Room{
		Id:   req.Rid,
		Name: req.Name,
	}
	res, total, err := dbutil.SelectListAndTotal(mysql.Mysql.DB(), query, dbutil.Paginate(req.Index, req.Size))
	setOwnerInfo(&res)
	return res, total, err
}

func setOwnerInfo(rooms *[]model.Room) {
	for i := 0; i < len(*rooms); i++ {
		user, err := getUserInfoByCache((*rooms)[i].Uid)
		if err != nil {
			continue
		}
		(*rooms)[i].Owner = model.OwnerInfo{
			Id:       user.Id,
			NickName: user.NickName,
			Avatar:   user.Avatar,
		}
	}
}
