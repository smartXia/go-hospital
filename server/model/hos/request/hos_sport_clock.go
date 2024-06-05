package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportClockSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Uid    string `json:"uid" form:"uid" `
	FlowId string `json:"flowId" form:"flowId" `
	request.PageInfo
}
