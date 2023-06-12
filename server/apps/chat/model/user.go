package model

import (
	"chat/toybox/common/dbutil"
	"chat/toybox/components/mysql"
	"encoding/json"
	"time"
)

type UserInfo struct {
	Id       string `gorm:"primaryKey" json:"id" redis:"id"`
	UserName string `gorm:"column:username" json:"userName" redis:"userName"`
	Passwd   string `gorm:"column:passwd" json:"-" redis:"-"`
	NickName string `gorm:"column:nickname" json:"nickName" redis:"nickName"`
	Avatar   string `gorm:"column:avatar" json:"avatar" redis:"avatar"`

	AccountType string    `gorm:"column:account_type" json:"-" redis:"-"`
	CreatedTime time.Time `gorm:"column:created_time" json:"-" redis:"-"`
}

func (u *UserInfo) BinaryMarshaler() ([]byte, error) {
	return json.Marshal(u)
}

func (UserInfo) TableName() string { return "t_user_info" }

func (u UserInfo) Query() (*UserInfo, error) {
	return dbutil.Select(mysql.Mysql.DB(), u)
}

func (u UserInfo) QueryList() ([]UserInfo, error) {
	// var users []UserInfo
	// err := mysql.Mysql.DB().Model(u).Where(u).Scan(&users).Error
	// if err != nil {
	// 	return nil, err
	// }
	// return users, nil
	return dbutil.SelectList(mysql.Mysql.DB(), u)
}

func (u UserInfo) GetUserInfoListByIds(ids []string) ([]UserInfo, error) {
	var users []UserInfo
	err := mysql.Mysql.DB().Model(u).Where("id IN ?").Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
