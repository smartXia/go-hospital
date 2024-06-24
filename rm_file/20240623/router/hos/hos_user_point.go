package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosUserPointRouter struct {
}

// InitHosUserPointRouter 初始化 hosUserPoint表 路由信息
func (s *HosUserPointRouter) InitHosUserPointRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosUserPointRouter := Router.Group("hosUserPoint").Use(middleware.OperationRecord())
	hosUserPointRouterWithoutRecord := Router.Group("hosUserPoint")
	hosUserPointRouterWithoutAuth := PublicRouter.Group("hosUserPoint")

	var hosUserPointApi = v1.ApiGroupApp.HosApiGroup.HosUserPointApi
	{
		hosUserPointRouter.POST("createHosUserPoint", hosUserPointApi.CreateHosUserPoint)             // 新建hosUserPoint表
		hosUserPointRouter.DELETE("deleteHosUserPoint", hosUserPointApi.DeleteHosUserPoint)           // 删除hosUserPoint表
		hosUserPointRouter.DELETE("deleteHosUserPointByIds", hosUserPointApi.DeleteHosUserPointByIds) // 批量删除hosUserPoint表
		hosUserPointRouter.PUT("updateHosUserPoint", hosUserPointApi.UpdateHosUserPoint)              // 更新hosUserPoint表
	}
	{
		hosUserPointRouterWithoutRecord.GET("findHosUserPoint", hosUserPointApi.FindHosUserPoint)       // 根据ID获取hosUserPoint表
		hosUserPointRouterWithoutRecord.GET("getHosUserPointList", hosUserPointApi.GetHosUserPointList) // 获取hosUserPoint表列表
	}
	{
		hosUserPointRouterWithoutAuth.GET("getHosUserPointPublic", hosUserPointApi.GetHosUserPointPublic) // 获取hosUserPoint表列表
	}
}
