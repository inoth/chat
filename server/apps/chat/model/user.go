package model

import (
	"chat/toybox/components/mysql"
	"time"
)

type UserInfo struct {
	Id       string `gorm:"primaryKey" json:"id"`
	UserName string `gorm:"column:username" json:"userName"`
	Passwd   string `gorm:"column:passwd" json:"-"`
	NickName string `gorm:"column:nickname" json:"nickName"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`

	AccountType string    `gorm:"column:account_type" json:"-"`
	CreatedTime time.Time `gorm:"column:created_time" json:"-"`
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

func (u UserInfo) GetUserInfoListByIds(ids []string) ([]UserInfo, error) {
	var users []UserInfo
	err := mysql.Mysql.DB().Model(u).Where("id IN ?").Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
