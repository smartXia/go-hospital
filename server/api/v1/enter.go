package v1

import (
	"devops-manage/api/v1/example"
	"devops-manage/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
