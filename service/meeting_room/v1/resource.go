// Code generated by Lark OpenAPI.

package larkmeeting_room

import (
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

type V1 struct {
	MeetingRoom *meetingRoom // 事件
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		MeetingRoom: &meetingRoom{config: config},
	}
}

type meetingRoom struct {
	config *larkcore.Config
}
