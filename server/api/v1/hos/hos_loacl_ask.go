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

type HosLoaclAskApi struct {
}

var hosLoaclAskService = service.ServiceGroupApp.HosServiceGroup.HosLoaclAskService

// CreateHosLoaclAsk 创建hosLoaclAsk表
// @Tags HosLoaclAsk
// @Summary 创建hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosLoaclAsk true "创建hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosLoaclAsk/createHosLoaclAsk [post]
func (hosLoaclAskApi *HosLoaclAskApi) CreateHosLoaclAsk(c *gin.Context) {
	var hosLoaclAsk hos.HosLoaclAsk
	err := c.ShouldBindJSON(&hosLoaclAsk)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosLoaclAskService.CreateHosLoaclAsk(&hosLoaclAsk, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHosLoaclAsk 删除hosLoaclAsk表
// @Tags HosLoaclAsk
// @Summary 删除hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosLoaclAsk true "删除hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosLoaclAsk/deleteHosLoaclAsk [delete]
func (hosLoaclAskApi *HosLoaclAskApi) DeleteHosLoaclAsk(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosLoaclAskService.DeleteHosLoaclAsk(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosLoaclAskByIds 批量删除hosLoaclAsk表
// @Tags HosLoaclAsk
// @Summary 批量删除hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosLoaclAsk/deleteHosLoaclAskByIds [delete]
func (hosLoaclAskApi *HosLoaclAskApi) DeleteHosLoaclAskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosLoaclAskService.DeleteHosLoaclAskByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosLoaclAsk 更新hosLoaclAsk表
// @Tags HosLoaclAsk
// @Summary 更新hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosLoaclAsk true "更新hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosLoaclAsk/updateHosLoaclAsk [put]
func (hosLoaclAskApi *HosLoaclAskApi) UpdateHosLoaclAsk(c *gin.Context) {
	var hosLoaclAsk hos.HosLoaclAsk
	err := c.ShouldBindJSON(&hosLoaclAsk)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosLoaclAskService.UpdateHosLoaclAsk(hosLoaclAsk, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosLoaclAsk 用id查询hosLoaclAsk表
// @Tags HosLoaclAsk
// @Summary 用id查询hosLoaclAsk表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosLoaclAsk true "用id查询hosLoaclAsk表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosLoaclAsk/findHosLoaclAsk [get]
func (hosLoaclAskApi *HosLoaclAskApi) FindHosLoaclAsk(c *gin.Context) {
	ID := c.Query("ID")
	if rehosLoaclAsk, err := hosLoaclAskService.GetHosLoaclAsk(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosLoaclAsk": rehosLoaclAsk}, c)
	}
}

// GetHosLoaclAskList 分页获取hosLoaclAsk表列表
// @Tags HosLoaclAsk
// @Summary 分页获取hosLoaclAsk表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosLoaclAskSearch true "分页获取hosLoaclAsk表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosLoaclAsk/getHosLoaclAskList [get]
func (hosLoaclAskApi *HosLoaclAskApi) GetHosLoaclAskList(c *gin.Context) {
	var pageInfo hosReq.HosLoaclAskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosLoaclAskService.GetHosLoaclAskInfoList(pageInfo, c); err != nil {
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

// GetHosLoaclAskPublic 不需要鉴权的hosLoaclAsk表接口
// @Tags HosLoaclAsk
// @Summary 不需要鉴权的hosLoaclAsk表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosLoaclAskSearch true "分页获取hosLoaclAsk表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosLoaclAsk/getHosLoaclAskPublic [get]
func (hosLoaclAskApi *HosLoaclAskApi) GetHosLoaclAskPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosLoaclAsk表接口信息",
	}, "获取成功", c)
}
