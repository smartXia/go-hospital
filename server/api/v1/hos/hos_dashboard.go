package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/response"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HosDashboardApi struct {
}

var hosDashboardService = service.ServiceGroupApp.HosServiceGroup.HosDashboardService

// CreateHosDashboard 创建hosDashboard表
// @Tags HosDashboard
// @Summary 创建hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosDashboard true "创建hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosDashboard/createHosDashboard [post]
func (hosDashboardApi *HosDashboardApi) CreateHosDashboard(c *gin.Context) {
	var hosDashboard hos.HosDashboard
	err := c.ShouldBindJSON(&hosDashboard)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, d := hosDashboardService.CreateHosDashboard(&hosDashboard, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteHosDashboard 删除hosDashboard表
// @Tags HosDashboard
// @Summary 删除hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosDashboard true "删除hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosDashboard/deleteHosDashboard [delete]
func (hosDashboardApi *HosDashboardApi) DeleteHosDashboard(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosDashboardService.DeleteHosDashboard(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosDashboardByIds 批量删除hosDashboard表
// @Tags HosDashboard
// @Summary 批量删除hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosDashboard/deleteHosDashboardByIds [delete]
func (hosDashboardApi *HosDashboardApi) DeleteHosDashboardByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosDashboardService.DeleteHosDashboardByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosDashboard 更新hosDashboard表
// @Tags HosDashboard
// @Summary 更新hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosDashboard true "更新hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosDashboard/updateHosDashboard [put]
func (hosDashboardApi *HosDashboardApi) UpdateHosDashboard(c *gin.Context) {
	var hosDashboard hos.HosDashboard
	err := c.ShouldBindJSON(&hosDashboard)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosDashboardService.UpdateHosDashboard(hosDashboard, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosDashboard 用id查询hosDashboard表
// @Tags HosDashboard
// @Summary 用id查询hosDashboard表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosDashboard true "用id查询hosDashboard表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosDashboard/findHosDashboard [get]
func (hosDashboardApi *HosDashboardApi) FindHosDashboard(c *gin.Context) {
	ID := c.Query("ID")
	if rehosDashboard, err := hosDashboardService.GetHosDashboard(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosDashboard": rehosDashboard}, c)
	}
}

// GetHosDashboardList 分页获取hosDashboard表列表
// @Tags HosDashboard
// @Summary 分页获取hosDashboard表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosDashboardSearch true "分页获取hosDashboard表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosDashboard/getHosDashboardList [get]
func (hosDashboardApi *HosDashboardApi) GetHosDashboardList(c *gin.Context) {
	var pageInfo hosReq.HosDashboardSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosDashboardService.GetHosDashboardInfoList(pageInfo, c); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetHosDashboardDataSource 获取HosDashboard的数据源
// @Tags HosDashboard
// @Summary 获取HosDashboard的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosDashboard/getHosDashboardDataSource [get]
func (hosDashboardApi *HosDashboardApi) GetHosDashboardDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := hosDashboardService.GetHosDashboardDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetHosDashboardPublic 不需要鉴权的hosDashboard表接口
// @Tags HosDashboard
// @Summary 不需要鉴权的hosDashboard表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosDashboardSearch true "分页获取hosDashboard表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosDashboard/getHosDashboardPublic [get]
func (hosDashboardApi *HosDashboardApi) GetHosDashboardPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosDashboard表接口信息",
	}, "获取成功", c)
}
