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

type HosLocalAskApi struct {
}

var hosLocalAskService = service.ServiceGroupApp.HosServiceGroup.HosLocalAskService

// CreateHosLocalAsk 创建hosLocalAsk表
// @Tags HosLocalAsk
// @Summary 创建hosLocalAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosLocalAsk true "创建hosLocalAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosLocalAsk/createHosLocalAsk [post]
func (hosLocalAskApi *HosLocalAskApi) CreateHosLocalAsk(c *gin.Context) {
	var hosLocalAsk hos.HosLocalAsk
	err := c.ShouldBindJSON(&hosLocalAsk)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, h := hosLocalAskService.CreateHosLocalAsk(&hosLocalAsk, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(h, c)
	}
}

// DeleteHosLocalAsk 删除hosLocalAsk表
// @Tags HosLocalAsk
// @Summary 删除hosLocalAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosLocalAsk true "删除hosLocalAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosLocalAsk/deleteHosLocalAsk [delete]
func (hosLocalAskApi *HosLocalAskApi) DeleteHosLocalAsk(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosLocalAskService.DeleteHosLocalAsk(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosLocalAskByIds 批量删除hosLocalAsk表
// @Tags HosLocalAsk
// @Summary 批量删除hosLocalAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosLocalAsk/deleteHosLocalAskByIds [delete]
func (hosLocalAskApi *HosLocalAskApi) DeleteHosLocalAskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosLocalAskService.DeleteHosLocalAskByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosLocalAsk 更新hosLocalAsk表
// @Tags HosLocalAsk
// @Summary 更新hosLocalAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosLocalAsk true "更新hosLocalAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosLocalAsk/updateHosLocalAsk [put]
func (hosLocalAskApi *HosLocalAskApi) UpdateHosLocalAsk(c *gin.Context) {
	var hosLocalAsk hos.HosLocalAsk
	err := c.ShouldBindJSON(&hosLocalAsk)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosLocalAskService.UpdateHosLocalAsk(hosLocalAsk, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosLocalAsk 用id查询hosLocalAsk表
// @Tags HosLocalAsk
// @Summary 用id查询hosLocalAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosLocalAsk true "用id查询hosLocalAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosLocalAsk/findHosLocalAsk [get]
func (hosLocalAskApi *HosLocalAskApi) FindHosLocalAsk(c *gin.Context) {
	ID := c.Query("ID")
	if rehosLocalAsk, err := hosLocalAskService.GetHosLocalAsk(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosLocalAsk": rehosLocalAsk}, c)
	}
}

// GetHosLocalAskList 分页获取hosLocalAsk表列表
// @Tags HosLocalAsk
// @Summary 分页获取hosLocalAsk表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosLocalAskSearch true "分页获取hosLocalAsk表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosLocalAsk/getHosLocalAskList [get]
func (hosLocalAskApi *HosLocalAskApi) GetHosLocalAskList(c *gin.Context) {
	var pageInfo hosReq.HosLocalAskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosLocalAskService.GetHosLocalAskInfoList(pageInfo, c); err != nil {
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

// GetHosLocalAskPublic 不需要鉴权的hosLocalAsk表接口
// @Tags HosLocalAsk
// @Summary 不需要鉴权的hosLocalAsk表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosLocalAskSearch true "分页获取hosLocalAsk表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosLocalAsk/getHosLocalAskPublic [get]
func (hosLocalAskApi *HosLocalAskApi) GetHosLocalAskPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosLocalAsk表接口信息",
	}, "获取成功", c)
}
