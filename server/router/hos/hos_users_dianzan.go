package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosUsersDianzanRouter struct {
}

// InitHosUsersDianzanRouter 初始化 hosUsersDianzan表 路由信息
func (s *HosUsersDianzanRouter) InitHosUsersDianzanRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosUsersDianzanRouter := Router.Group("hosUsersDianzan").Use(middleware.OperationRecord())
	hosUsersDianzanRouterWithoutRecord := Router.Group("hosUsersDianzan")
	hosUsersDianzanRouterWithoutAuth := PublicRouter.Group("hosUsersDianzan")

	var hosUsersDianzanApi = v1.ApiGroupApp.HosApiGroup.HosUsersDianzanApi
	{
		hosUsersDianzanRouter.POST("createHosUsersDianzan", hosUsersDianzanApi.CreateHosUsersDianzan)             // 新建hosUsersDianzan表
		hosUsersDianzanRouter.DELETE("deleteHosUsersDianzan", hosUsersDianzanApi.DeleteHosUsersDianzan)           // 删除hosUsersDianzan表
		hosUsersDianzanRouter.DELETE("deleteHosUsersDianzanByIds", hosUsersDianzanApi.DeleteHosUsersDianzanByIds) // 批量删除hosUsersDianzan表
		hosUsersDianzanRouter.PUT("updateHosUsersDianzan", hosUsersDianzanApi.UpdateHosUsersDianzan)              // 更新hosUsersDianzan表
	}
	{
		hosUsersDianzanRouterWithoutRecord.GET("findHosUsersDianzan", hosUsersDianzanApi.FindHosUsersDianzan)       // 根据ID获取hosUsersDianzan表
		hosUsersDianzanRouterWithoutRecord.GET("getHosUsersDianzanList", hosUsersDianzanApi.GetHosUsersDianzanList) // 获取hosUsersDianzan表列表
	}
	{
		hosUsersDianzanRouterWithoutAuth.GET("getHosUsersDianzanPublic", hosUsersDianzanApi.GetHosUsersDianzanPublic) // 获取hosUsersDianzan表列表
	}
}
