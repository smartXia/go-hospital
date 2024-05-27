package request

import (
	"devops-manage/model/common/request"
	"time"
)

type SysUsersSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Uuid     string `json:"uuid" form:"uuid" `
	Username string `json:"username" form:"username" `
	request.PageInfo
}
