package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosUsersSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Uuid        string `json:"uuid" form:"uuid" `
	Username    string `json:"username" form:"username" `
	NickName    string `json:"nickName" form:"nickName" `
	Phone       string `json:"phone" form:"phone" `
	JianhuPhone string `json:"jianhuPhone" form:"jianhuPhone" `
	Jianhuren   string `json:"jianhuren" form:"jianhuren" `
	request.PageInfo
}
