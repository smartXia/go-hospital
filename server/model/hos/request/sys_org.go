package request

import (
	"devops-manage/model/common/request"
	"time"
)

type SysOrgSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name string `json:"name" form:"name" `
	Desc string `json:"desc" form:"desc" `
	request.PageInfo
}
