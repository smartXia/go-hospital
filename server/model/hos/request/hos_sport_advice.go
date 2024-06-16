package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportAdviceSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	FlowId      *int   `json:"flowId" form:"flowId" `
	HosUserId   string `json:"hosUserId" form:"hosUserId" `
	SportModeId *int   `json:"sportModeId" form:"sportModeId" `
	Name        string `json:"name" form:"name" `
	Fuzhenriqi  string `json:"fuzhenriqi" form:"fuzhenriqi" `
	SyncWx      *int   `json:"syncWx" form:"syncWx" `
	request.PageInfo
}
