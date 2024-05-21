package response

import (
	"devops-manage/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
