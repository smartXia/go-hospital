package hos

import (
	"devops-manage/api/v1"
	"github.com/gin-gonic/gin"
)

type SysOrgRouter struct {
}

// InitSysOrgRouter 初始化 sysOrg表 路由信息
func (s *SysOrgRouter) InitSysOrgRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysOrgRouter := Router.Group("sysOrg")
	sysOrgRouterWithoutRecord := Router.Group("sysOrg")
	sysOrgRouterWithoutAuth := PublicRouter.Group("sysOrg")

	var sysOrgApi = v1.ApiGroupApp.HosApiGroup.SysOrgApi
	{
		sysOrgRouter.POST("createSysOrg", sysOrgApi.CreateSysOrg)             // 新建sysOrg表
		sysOrgRouter.DELETE("deleteSysOrg", sysOrgApi.DeleteSysOrg)           // 删除sysOrg表
		sysOrgRouter.DELETE("deleteSysOrgByIds", sysOrgApi.DeleteSysOrgByIds) // 批量删除sysOrg表
		sysOrgRouter.PUT("updateSysOrg", sysOrgApi.UpdateSysOrg)              // 更新sysOrg表
	}
	{
		sysOrgRouterWithoutRecord.GET("findSysOrg", sysOrgApi.FindSysOrg)       // 根据ID获取sysOrg表
		sysOrgRouterWithoutRecord.GET("getSysOrgList", sysOrgApi.GetSysOrgList) // 获取sysOrg表列表
	}
	{
		sysOrgRouterWithoutAuth.GET("getSysOrgPublic", sysOrgApi.GetSysOrgPublic) // 获取sysOrg表列表
	}
}
