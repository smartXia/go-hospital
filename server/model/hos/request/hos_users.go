package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosUsersSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Username    string `json:"username" form:"username" `
	NickName    string `json:"nickName" form:"nickName" `
	Phone       string `json:"phone" form:"phone" `
	JianhuPhone string `json:"jianhuPhone" form:"jianhuPhone" `
	CardNo      string `json:"cardNo" form:"cardNo" `
	request.PageInfo
}
