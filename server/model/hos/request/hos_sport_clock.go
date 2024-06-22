package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportClockSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	HosUserId      int        `json:"hosUserId" form:"hosUserId" `
	FlowId         string     `json:"flowId" form:"flowId" `
	request.PageInfo
}
