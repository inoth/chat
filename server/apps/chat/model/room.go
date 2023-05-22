package model

import (
	"chat/toybox/components/mysql"
	"encoding/json"
	"time"
)

type Room struct {
	Id          string    `gorm:"primaryKey" json:"id"`
	Uid         string    `gorm:"column:uid" json:"-"`
	Owner       RoomOwner `gorm:"-" json:"owner"`
	RoomName    string    `gorm:"column:roomName" json:"roomName"`
	Desc        string    `gorm:"column:description" json:"desc"`
	Passwd      string    `gorm:"column:password" json:"passwd"`
	UserLimit   int       `gorm:"column:userLimit" json:"userLimit"`
	CreatedTime time.Time `gorm:"column:createdTime" json:"createdTime"`
}

type RoomOwner struct {
	Id       string `json:"id"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

func (ri Room) String() string {
	buf, _ := json.Marshal(ri)
	return string(buf)
}

func (Room) TableName() string { return "t_room" }

func (r *Room) Add() error {
	return mysql.Mysql.DB().Create(r).Error
}
