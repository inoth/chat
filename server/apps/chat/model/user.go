package model

import (
	"chat/toybox/components/mysql"
)

type UserInfo struct {
	Id       string `gorm:"column:id" json:"id"`
	UserName string `gorm:"column:username" json:"userName"`
	Passwd   string `gorm:"column:passwd" json:"-"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
}

func (UserInfo) TableName() string { return "t_user_info" }

func (u UserInfo) Query() (*UserInfo, error) {
	var user UserInfo
	err := mysql.Mysql.DB().Model(u).Where(u).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserInfo) QueryList() ([]UserInfo, error) {
	var users []UserInfo
	err := mysql.Mysql.DB().Model(u).Where(u).Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
