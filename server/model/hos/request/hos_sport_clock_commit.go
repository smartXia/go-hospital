package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportClockCommitSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	HosUserId      int        `json:"hosUserId" form:"hosUserId" `

	FlowId   string `json:"flowId" form:"flowId" `
	AdviceId string `json:"adviceId" form:"adviceId" `
	request.PageInfo
}
