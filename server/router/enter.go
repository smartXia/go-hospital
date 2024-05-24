package router

import (
	"devops-manage/router/example"
	"devops-manage/router/hos"
	"devops-manage/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Hos     hos.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
