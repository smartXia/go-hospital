package request

import (
	"devops-manage/model/common/request"
	"time"
)

type HosSportModeSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Category  string `json:"category" form:"category" `
	Buwei     string `json:"buwei" form:"buwei" `
	Tiduan    string `json:"tiduan" form:"tiduan" `
	Fangxiang string `json:"fangxiang" form:"fangxiang" `
	Weizhi    string `json:"weizhi" form:"weizhi" `
	Source    string `json:"source" form:"source" `
	request.PageInfo
}
