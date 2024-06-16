package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportAdviceSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	FlowId *int   `json:"flowId" form:"flowId" `
	UserId *int   `json:"userId" form:"userId" `
	Period string `json:"period" form:"period" `
	Source string `json:"source" form:"source" `
	SyncWx *int   `json:"syncWx" form:"syncWx" `
	request.PageInfo
}
