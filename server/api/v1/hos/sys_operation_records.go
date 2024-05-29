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

type SysOperationRecordsApi struct {
}

var sysOperationRecordsService = service.ServiceGroupApp.HosServiceGroup.SysOperationRecordsService

// CreateSysOperationRecords 创建sysOperationRecords表
// @Tags SysOperationRecords
// @Summary 创建sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysOperationRecords true "创建sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysOperationRecords/createSysOperationRecords [post]
func (sysOperationRecordsApi *SysOperationRecordsApi) CreateSysOperationRecords(c *gin.Context) {
	var sysOperationRecords hos.SysOperationRecords
	err := c.ShouldBindJSON(&sysOperationRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysOperationRecordsService.CreateSysOperationRecords(&sysOperationRecords, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysOperationRecords 删除sysOperationRecords表
// @Tags SysOperationRecords
// @Summary 删除sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysOperationRecords true "删除sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecords/deleteSysOperationRecords [delete]
func (sysOperationRecordsApi *SysOperationRecordsApi) DeleteSysOperationRecords(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysOperationRecordsService.DeleteSysOperationRecords(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysOperationRecordsByIds 批量删除sysOperationRecords表
// @Tags SysOperationRecords
// @Summary 批量删除sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysOperationRecords/deleteSysOperationRecordsByIds [delete]
func (sysOperationRecordsApi *SysOperationRecordsApi) DeleteSysOperationRecordsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysOperationRecordsService.DeleteSysOperationRecordsByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysOperationRecords 更新sysOperationRecords表
// @Tags SysOperationRecords
// @Summary 更新sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysOperationRecords true "更新sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysOperationRecords/updateSysOperationRecords [put]
func (sysOperationRecordsApi *SysOperationRecordsApi) UpdateSysOperationRecords(c *gin.Context) {
	var sysOperationRecords hos.SysOperationRecords
	err := c.ShouldBindJSON(&sysOperationRecords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysOperationRecordsService.UpdateSysOperationRecords(sysOperationRecords, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysOperationRecords 用id查询sysOperationRecords表
// @Tags SysOperationRecords
// @Summary 用id查询sysOperationRecords表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.SysOperationRecords true "用id查询sysOperationRecords表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecords/findSysOperationRecords [get]
func (sysOperationRecordsApi *SysOperationRecordsApi) FindSysOperationRecords(c *gin.Context) {
	ID := c.Query("ID")
	if resysOperationRecords, err := sysOperationRecordsService.GetSysOperationRecords(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysOperationRecords": resysOperationRecords}, c)
	}
}

// GetSysOperationRecordsList 分页获取sysOperationRecords表列表
// @Tags SysOperationRecords
// @Summary 分页获取sysOperationRecords表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysOperationRecordsSearch true "分页获取sysOperationRecords表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecords/getSysOperationRecordsList [get]
func (sysOperationRecordsApi *SysOperationRecordsApi) GetSysOperationRecordsList(c *gin.Context) {
	var pageInfo hosReq.SysOperationRecordsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysOperationRecordsService.GetSysOperationRecordsInfoList(pageInfo, c); err != nil {
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

// GetSysOperationRecordsPublic 不需要鉴权的sysOperationRecords表接口
// @Tags SysOperationRecords
// @Summary 不需要鉴权的sysOperationRecords表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysOperationRecordsSearch true "分页获取sysOperationRecords表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecords/getSysOperationRecordsPublic [get]
func (sysOperationRecordsApi *SysOperationRecordsApi) GetSysOperationRecordsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的sysOperationRecords表接口信息",
	}, "获取成功", c)
}
