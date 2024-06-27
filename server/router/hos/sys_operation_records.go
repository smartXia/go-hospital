package hos

import (
	"devops-manage/api/v1"
	"github.com/gin-gonic/gin"
)

type SysOperationRecordsRouter struct {
}

// InitSysOperationRecordsRouter 初始化 sysOperationRecords表 路由信息
func (s *SysOperationRecordsRouter) InitSysOperationRecordsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sysOperationRecordsRouter := Router.Group("sysOperationRecords")
	sysOperationRecordsRouterWithoutRecord := Router.Group("sysOperationRecords")
	sysOperationRecordsRouterWithoutAuth := PublicRouter.Group("sysOperationRecords")

	var sysOperationRecordsApi = v1.ApiGroupApp.HosApiGroup.SysOperationRecordsApi
	{
		sysOperationRecordsRouter.POST("createSysOperationRecords", sysOperationRecordsApi.CreateSysOperationRecords)             // 新建sysOperationRecords表
		sysOperationRecordsRouter.DELETE("deleteSysOperationRecords", sysOperationRecordsApi.DeleteSysOperationRecords)           // 删除sysOperationRecords表
		sysOperationRecordsRouter.DELETE("deleteSysOperationRecordsByIds", sysOperationRecordsApi.DeleteSysOperationRecordsByIds) // 批量删除sysOperationRecords表
		sysOperationRecordsRouter.PUT("updateSysOperationRecords", sysOperationRecordsApi.UpdateSysOperationRecords)              // 更新sysOperationRecords表
	}
	{
		sysOperationRecordsRouterWithoutRecord.GET("findSysOperationRecords", sysOperationRecordsApi.FindSysOperationRecords)       // 根据ID获取sysOperationRecords表
		sysOperationRecordsRouterWithoutRecord.GET("getSysOperationRecordsList", sysOperationRecordsApi.GetSysOperationRecordsList) // 获取sysOperationRecords表列表
	}
	{
		sysOperationRecordsRouterWithoutAuth.GET("getSysOperationRecordsPublic", sysOperationRecordsApi.GetSysOperationRecordsPublic) // 获取sysOperationRecords表列表
	}
}
