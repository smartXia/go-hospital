package system

import (
	v1 "devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base").Use(middleware.OperationRecord())
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("loginByPhone", baseApi.LoginByPhone)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
