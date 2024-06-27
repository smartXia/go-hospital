package hos

import (
	"devops-manage/api/v1"
	"github.com/gin-gonic/gin"
)

type HosSportClockCommitRouter struct {
}

// InitHosSportClockCommitRouter 初始化 hosSportClockCommit表 路由信息
func (s *HosSportClockCommitRouter) InitHosSportClockCommitRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	hosSportClockCommitRouter := Router.Group("hosSportClockCommit")
	hosSportClockCommitRouterWithoutRecord := Router.Group("hosSportClockCommit")
	hosSportClockCommitRouterWithoutAuth := PublicRouter.Group("hosSportClockCommit")

	var hosSportClockCommitApi = v1.ApiGroupApp.HosApiGroup.HosSportClockCommitApi
	{
		hosSportClockCommitRouter.POST("createHosSportClockCommit", hosSportClockCommitApi.CreateHosSportClockCommit)             // 新建hosSportClockCommit表
		hosSportClockCommitRouter.DELETE("deleteHosSportClockCommit", hosSportClockCommitApi.DeleteHosSportClockCommit)           // 删除hosSportClockCommit表
		hosSportClockCommitRouter.DELETE("deleteHosSportClockCommitByIds", hosSportClockCommitApi.DeleteHosSportClockCommitByIds) // 批量删除hosSportClockCommit表
		hosSportClockCommitRouter.PUT("updateHosSportClockCommit", hosSportClockCommitApi.UpdateHosSportClockCommit)              // 更新hosSportClockCommit表
	}
	{
		hosSportClockCommitRouterWithoutRecord.GET("findHosSportClockCommit", hosSportClockCommitApi.FindHosSportClockCommit)       // 根据ID获取hosSportClockCommit表
		hosSportClockCommitRouterWithoutRecord.GET("getHosSportClockCommitList", hosSportClockCommitApi.GetHosSportClockCommitList) // 获取hosSportClockCommit表列表
	}
	{
		hosSportClockCommitRouterWithoutAuth.GET("getHosSportClockCommitPublic", hosSportClockCommitApi.GetHosSportClockCommitPublic) // 获取hosSportClockCommit表列表
	}
}
