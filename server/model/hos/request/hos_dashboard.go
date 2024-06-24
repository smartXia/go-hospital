package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosDashboardSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	OrgId *int `json:"orgId" form:"orgId" `
	request.PageInfo
}
