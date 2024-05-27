package request

import (
	"devops-manage/model/common/request"
	"time"
)

type SysPostSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Name   string `json:"name" form:"name" `
	Code   string `json:"code" form:"code" `
	Enable *int   `json:"enable" form:"enable" `
	request.PageInfo
}
