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

type HosUsersDianzanApi struct {
}

var hosUsersDianzanService = service.ServiceGroupApp.HosServiceGroup.HosUsersDianzanService

// CreateHosUsersDianzan 创建hosUsersDianzan表
// @Tags HosUsersDianzan
// @Summary 创建hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUsersDianzan true "创建hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosUsersDianzan/createHosUsersDianzan [post]
func (hosUsersDianzanApi *HosUsersDianzanApi) CreateHosUsersDianzan(c *gin.Context) {
	var hosUsersDianzan hos.HosUsersDianzan
	err := c.ShouldBindJSON(&hosUsersDianzan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, d := hosUsersDianzanService.CreateHosUsersDianzan(&hosUsersDianzan, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteHosUsersDianzan 删除hosUsersDianzan表
// @Tags HosUsersDianzan
// @Summary 删除hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUsersDianzan true "删除hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUsersDianzan/deleteHosUsersDianzan [delete]
func (hosUsersDianzanApi *HosUsersDianzanApi) DeleteHosUsersDianzan(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosUsersDianzanService.DeleteHosUsersDianzan(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosUsersDianzanByIds 批量删除hosUsersDianzan表
// @Tags HosUsersDianzan
// @Summary 批量删除hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosUsersDianzan/deleteHosUsersDianzanByIds [delete]
func (hosUsersDianzanApi *HosUsersDianzanApi) DeleteHosUsersDianzanByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosUsersDianzanService.DeleteHosUsersDianzanByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosUsersDianzan 更新hosUsersDianzan表
// @Tags HosUsersDianzan
// @Summary 更新hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUsersDianzan true "更新hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosUsersDianzan/updateHosUsersDianzan [put]
func (hosUsersDianzanApi *HosUsersDianzanApi) UpdateHosUsersDianzan(c *gin.Context) {
	var hosUsersDianzan hos.HosUsersDianzan
	err := c.ShouldBindJSON(&hosUsersDianzan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosUsersDianzanService.UpdateHosUsersDianzan(hosUsersDianzan, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosUsersDianzan 用id查询hosUsersDianzan表
// @Tags HosUsersDianzan
// @Summary 用id查询hosUsersDianzan表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosUsersDianzan true "用id查询hosUsersDianzan表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosUsersDianzan/findHosUsersDianzan [get]
func (hosUsersDianzanApi *HosUsersDianzanApi) FindHosUsersDianzan(c *gin.Context) {
	ID := c.Query("ID")
	if rehosUsersDianzan, err := hosUsersDianzanService.GetHosUsersDianzan(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosUsersDianzan": rehosUsersDianzan}, c)
	}
}

// GetHosUsersDianzanList 分页获取hosUsersDianzan表列表
// @Tags HosUsersDianzan
// @Summary 分页获取hosUsersDianzan表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosUsersDianzanSearch true "分页获取hosUsersDianzan表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUsersDianzan/getHosUsersDianzanList [get]
func (hosUsersDianzanApi *HosUsersDianzanApi) GetHosUsersDianzanList(c *gin.Context) {
	var pageInfo hosReq.HosUsersDianzanSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosUsersDianzanService.GetHosUsersDianzanInfoList(pageInfo, c); err != nil {
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

// GetHosUsersDianzanPublic 不需要鉴权的hosUsersDianzan表接口
// @Tags HosUsersDianzan
// @Summary 不需要鉴权的hosUsersDianzan表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosUsersDianzanSearch true "分页获取hosUsersDianzan表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUsersDianzan/getHosUsersDianzanPublic [get]
func (hosUsersDianzanApi *HosUsersDianzanApi) GetHosUsersDianzanPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosUsersDianzan表接口信息",
	}, "获取成功", c)
}
