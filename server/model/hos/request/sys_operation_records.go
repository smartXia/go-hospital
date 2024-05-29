package request

import (
	"devops-manage/model/common/request"
	"time"
)

type SysOperationRecordsSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Method string `json:"method" form:"method" `
	Action string `json:"action" form:"action" `
	request.PageInfo
}
