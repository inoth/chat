package service

import (
	"chat/apps/chat/model/request"
	"reflect"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	type args struct {
		uid      string
		nickName string
		avatar   string
		req      request.RoomRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateRoom(tt.args.uid, tt.args.nickName, tt.args.avatar, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRoom() got = %v, want %v", got, tt.want)
			}
		})
	}
}
