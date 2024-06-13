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

type HosScaleApi struct {
}

var hosScaleService = service.ServiceGroupApp.HosServiceGroup.HosScaleService

// CreateHosScale 创建hosScale表
// @Tags HosScale
// @Summary 创建hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosScale true "创建hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosScale/createHosScale [post]
func (hosScaleApi *HosScaleApi) CreateHosScale(c *gin.Context) {
	var hosScale hos.HosScale
	err := c.ShouldBindJSON(&hosScale)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, d := hosScaleService.CreateHosScale(&hosScale, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteHosScale 删除hosScale表
// @Tags HosScale
// @Summary 删除hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosScale true "删除hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosScale/deleteHosScale [delete]
func (hosScaleApi *HosScaleApi) DeleteHosScale(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosScaleService.DeleteHosScale(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosScaleByIds 批量删除hosScale表
// @Tags HosScale
// @Summary 批量删除hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosScale/deleteHosScaleByIds [delete]
func (hosScaleApi *HosScaleApi) DeleteHosScaleByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosScaleService.DeleteHosScaleByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosScale 更新hosScale表
// @Tags HosScale
// @Summary 更新hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosScale true "更新hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosScale/updateHosScale [put]
func (hosScaleApi *HosScaleApi) UpdateHosScale(c *gin.Context) {
	var hosScale hos.HosScale
	err := c.ShouldBindJSON(&hosScale)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosScaleService.UpdateHosScale(hosScale, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosScale 用id查询hosScale表
// @Tags HosScale
// @Summary 用id查询hosScale表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosScale true "用id查询hosScale表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosScale/findHosScale [get]
func (hosScaleApi *HosScaleApi) FindHosScale(c *gin.Context) {
	ID := c.Query("ID")
	if rehosScale, err := hosScaleService.GetHosScale(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosScale": rehosScale}, c)
	}
}

// GetHosScaleList 分页获取hosScale表列表
// @Tags HosScale
// @Summary 分页获取hosScale表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosScaleSearch true "分页获取hosScale表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosScale/getHosScaleList [get]
func (hosScaleApi *HosScaleApi) GetHosScaleList(c *gin.Context) {
	var pageInfo hosReq.HosScaleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosScaleService.GetHosScaleInfoList(pageInfo, c); err != nil {
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

// GetHosScalePublic 不需要鉴权的hosScale表接口
// @Tags HosScale
// @Summary 不需要鉴权的hosScale表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosScaleSearch true "分页获取hosScale表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosScale/getHosScalePublic [get]
func (hosScaleApi *HosScaleApi) GetHosScalePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosScale表接口信息",
	}, "获取成功", c)
}
