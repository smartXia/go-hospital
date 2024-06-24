package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportClockCommitSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	FlowId string `json:"flowId" form:"flowId" `
	request.PageInfo
}
