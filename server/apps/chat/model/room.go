package model

import (
	"chat/toybox/components/mysql"
	"encoding/json"
	"time"
)

type Room struct {
	Id          string    `gorm:"primaryKey;column:id" json:"id"`
	Uid         string    `gorm:"column:uid" json:"-"`
	Owner       OwnerInfo `gorm:"-" json:"owner"`
	Name        string    `gorm:"column:name" json:"name"`
	Desc        string    `gorm:"column:desc" json:"desc"`
	Passwd      string    `gorm:"column:passwd" json:"passwd"`
	Limit       int       `gorm:"column:limit" json:"limit"`
	Tags        string    `gorm:"column:tags" json:"tags"`
	CreatedTime time.Time `gorm:"column:created_time" json:"createdTime"`
}

type OwnerInfo struct {
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
