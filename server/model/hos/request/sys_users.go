package request

import (
	"devops-manage/model/common/request"
	"time"
)

type SysUsersSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Phone    string `json:"phone" form:"phone" `
	Enable   *int   `json:"enable" form:"enable" `
	Hospital int    `json:"hospital" form:"hospital" `
	NickName string `json:"nickName" form:"nickName" `
	Dept     string `json:"dept" form:"dept" `
	Post     string `json:"post" form:"post" `
	TenantId int    `json:"tenantId" form:"tenantId" `
	request.PageInfo
}
