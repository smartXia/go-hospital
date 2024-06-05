package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportAdviceSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	FlowId *int   `json:"flowId" form:"flowId" `
	Name   string `json:"name" form:"name" `
	Source string `json:"source" form:"source" `
	request.PageInfo
}
