package request

import (
	"devops-manage/model/common/request"
	"devops-manage/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
