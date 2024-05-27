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

type SysUsersApi struct {
}

var sysUsersService = service.ServiceGroupApp.HosServiceGroup.SysUsersService

// CreateSysUsers 创建sysUsers表
// @Tags SysUsers
// @Summary 创建sysUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysUsers true "创建sysUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysUsers/createSysUsers [post]
func (sysUsersApi *SysUsersApi) CreateSysUsers(c *gin.Context) {
	var sysUsers hos.SysUsers
	err := c.ShouldBindJSON(&sysUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysUsersService.CreateSysUsers(&sysUsers, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysUsers 删除sysUsers表
// @Tags SysUsers
// @Summary 删除sysUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysUsers true "删除sysUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysUsers/deleteSysUsers [delete]
func (sysUsersApi *SysUsersApi) DeleteSysUsers(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysUsersService.DeleteSysUsers(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysUsersByIds 批量删除sysUsers表
// @Tags SysUsers
// @Summary 批量删除sysUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysUsers/deleteSysUsersByIds [delete]
func (sysUsersApi *SysUsersApi) DeleteSysUsersByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysUsersService.DeleteSysUsersByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysUsers 更新sysUsers表
// @Tags SysUsers
// @Summary 更新sysUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysUsers true "更新sysUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysUsers/updateSysUsers [put]
func (sysUsersApi *SysUsersApi) UpdateSysUsers(c *gin.Context) {
	var sysUsers hos.SysUsers
	err := c.ShouldBindJSON(&sysUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysUsersService.UpdateSysUsers(sysUsers, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysUsers 用id查询sysUsers表
// @Tags SysUsers
// @Summary 用id查询sysUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.SysUsers true "用id查询sysUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysUsers/findSysUsers [get]
func (sysUsersApi *SysUsersApi) FindSysUsers(c *gin.Context) {
	ID := c.Query("ID")
	if resysUsers, err := sysUsersService.GetSysUsers(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysUsers": resysUsers}, c)
	}
}

// GetSysUsersList 分页获取sysUsers表列表
// @Tags SysUsers
// @Summary 分页获取sysUsers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysUsersSearch true "分页获取sysUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUsers/getSysUsersList [get]
func (sysUsersApi *SysUsersApi) GetSysUsersList(c *gin.Context) {
	var pageInfo hosReq.SysUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysUsersService.GetSysUsersInfoList(pageInfo, c); err != nil {
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

// GetSysUsersPublic 不需要鉴权的sysUsers表接口
// @Tags SysUsers
// @Summary 不需要鉴权的sysUsers表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysUsersSearch true "分页获取sysUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysUsers/getSysUsersPublic [get]
func (sysUsersApi *SysUsersApi) GetSysUsersPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的sysUsers表接口信息",
	}, "获取成功", c)
}
