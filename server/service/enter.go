package service

import (
	"devops-manage/service/example"
	"devops-manage/service/hos"
	"devops-manage/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	HosServiceGroup     hos.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
