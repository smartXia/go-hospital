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

type HosSportModeApi struct {
}

var hosSportModeService = service.ServiceGroupApp.HosServiceGroup.HosSportModeService

// CreateHosSportMode 创建hosSportMode表
// @Tags HosSportMode
// @Summary 创建hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportMode true "创建hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosSportMode/createHosSportMode [post]
func (hosSportModeApi *HosSportModeApi) CreateHosSportMode(c *gin.Context) {
	var hosSportMode hos.HosSportMode
	err := c.ShouldBindJSON(&hosSportMode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, d := hosSportModeService.CreateHosSportMode(&hosSportMode, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteHosSportMode 删除hosSportMode表
// @Tags HosSportMode
// @Summary 删除hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportMode true "删除hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportMode/deleteHosSportMode [delete]
func (hosSportModeApi *HosSportModeApi) DeleteHosSportMode(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosSportModeService.DeleteHosSportMode(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosSportModeByIds 批量删除hosSportMode表
// @Tags HosSportMode
// @Summary 批量删除hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosSportMode/deleteHosSportModeByIds [delete]
func (hosSportModeApi *HosSportModeApi) DeleteHosSportModeByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosSportModeService.DeleteHosSportModeByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosSportMode 更新hosSportMode表
// @Tags HosSportMode
// @Summary 更新hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportMode true "更新hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosSportMode/updateHosSportMode [put]
func (hosSportModeApi *HosSportModeApi) UpdateHosSportMode(c *gin.Context) {
	var hosSportMode hos.HosSportMode
	err := c.ShouldBindJSON(&hosSportMode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosSportModeService.UpdateHosSportMode(hosSportMode, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosSportMode 用id查询hosSportMode表
// @Tags HosSportMode
// @Summary 用id查询hosSportMode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosSportMode true "用id查询hosSportMode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosSportMode/findHosSportMode [get]
func (hosSportModeApi *HosSportModeApi) FindHosSportMode(c *gin.Context) {
	ID := c.Query("ID")
	if rehosSportMode, err := hosSportModeService.GetHosSportMode(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosSportMode": rehosSportMode}, c)
	}
}

// GetHosSportModeList 分页获取hosSportMode表列表
// @Tags HosSportMode
// @Summary 分页获取hosSportMode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportModeSearch true "分页获取hosSportMode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportMode/getHosSportModeList [get]
func (hosSportModeApi *HosSportModeApi) GetHosSportModeList(c *gin.Context) {
	var pageInfo hosReq.HosSportModeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosSportModeService.GetHosSportModeInfoList(pageInfo, c); err != nil {
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

// GetHosSportModePublic 不需要鉴权的hosSportMode表接口
// @Tags HosSportMode
// @Summary 不需要鉴权的hosSportMode表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportModeSearch true "分页获取hosSportMode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportMode/getHosSportModePublic [get]
func (hosSportModeApi *HosSportModeApi) GetHosSportModePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosSportMode表接口信息",
	}, "获取成功", c)
}

// GetHosSportModeMatrix 分页获取hosSportMode表列表
// @Tags HosSportMode
// @Summary 分页获取hosSportMode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportModeSearch true "分页获取hosSportMode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportMode/getHosSportModeList [get]
func (hosSportModeApi *HosSportModeApi) GetHosSportModeMatrix(c *gin.Context) {
	var pageInfo hosReq.HosSportModeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosSportModeService.GetHosSportModeMatrix(pageInfo, c); err != nil {
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
