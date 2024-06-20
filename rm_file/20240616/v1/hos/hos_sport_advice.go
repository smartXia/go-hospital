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

type HosSportAdviceApi struct {
}

var hosSportAdviceService = service.ServiceGroupApp.HosServiceGroup.HosSportAdviceService

// CreateHosSportAdvice 创建hosSportAdvice表
// @Tags HosSportAdvice
// @Summary 创建hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportAdvice true "创建hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosSportAdvice/createHosSportAdvice [post]
func (hosSportAdviceApi *HosSportAdviceApi) CreateHosSportAdvice(c *gin.Context) {
	var hosSportAdvice hos.HosSportAdvice
	err := c.ShouldBindJSON(&hosSportAdvice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, d := hosSportAdviceService.CreateHosSportAdvice(&hosSportAdvice, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteHosSportAdvice 删除hosSportAdvice表
// @Tags HosSportAdvice
// @Summary 删除hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportAdvice true "删除hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportAdvice/deleteHosSportAdvice [delete]
func (hosSportAdviceApi *HosSportAdviceApi) DeleteHosSportAdvice(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosSportAdviceService.DeleteHosSportAdvice(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosSportAdviceByIds 批量删除hosSportAdvice表
// @Tags HosSportAdvice
// @Summary 批量删除hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosSportAdvice/deleteHosSportAdviceByIds [delete]
func (hosSportAdviceApi *HosSportAdviceApi) DeleteHosSportAdviceByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosSportAdviceService.DeleteHosSportAdviceByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosSportAdvice 更新hosSportAdvice表
// @Tags HosSportAdvice
// @Summary 更新hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportAdvice true "更新hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosSportAdvice/updateHosSportAdvice [put]
func (hosSportAdviceApi *HosSportAdviceApi) UpdateHosSportAdvice(c *gin.Context) {
	var hosSportAdvice hos.HosSportAdvice
	err := c.ShouldBindJSON(&hosSportAdvice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosSportAdviceService.UpdateHosSportAdvice(hosSportAdvice, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosSportAdvice 用id查询hosSportAdvice表
// @Tags HosSportAdvice
// @Summary 用id查询hosSportAdvice表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosSportAdvice true "用id查询hosSportAdvice表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosSportAdvice/findHosSportAdvice [get]
func (hosSportAdviceApi *HosSportAdviceApi) FindHosSportAdvice(c *gin.Context) {
	ID := c.Query("ID")
	if rehosSportAdvice, err := hosSportAdviceService.GetHosSportAdvice(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosSportAdvice": rehosSportAdvice}, c)
	}
}

// GetHosSportAdviceList 分页获取hosSportAdvice表列表
// @Tags HosSportAdvice
// @Summary 分页获取hosSportAdvice表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportAdviceSearch true "分页获取hosSportAdvice表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportAdvice/getHosSportAdviceList [get]
func (hosSportAdviceApi *HosSportAdviceApi) GetHosSportAdviceList(c *gin.Context) {
	var pageInfo hosReq.HosSportAdviceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosSportAdviceService.GetHosSportAdviceInfoList(pageInfo, c); err != nil {
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

// GetHosSportAdvicePublic 不需要鉴权的hosSportAdvice表接口
// @Tags HosSportAdvice
// @Summary 不需要鉴权的hosSportAdvice表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportAdviceSearch true "分页获取hosSportAdvice表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportAdvice/getHosSportAdvicePublic [get]
func (hosSportAdviceApi *HosSportAdviceApi) GetHosSportAdvicePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosSportAdvice表接口信息",
	}, "获取成功", c)
}
