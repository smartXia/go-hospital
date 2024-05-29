package request

import (
	"devops-manage/model/common/request"
	"time"
)

type SysUsersSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Username string `json:"username" form:"username" `
	request.PageInfo
}
