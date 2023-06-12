package service

import (
	"fmt"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	initComponents()

	type args struct {
		userId string
	}
	tests := []struct {
		args args
	}{
		// TODO: Add test cases.
		{args: args{
			userId: "test001",
		}},
	}
	for _, tt := range tests {
		got, err := GetUserInfo(tt.args.userId)
		if err != nil {
			t.Errorf("GetUserInfo() error = %v", err)
			return
		}
		fmt.Printf("GetUserInfo() => %+v", got)
	}
}
