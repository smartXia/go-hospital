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

type HosUserPointApi struct {
}

var hosUserPointService = service.ServiceGroupApp.HosServiceGroup.HosUserPointService

// CreateHosUserPoint 创建hosUserPoint表
// @Tags HosUserPoint
// @Summary 创建hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUserPoint true "创建hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosUserPoint/createHosUserPoint [post]
func (hosUserPointApi *HosUserPointApi) CreateHosUserPoint(c *gin.Context) {
	var hosUserPoint hos.HosUserPoint
	err := c.ShouldBindJSON(&hosUserPoint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosUserPointService.CreateHosUserPoint(&hosUserPoint, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteHosUserPoint 删除hosUserPoint表
// @Tags HosUserPoint
// @Summary 删除hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUserPoint true "删除hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUserPoint/deleteHosUserPoint [delete]
func (hosUserPointApi *HosUserPointApi) DeleteHosUserPoint(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosUserPointService.DeleteHosUserPoint(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosUserPointByIds 批量删除hosUserPoint表
// @Tags HosUserPoint
// @Summary 批量删除hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosUserPoint/deleteHosUserPointByIds [delete]
func (hosUserPointApi *HosUserPointApi) DeleteHosUserPointByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosUserPointService.DeleteHosUserPointByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosUserPoint 更新hosUserPoint表
// @Tags HosUserPoint
// @Summary 更新hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUserPoint true "更新hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosUserPoint/updateHosUserPoint [put]
func (hosUserPointApi *HosUserPointApi) UpdateHosUserPoint(c *gin.Context) {
	var hosUserPoint hos.HosUserPoint
	err := c.ShouldBindJSON(&hosUserPoint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosUserPointService.UpdateHosUserPoint(hosUserPoint, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosUserPoint 用id查询hosUserPoint表
// @Tags HosUserPoint
// @Summary 用id查询hosUserPoint表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosUserPoint true "用id查询hosUserPoint表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosUserPoint/findHosUserPoint [get]
func (hosUserPointApi *HosUserPointApi) FindHosUserPoint(c *gin.Context) {
	ID := c.Query("ID")
	if rehosUserPoint, err := hosUserPointService.GetHosUserPoint(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosUserPoint": rehosUserPoint}, c)
	}
}

// GetHosUserPointList 分页获取hosUserPoint表列表
// @Tags HosUserPoint
// @Summary 分页获取hosUserPoint表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosUserPointSearch true "分页获取hosUserPoint表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUserPoint/getHosUserPointList [get]
func (hosUserPointApi *HosUserPointApi) GetHosUserPointList(c *gin.Context) {
	var pageInfo hosReq.HosUserPointSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosUserPointService.GetHosUserPointInfoList(pageInfo, c); err != nil {
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

// GetHosUserPointPublic 不需要鉴权的hosUserPoint表接口
// @Tags HosUserPoint
// @Summary 不需要鉴权的hosUserPoint表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosUserPointSearch true "分页获取hosUserPoint表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUserPoint/getHosUserPointPublic [get]
func (hosUserPointApi *HosUserPointApi) GetHosUserPointPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosUserPoint表接口信息",
	}, "获取成功", c)
}
