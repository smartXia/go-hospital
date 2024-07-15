package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosFlowSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Enable         *int       `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"` //状态
	Name           string     `json:"name" form:"name" `
	Uuid           string     `json:"uuid" form:"uuid" `
	Type           string     `json:"type" form:"type" `
	AskId          *int       `json:"askId" form:"askId" `
	HosUserId      *int       `json:"hosUserId" form:"hosUserId" `
	AdviceId       *int       `json:"adviceId" form:"adviceId" `
	request.PageInfo
}
