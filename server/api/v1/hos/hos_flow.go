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

type HosFlowApi struct {
}

var hosFlowService = service.ServiceGroupApp.HosServiceGroup.HosFlowService

// CreateHosFlow 创建hosFlow表
// @Tags HosFlow
// @Summary 创建hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosFlow true "创建hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosFlow/createHosFlow [post]
func (hosFlowApi *HosFlowApi) CreateHosFlow(c *gin.Context) {
	var hosFlow hos.HosFlow
	err := c.ShouldBindJSON(&hosFlow)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, h := hosFlowService.CreateHosFlow(&hosFlow, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(h, c)
	}
}

// DeleteHosFlow 删除hosFlow表
// @Tags HosFlow
// @Summary 删除hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosFlow true "删除hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosFlow/deleteHosFlow [delete]
func (hosFlowApi *HosFlowApi) DeleteHosFlow(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosFlowService.DeleteHosFlow(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosFlowByIds 批量删除hosFlow表
// @Tags HosFlow
// @Summary 批量删除hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosFlow/deleteHosFlowByIds [delete]
func (hosFlowApi *HosFlowApi) DeleteHosFlowByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosFlowService.DeleteHosFlowByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosFlow 更新hosFlow表
// @Tags HosFlow
// @Summary 更新hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosFlow true "更新hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosFlow/updateHosFlow [put]
func (hosFlowApi *HosFlowApi) UpdateHosFlow(c *gin.Context) {
	var hosFlow hos.HosFlow
	err := c.ShouldBindJSON(&hosFlow)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosFlowService.UpdateHosFlow(hosFlow, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosFlow 用id查询hosFlow表
// @Tags HosFlow
// @Summary 用id查询hosFlow表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosFlow true "用id查询hosFlow表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosFlow/findHosFlow [get]
func (hosFlowApi *HosFlowApi) FindHosFlow(c *gin.Context) {
	ID := c.Query("ID")
	if rehosFlow, err := hosFlowService.GetHosFlow(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosFlow": rehosFlow}, c)
	}
}

// GetHosFlowList 分页获取hosFlow表列表
// @Tags HosFlow
// @Summary 分页获取hosFlow表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosFlowSearch true "分页获取hosFlow表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosFlow/getHosFlowList [get]
func (hosFlowApi *HosFlowApi) GetHosFlowList(c *gin.Context) {
	var pageInfo hosReq.HosFlowSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosFlowService.GetHosFlowInfoList(pageInfo, c); err != nil {
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

func (hosFlowApi *HosFlowApi) GetCurrentHosFlowList(c *gin.Context) {
	var pageInfo hosReq.HosFlowSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosFlowService.GetHosFlowInfoList(pageInfo, c); err != nil {
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

// GetHosFlowDataSource 获取HosFlow的数据源
// @Tags HosFlow
// @Summary 获取HosFlow的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosFlow/getHosFlowDataSource [get]
func (hosFlowApi *HosFlowApi) GetHosFlowDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := hosFlowService.GetHosFlowDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetHosFlowPublic 不需要鉴权的hosFlow表接口
// @Tags HosFlow
// @Summary 不需要鉴权的hosFlow表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosFlowSearch true "分页获取hosFlow表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosFlow/getHosFlowPublic [get]
func (hosFlowApi *HosFlowApi) GetHosFlowPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosFlow表接口信息",
	}, "获取成功", c)
}
