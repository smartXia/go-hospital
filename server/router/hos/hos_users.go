package hos

import (
	"devops-manage/api/v1"
	"github.com/gin-gonic/gin"
)

type HosUsersRouter struct {
}

// InitHosUsersRouter 初始化 hosUsers表 路由信息
func (s *HosUsersRouter) InitHosUsersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosUsersRouter := Router.Group("hosUsers")
	hosUsersRouterWithoutRecord := Router.Group("hosUsers")
	hosUsersRouterWithoutAuth := PublicRouter.Group("hosUsers")

	var hosUsersApi = v1.ApiGroupApp.HosApiGroup.HosUsersApi
	{
		hosUsersRouter.POST("createHosUsers", hosUsersApi.CreateHosUsers)             // 新建hosUsers表
		hosUsersRouter.DELETE("deleteHosUsers", hosUsersApi.DeleteHosUsers)           // 删除hosUsers表
		hosUsersRouter.DELETE("deleteHosUsersByIds", hosUsersApi.DeleteHosUsersByIds) // 批量删除hosUsers表
		hosUsersRouter.PUT("updateHosUsers", hosUsersApi.UpdateHosUsers)              // 更新hosUsers表
	}
	{
		hosUsersRouterWithoutRecord.GET("findHosUsers", hosUsersApi.FindHosUsers)                           // 根据ID获取hosUsers表
		hosUsersRouterWithoutRecord.GET("getHosUsersList", hosUsersApi.GetHosUsersList)                     // 获取hosUsers表列表
		hosUsersRouterWithoutRecord.GET("GetCurrentHosUsersList", hosUsersApi.GetCurrentHosUsersList)       // 获取hosUsers表列表
		hosUsersRouterWithoutRecord.GET("getCurrentUsersFuzhenList", hosUsersApi.GetCurrentUsersFuzhenList) // 获取hosUsers表列表
	}
	{
		hosUsersRouterWithoutAuth.GET("getHosUsersDataSource", hosUsersApi.GetHosUsersDataSource) // 获取hosUsers表数据源
		hosUsersRouterWithoutAuth.GET("getHosUsersPublic", hosUsersApi.GetHosUsersPublic)         // 获取hosUsers表列表
		hosUsersRouterWithoutAuth.GET("getHosUsersLastly", hosUsersApi.GetHosUsersLastly)         // 获取hosUsers表列表
	}
}
