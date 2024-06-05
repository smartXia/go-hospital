package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosLoaclAskRouter struct {
}

// InitHosLoaclAskRouter 初始化 hosLoaclAsk表 路由信息
func (s *HosLoaclAskRouter) InitHosLoaclAskRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosLoaclAskRouter := Router.Group("hosLoaclAsk").Use(middleware.OperationRecord())
	hosLoaclAskRouterWithoutRecord := Router.Group("hosLoaclAsk")
	hosLoaclAskRouterWithoutAuth := PublicRouter.Group("hosLoaclAsk")

	var hosLoaclAskApi = v1.ApiGroupApp.HosApiGroup.HosLoaclAskApi
	{
		hosLoaclAskRouter.POST("createHosLoaclAsk", hosLoaclAskApi.CreateHosLoaclAsk)             // 新建hosLoaclAsk表
		hosLoaclAskRouter.DELETE("deleteHosLoaclAsk", hosLoaclAskApi.DeleteHosLoaclAsk)           // 删除hosLoaclAsk表
		hosLoaclAskRouter.DELETE("deleteHosLoaclAskByIds", hosLoaclAskApi.DeleteHosLoaclAskByIds) // 批量删除hosLoaclAsk表
		hosLoaclAskRouter.PUT("updateHosLoaclAsk", hosLoaclAskApi.UpdateHosLoaclAsk)              // 更新hosLoaclAsk表
	}
	{
		hosLoaclAskRouterWithoutRecord.GET("findHosLoaclAsk", hosLoaclAskApi.FindHosLoaclAsk)       // 根据ID获取hosLoaclAsk表
		hosLoaclAskRouterWithoutRecord.GET("getHosLoaclAskList", hosLoaclAskApi.GetHosLoaclAskList) // 获取hosLoaclAsk表列表
	}
	{
		hosLoaclAskRouterWithoutAuth.GET("getHosLoaclAskPublic", hosLoaclAskApi.GetHosLoaclAskPublic) // 获取hosLoaclAsk表列表
	}
}
