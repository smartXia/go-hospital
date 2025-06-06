package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportClockSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name           string `json:"name" form:"name" `
	HosUserId      *int   `json:"hosUserId" form:"hosUserId" `
	FlowId         *int   `json:"flowId" form:"flowId" `
	AdviceId       *int   `json:"adviceId" form:"adviceId" `
	ClockStartTime string `json:"clockStartTime" form:"clockStartTime" `
	ClockEndTime   string `json:"clockEndTime" form:"clockEndTime" `
	Phone          string `json:"phone" form:"phone" `
	YuyueType      string `json:"yuyueType" form:"yuyueType"`
	request.PageInfo
}
