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

type HosSportClockApi struct {
}

var hosSportClockService = service.ServiceGroupApp.HosServiceGroup.HosSportClockService

// CreateHosSportClock 创建hosSportClock表
// @Tags HosSportClock
// @Summary 创建hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportClock true "创建hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosSportClock/createHosSportClock [post]
func (hosSportClockApi *HosSportClockApi) CreateHosSportClock(c *gin.Context) {
	var hosSportClock hos.HosSportClock
	err := c.ShouldBindJSON(&hosSportClock)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, d := hosSportClockService.CreateHosSportClock(&hosSportClock, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteHosSportClock 删除hosSportClock表
// @Tags HosSportClock
// @Summary 删除hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportClock true "删除hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportClock/deleteHosSportClock [delete]
func (hosSportClockApi *HosSportClockApi) DeleteHosSportClock(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosSportClockService.DeleteHosSportClock(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosSportClockByIds 批量删除hosSportClock表
// @Tags HosSportClock
// @Summary 批量删除hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosSportClock/deleteHosSportClockByIds [delete]
func (hosSportClockApi *HosSportClockApi) DeleteHosSportClockByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosSportClockService.DeleteHosSportClockByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosSportClock 更新hosSportClock表
// @Tags HosSportClock
// @Summary 更新hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportClock true "更新hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosSportClock/updateHosSportClock [put]
func (hosSportClockApi *HosSportClockApi) UpdateHosSportClock(c *gin.Context) {
	var hosSportClock hos.HosSportClock
	err := c.ShouldBindJSON(&hosSportClock)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosSportClockService.UpdateHosSportClock(hosSportClock, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosSportClock 用id查询hosSportClock表
// @Tags HosSportClock
// @Summary 用id查询hosSportClock表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosSportClock true "用id查询hosSportClock表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosSportClock/findHosSportClock [get]
func (hosSportClockApi *HosSportClockApi) FindHosSportClock(c *gin.Context) {
	ID := c.Query("ID")
	if rehosSportClock, err := hosSportClockService.GetHosSportClock(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosSportClock": rehosSportClock}, c)
	}
}

// GetHosSportClockList 分页获取hosSportClock表列表
// @Tags HosSportClock
// @Summary 分页获取hosSportClock表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportClockSearch true "分页获取hosSportClock表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportClock/getHosSportClockList [get]
func (hosSportClockApi *HosSportClockApi) GetHosSportClockList(c *gin.Context) {
	var pageInfo hosReq.HosSportClockSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosSportClockService.GetHosSportClockInfoList(pageInfo, c); err != nil {
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

func (hosSportClockApi *HosSportClockApi) GetHosSportDistinctClockList(c *gin.Context) {
	var pageInfo hosReq.HosSportClockSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosSportClockService.GetHosSportDistinctClockList(pageInfo, c); err != nil {
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

// GetHosSportClockDataSource 获取HosSportClock的数据源
// @Tags HosSportClock
// @Summary 获取HosSportClock的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportClock/getHosSportClockDataSource [get]
func (hosSportClockApi *HosSportClockApi) GetHosSportClockDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := hosSportClockService.GetHosSportClockDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetHosSportClockPublic 不需要鉴权的hosSportClock表接口
// @Tags HosSportClock
// @Summary 不需要鉴权的hosSportClock表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportClockSearch true "分页获取hosSportClock表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportClock/getHosSportClockPublic [get]
func (hosSportClockApi *HosSportClockApi) GetHosSportClockPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosSportClock表接口信息",
	}, "获取成功", c)
}

// GetHosSportClockList 分页获取hosSportClock表列表
// @Tags HosSportClock
// @Summary 分页获取hosSportClock表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportClockSearch true "分页获取hosSportClock表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportClock/getHosSportClockList [get]
func (hosSportClockApi *HosSportClockApi) GetCurrentUserHosSportClockList(c *gin.Context) {
	var pageInfo hosReq.HosSportClockSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosSportClockService.GetCurrentUserHosSportClockList(pageInfo, c); err != nil {
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
