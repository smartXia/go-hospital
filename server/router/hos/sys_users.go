package hos

import (
	"devops-manage/api/v1"
	"devops-manage/middleware"
	"github.com/gin-gonic/gin"
)

type SysUsersRouter struct {
}

// InitSysUsersRouter 初始化 sysUsers表 路由信息
func (s *SysUsersRouter) InitSysUsersRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysUsersRouter := Router.Group("sysUsers").Use(middleware.OperationRecord())
	sysUsersRouterWithoutRecord := Router.Group("sysUsers")
	sysUsersRouterWithoutAuth := PublicRouter.Group("sysUsers")

	var sysUsersApi = v1.ApiGroupApp.HosApiGroup.SysUsersApi
	{
		sysUsersRouter.POST("createSysUsers", sysUsersApi.CreateSysUsers)             // 新建sysUsers表
		sysUsersRouter.DELETE("deleteSysUsers", sysUsersApi.DeleteSysUsers)           // 删除sysUsers表
		sysUsersRouter.DELETE("deleteSysUsersByIds", sysUsersApi.DeleteSysUsersByIds) // 批量删除sysUsers表
		sysUsersRouter.PUT("updateSysUsers", sysUsersApi.UpdateSysUsers)              // 更新sysUsers表
	}
	{
		sysUsersRouterWithoutRecord.GET("findSysUsers", sysUsersApi.FindSysUsers)       // 根据ID获取sysUsers表
		sysUsersRouterWithoutRecord.GET("getSysUsersList", sysUsersApi.GetSysUsersList) // 获取sysUsers表列表
	}
	{
		sysUsersRouterWithoutAuth.GET("getSysUsersPublic", sysUsersApi.GetSysUsersPublic) // 获取sysUsers表列表
	}
}
