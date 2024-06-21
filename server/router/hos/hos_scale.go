package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosScaleRouter struct {
}

// InitHosScaleRouter 初始化 hosScale表 路由信息
func (s *HosScaleRouter) InitHosScaleRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosScaleRouter := Router.Group("hosScale").Use(middleware.OperationRecord())
	hosScaleRouterWithoutRecord := Router.Group("hosScale")
	hosScaleRouterWithoutAuth := PublicRouter.Group("hosScale")

	var hosScaleApi = v1.ApiGroupApp.HosApiGroup.HosScaleApi
	{
		hosScaleRouter.POST("createHosScale", hosScaleApi.CreateHosScale)             // 新建hosScale表
		hosScaleRouter.DELETE("deleteHosScale", hosScaleApi.DeleteHosScale)           // 删除hosScale表
		hosScaleRouter.DELETE("deleteHosScaleByIds", hosScaleApi.DeleteHosScaleByIds) // 批量删除hosScale表
		hosScaleRouter.PUT("updateHosScale", hosScaleApi.UpdateHosScale)              // 更新hosScale表
	}
	{
		hosScaleRouterWithoutRecord.GET("findHosScale", hosScaleApi.FindHosScale)                                // 根据ID获取hosScale表
		hosScaleRouterWithoutRecord.GET("getHosScaleList", hosScaleApi.GetHosScaleList)                          // 获取hosScale表列表
		hosScaleRouterWithoutRecord.GET("getGetCurrentUserHosScaleList", hosScaleApi.GetCurrentUserHosScaleList) // 获取当前用户的量表记录hosScale表列表
	}
	{
		hosScaleRouterWithoutAuth.GET("getHosScalePublic", hosScaleApi.GetHosScalePublic) // 获取hosScale表列表
	}
}
