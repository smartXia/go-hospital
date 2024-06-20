package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosSportModeRouter struct {
}

// InitHosSportModeRouter 初始化 hosSportMode表 路由信息
func (s *HosSportModeRouter) InitHosSportModeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosSportModeRouter := Router.Group("hosSportMode").Use(middleware.OperationRecord())
	hosSportModeRouterWithoutRecord := Router.Group("hosSportMode")
	hosSportModeRouterWithoutAuth := PublicRouter.Group("hosSportMode")

	var hosSportModeApi = v1.ApiGroupApp.HosApiGroup.HosSportModeApi
	{
		hosSportModeRouter.POST("createHosSportMode", hosSportModeApi.CreateHosSportMode)             // 新建hosSportMode表
		hosSportModeRouter.DELETE("deleteHosSportMode", hosSportModeApi.DeleteHosSportMode)           // 删除hosSportMode表
		hosSportModeRouter.DELETE("deleteHosSportModeByIds", hosSportModeApi.DeleteHosSportModeByIds) // 批量删除hosSportMode表
		hosSportModeRouter.PUT("updateHosSportMode", hosSportModeApi.UpdateHosSportMode)              // 更新hosSportMode表
	}
	{
		hosSportModeRouterWithoutRecord.GET("findHosSportMode", hosSportModeApi.FindHosSportMode)           // 根据ID获取hosSportMode表
		hosSportModeRouterWithoutRecord.GET("getHosSportModeList", hosSportModeApi.GetHosSportModeList)     // 获取hosSportMode表列表
		hosSportModeRouterWithoutRecord.GET("getHosSportModeMatrix", hosSportModeApi.GetHosSportModeMatrix) // 获取hosSportMode表列表
	}
	{
		hosSportModeRouterWithoutAuth.GET("getHosSportModePublic", hosSportModeApi.GetHosSportModePublic) // 获取hosSportMode表列表
	}
}
