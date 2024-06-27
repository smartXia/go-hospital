package hos

import (
	"devops-manage/api/v1"
	"github.com/gin-gonic/gin"
)

type SysPostRouter struct {
}

// InitSysPostRouter 初始化 sysPost表 路由信息
func (s *SysPostRouter) InitSysPostRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysPostRouter := Router.Group("sysPost")
	sysPostRouterWithoutRecord := Router.Group("sysPost")
	sysPostRouterWithoutAuth := PublicRouter.Group("sysPost")

	var sysPostApi = v1.ApiGroupApp.HosApiGroup.SysPostApi
	{
		sysPostRouter.POST("createSysPost", sysPostApi.CreateSysPost)             // 新建sysPost表
		sysPostRouter.DELETE("deleteSysPost", sysPostApi.DeleteSysPost)           // 删除sysPost表
		sysPostRouter.DELETE("deleteSysPostByIds", sysPostApi.DeleteSysPostByIds) // 批量删除sysPost表
		sysPostRouter.PUT("updateSysPost", sysPostApi.UpdateSysPost)              // 更新sysPost表
	}
	{
		sysPostRouterWithoutRecord.GET("findSysPost", sysPostApi.FindSysPost)       // 根据ID获取sysPost表
		sysPostRouterWithoutRecord.GET("getSysPostList", sysPostApi.GetSysPostList) // 获取sysPost表列表
	}
	{
		sysPostRouterWithoutAuth.GET("getSysPostPublic", sysPostApi.GetSysPostPublic) // 获取sysPost表列表
	}
}
