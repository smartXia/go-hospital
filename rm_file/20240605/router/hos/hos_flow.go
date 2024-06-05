package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type HosFlowRouter struct {
}

// InitHosFlowRouter 初始化 会诊流程 路由信息
func (s *HosFlowRouter) InitHosFlowRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	hosFlowRouter := Router.Group("hosFlow").Use(middleware.OperationRecord())
	hosFlowRouterWithoutRecord := Router.Group("hosFlow")
	hosFlowRouterWithoutAuth := PublicRouter.Group("hosFlow")

	var hosFlowApi = v1.ApiGroupApp.HosApiGroup.HosFlowApi
	{
		hosFlowRouter.POST("createHosFlow", hosFlowApi.CreateHosFlow)   // 新建会诊流程
		hosFlowRouter.DELETE("deleteHosFlow", hosFlowApi.DeleteHosFlow) // 删除会诊流程
		hosFlowRouter.DELETE("deleteHosFlowByIds", hosFlowApi.DeleteHosFlowByIds) // 批量删除会诊流程
		hosFlowRouter.PUT("updateHosFlow", hosFlowApi.UpdateHosFlow)    // 更新会诊流程
	}
	{
		hosFlowRouterWithoutRecord.GET("findHosFlow", hosFlowApi.FindHosFlow)        // 根据ID获取会诊流程
		hosFlowRouterWithoutRecord.GET("getHosFlowList", hosFlowApi.GetHosFlowList)  // 获取会诊流程列表
	}
	{
	    hosFlowRouterWithoutAuth.GET("getHosFlowDataSource", hosFlowApi.GetHosFlowDataSource)  // 获取会诊流程数据源
	    hosFlowRouterWithoutAuth.GET("getHosFlowPublic", hosFlowApi.GetHosFlowPublic)  // 获取会诊流程列表
	}
}
