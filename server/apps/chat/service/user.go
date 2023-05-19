package service

import "chat/apps/chat/model"

func CheckUserInfo(username, passwd string) (*model.UserInfo, error) {
	query := model.UserInfo{UserName: username, Passwd: passwd}
	user, err := query.Query()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserInfo(userId string) (*model.UserInfo, error) {
	query := model.UserInfo{Id: userId}
	user, err := query.Query()
	if err != nil {
		return nil, err
	}
	return user, nil
}
