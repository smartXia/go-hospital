package hos

import (
	"devops-manage/api/v1"
	"github.com/gin-gonic/gin"
)

type HosDashboardRouter struct {
}

// InitHosDashboardRouter 初始化 hosDashboard表 路由信息
func (s *HosDashboardRouter) InitHosDashboardRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosDashboardRouter := Router.Group("hosDashboard")
	hosDashboardRouterWithoutRecord := Router.Group("hosDashboard")
	hosDashboardRouterWithoutAuth := PublicRouter.Group("hosDashboard")

	var hosDashboardApi = v1.ApiGroupApp.HosApiGroup.HosDashboardApi
	{
		hosDashboardRouter.POST("createHosDashboard", hosDashboardApi.CreateHosDashboard)             // 新建hosDashboard表
		hosDashboardRouter.DELETE("deleteHosDashboard", hosDashboardApi.DeleteHosDashboard)           // 删除hosDashboard表
		hosDashboardRouter.DELETE("deleteHosDashboardByIds", hosDashboardApi.DeleteHosDashboardByIds) // 批量删除hosDashboard表
		hosDashboardRouter.PUT("updateHosDashboard", hosDashboardApi.UpdateHosDashboard)              // 更新hosDashboard表
	}
	{
		hosDashboardRouterWithoutRecord.GET("findHosDashboard", hosDashboardApi.FindHosDashboard)               // 根据ID获取hosDashboard表
		hosDashboardRouterWithoutRecord.GET("getHosDashboardList", hosDashboardApi.GetHosDashboardList)         // 获取hosDashboard表列表
		hosDashboardRouterWithoutRecord.GET("getCurrentDashBoardInfo", hosDashboardApi.GetCurrentDashBoardInfo) // 获取hosDashboard表列表
	}
	{
		hosDashboardRouterWithoutAuth.GET("getHosDashboardDataSource", hosDashboardApi.GetHosDashboardDataSource) // 获取hosDashboard表数据源
		hosDashboardRouterWithoutAuth.GET("getHosDashboardPublic", hosDashboardApi.GetHosDashboardPublic)         // 获取hosDashboard表列表
	}
}
