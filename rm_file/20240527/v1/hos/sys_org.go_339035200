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

type SysOrgApi struct {
}

var sysOrgService = service.ServiceGroupApp.HosServiceGroup.SysOrgService

// CreateSysOrg 创建sysOrg表
// @Tags SysOrg
// @Summary 创建sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysOrg true "创建sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysOrg/createSysOrg [post]
func (sysOrgApi *SysOrgApi) CreateSysOrg(c *gin.Context) {
	var sysOrg hos.SysOrg
	err := c.ShouldBindJSON(&sysOrg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysOrgService.CreateSysOrg(&sysOrg); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysOrg 删除sysOrg表
// @Tags SysOrg
// @Summary 删除sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysOrg true "删除sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOrg/deleteSysOrg [delete]
func (sysOrgApi *SysOrgApi) DeleteSysOrg(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysOrgService.DeleteSysOrg(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysOrgByIds 批量删除sysOrg表
// @Tags SysOrg
// @Summary 批量删除sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysOrg/deleteSysOrgByIds [delete]
func (sysOrgApi *SysOrgApi) DeleteSysOrgByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysOrgService.DeleteSysOrgByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysOrg 更新sysOrg表
// @Tags SysOrg
// @Summary 更新sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysOrg true "更新sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysOrg/updateSysOrg [put]
func (sysOrgApi *SysOrgApi) UpdateSysOrg(c *gin.Context) {
	var sysOrg hos.SysOrg
	err := c.ShouldBindJSON(&sysOrg)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysOrgService.UpdateSysOrg(sysOrg); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysOrg 用id查询sysOrg表
// @Tags SysOrg
// @Summary 用id查询sysOrg表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.SysOrg true "用id查询sysOrg表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOrg/findSysOrg [get]
func (sysOrgApi *SysOrgApi) FindSysOrg(c *gin.Context) {
	ID := c.Query("ID")
	if resysOrg, err := sysOrgService.GetSysOrg(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysOrg": resysOrg}, c)
	}
}

// GetSysOrgList 分页获取sysOrg表列表
// @Tags SysOrg
// @Summary 分页获取sysOrg表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysOrgSearch true "分页获取sysOrg表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOrg/getSysOrgList [get]
func (sysOrgApi *SysOrgApi) GetSysOrgList(c *gin.Context) {
	var pageInfo hosReq.SysOrgSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysOrgService.GetSysOrgInfoList(pageInfo); err != nil {
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

// GetSysOrgPublic 不需要鉴权的sysOrg表接口
// @Tags SysOrg
// @Summary 不需要鉴权的sysOrg表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysOrgSearch true "分页获取sysOrg表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOrg/getSysOrgPublic [get]
func (sysOrgApi *SysOrgApi) GetSysOrgPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的sysOrg表接口信息",
	}, "获取成功", c)
}
