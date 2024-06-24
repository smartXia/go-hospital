package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosUserPointSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	HosUserId *int `json:"hosUserId" form:"hosUserId" `
	request.PageInfo
}
