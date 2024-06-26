package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportClockCommitSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	FlowId    int `json:"flowId" form:"flowId" `
	HosUserId int `json:"hosUserId" form:"hosUserId" `
	ClockId   int `json:"clockId" form:"clockId" `
	request.PageInfo
}
