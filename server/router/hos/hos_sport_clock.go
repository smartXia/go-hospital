package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosSportClockRouter struct {
}

// InitHosSportClockRouter 初始化 hosSportClock表 路由信息
func (s *HosSportClockRouter) InitHosSportClockRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosSportClockRouter := Router.Group("hosSportClock").Use(middleware.OperationRecord())
	hosSportClockRouterWithoutRecord := Router.Group("hosSportClock")
	hosSportClockRouterWithoutAuth := PublicRouter.Group("hosSportClock")

	var hosSportClockApi = v1.ApiGroupApp.HosApiGroup.HosSportClockApi
	{
		hosSportClockRouter.POST("createHosSportClock", hosSportClockApi.CreateHosSportClock)             // 新建hosSportClock表
		hosSportClockRouter.DELETE("deleteHosSportClock", hosSportClockApi.DeleteHosSportClock)           // 删除hosSportClock表
		hosSportClockRouter.DELETE("deleteHosSportClockByIds", hosSportClockApi.DeleteHosSportClockByIds) // 批量删除hosSportClock表
		hosSportClockRouter.PUT("updateHosSportClock", hosSportClockApi.UpdateHosSportClock)              // 更新hosSportClock表
	}
	{
		hosSportClockRouterWithoutRecord.GET("findHosSportClock", hosSportClockApi.FindHosSportClock)                             // 根据ID获取hosSportClock表
		hosSportClockRouterWithoutRecord.GET("getHosSportClockList", hosSportClockApi.GetHosSportClockList)                       // 获取hosSportClock表列表
		hosSportClockRouterWithoutRecord.GET("getCurrentUserHosSportClockList", hosSportClockApi.GetCurrentUserHosSportClockList) // 获取hosSportClock表列表
	}
	{
		hosSportClockRouterWithoutAuth.GET("getHosSportClockDataSource", hosSportClockApi.GetHosSportClockDataSource) // 获取hosSportClock表数据源
		hosSportClockRouterWithoutAuth.GET("getHosSportClockPublic", hosSportClockApi.GetHosSportClockPublic)         // 获取hosSportClock表列表
	}
}
