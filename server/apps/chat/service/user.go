package service

import (
	"chat/apps/chat/model"
	"chat/toybox/components/logger"
	"chat/toybox/components/redis"
	"context"
	"fmt"
)

const (
	userPrefix = "P5Chat:User:"
)

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

	user, err := getUserInfoByCache(userId)
	if err == nil {
		return user, err
	}

	user, err = query.Query()
	if err != nil {
		return nil, err
	}
	cacheMap := map[string]interface{}{
		"id":       user.Id,
		"userName": user.UserName,
		"nickName": user.NickName,
		"avatar":   user.Avatar,
	}
	err = redis.Rdc.Redis().HMSet(context.Background(), userPrefix+userId, cacheMap).Err()
	if err != nil {
		logger.Log.Warnf("user join cache err:%v", err)
	}
	return user, nil
}

func GetUserInfoListByIds(ids []string) {}

func getUserInfoByCache(userId string) (*model.UserInfo, error) {
	var cacheUser model.UserInfo
	err := redis.Rdc.Redis().HGetAll(context.Background(), userPrefix+userId).Scan(&cacheUser)
	if err != redis.Nil && len(cacheUser.Id) > 0 {
		return &cacheUser, nil
	}
	return nil, fmt.Errorf("user info no cache")
}
