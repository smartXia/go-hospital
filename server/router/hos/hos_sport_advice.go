package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosSportAdviceRouter struct {
}

// InitHosSportAdviceRouter 初始化 hosSportAdvice表 路由信息
func (s *HosSportAdviceRouter) InitHosSportAdviceRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosSportAdviceRouter := Router.Group("hosSportAdvice").Use(middleware.OperationRecord())
	hosSportAdviceRouterWithoutRecord := Router.Group("hosSportAdvice")
	hosSportAdviceRouterWithoutAuth := PublicRouter.Group("hosSportAdvice")

	var hosSportAdviceApi = v1.ApiGroupApp.HosApiGroup.HosSportAdviceApi
	{
		hosSportAdviceRouter.POST("createHosSportAdvice", hosSportAdviceApi.CreateHosSportAdvice)             // 新建hosSportAdvice表
		hosSportAdviceRouter.DELETE("deleteHosSportAdvice", hosSportAdviceApi.DeleteHosSportAdvice)           // 删除hosSportAdvice表
		hosSportAdviceRouter.DELETE("deleteHosSportAdviceByIds", hosSportAdviceApi.DeleteHosSportAdviceByIds) // 批量删除hosSportAdvice表
		hosSportAdviceRouter.PUT("updateHosSportAdvice", hosSportAdviceApi.UpdateHosSportAdvice)              // 更新hosSportAdvice表
	}
	{
		hosSportAdviceRouterWithoutRecord.GET("findHosSportAdvice", hosSportAdviceApi.FindHosSportAdvice)                     // 根据ID获取hosSportAdvice表
		hosSportAdviceRouterWithoutRecord.GET("getHosSportAdviceList", hosSportAdviceApi.GetHosSportAdviceList)               // 获取hosSportAdvice表列表
		hosSportAdviceRouterWithoutRecord.GET("getCurrentHosSportAdviceList", hosSportAdviceApi.GetCurrentHosSportAdviceList) // 获取hosSportAdvice表列表
	}
	{
		hosSportAdviceRouterWithoutAuth.GET("getHosSportAdvicePublic", hosSportAdviceApi.GetHosSportAdvicePublic) // 获取hosSportAdvice表列表
	}
}
