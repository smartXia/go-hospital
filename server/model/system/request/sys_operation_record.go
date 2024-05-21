package request

import (
	"devops-manage/model/common/request"
	"devops-manage/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
