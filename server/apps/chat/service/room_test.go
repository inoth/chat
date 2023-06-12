package service

import (
	"chat/apps/chat/model/request"
	"chat/toybox"
	"chat/toybox/components/config"
	"chat/toybox/components/logger/zap"
	"chat/toybox/components/mysql"
	"chat/toybox/components/redis"
	"fmt"
	"os"
	"testing"
)

func initComponents() {
	os.Setenv("GORUNEVN", "dev")
	toybox.New(
		toybox.WithComponentCfgPath("../../../config"),
		toybox.EnableComponents(
			zap.New(),
			config.New(),
			redis.New(),
			mysql.New(),
		),
	).InitComponents()
}

func TestCreateRoom(t *testing.T) {
	initComponents()

	type args struct {
		uid      string
		nickName string
		avatar   string
		req      request.RoomRequest
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{args: args{
			uid:      "v1",
			nickName: "inoth",
			avatar:   "001.png",
			req: request.RoomRequest{
				Name:   "inoth's room",
				Desc:   "",
				Passwd: "",
				Limit:  10,
				Tags:   "xxx,aaa",
			},
		}},
	}
	for _, tt := range tests {
		got, err := CreateRoom(tt.args.uid, tt.args.nickName, tt.args.avatar, tt.args.req)
		if err != nil {
			t.Errorf("CreateRoom() error = %v", err)
			return
		}
		fmt.Printf("CreateRoom() => %+v\n", got)
	}
}

func TestQueryRoomList(t *testing.T) {
	initComponents()

	type args struct {
		req request.RoomQueryRequest
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{args: args{req: request.RoomQueryRequest{
			Rid:  "",
			Name: "",
			Page: request.Page{
				1, 10,
			}},
		}},
	}
	for _, tt := range tests {
		got, total, err := QueryRoomList(tt.args.req)
		if err != nil {
			t.Errorf("QueryRoomList() error = %v", err)
			return
		}
		fmt.Printf("QueryRoomList() got = %+v, total %v\n", got, total)
	}
}
