package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosLocalAskRouter struct {
}

// InitHosLocalAskRouter 初始化 hosLocalAsk表 路由信息
func (s *HosLocalAskRouter) InitHosLocalAskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosLocalAskRouter := Router.Group("hosLoaclAsk").Use(middleware.OperationRecord())
	hosLocalAskRouterWithoutRecord := Router.Group("hosLoaclAsk")
	hosLocalAskRouterWithoutAuth := PublicRouter.Group("hosLoaclAsk")

	var hosLocalAskApi = v1.ApiGroupApp.HosApiGroup.HosLocalAskApi
	{
		hosLocalAskRouter.POST("createHosLoaclAsk", hosLocalAskApi.CreateHosLocalAsk)             // 新建hosLocalAsk表
		hosLocalAskRouter.DELETE("deleteHosLoaclAsk", hosLocalAskApi.DeleteHosLocalAsk)           // 删除hosLocalAsk表
		hosLocalAskRouter.DELETE("deleteHosLoaclAskByIds", hosLocalAskApi.DeleteHosLocalAskByIds) // 批量删除hosLocalAsk表
		hosLocalAskRouter.PUT("updateHosLoaclAsk", hosLocalAskApi.UpdateHosLocalAsk)              // 更新hosLocalAsk表
	}
	{
		hosLocalAskRouterWithoutRecord.GET("findHosLoaclAsk", hosLocalAskApi.FindHosLocalAsk)       // 根据ID获取hosLocalAsk表
		hosLocalAskRouterWithoutRecord.GET("getHosLoaclAskList", hosLocalAskApi.GetHosLocalAskList) // 获取hosLocalAsk表列表
	}
	{
		hosLocalAskRouterWithoutAuth.GET("getHosLoaclAskPublic", hosLocalAskApi.GetHosLocalAskPublic) // 获取hosLocalAsk表列表
	}
}
