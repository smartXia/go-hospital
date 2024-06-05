package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosFlowSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name     string `json:"name" form:"name" `
	Uuid     string `json:"uuid" form:"uuid" `
	AskId    *int   `json:"askId" form:"askId" `
	AdviceId *int   `json:"adviceId" form:"adviceId" `
	request.PageInfo
}
