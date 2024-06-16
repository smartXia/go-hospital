package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/response"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/service"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HosUsersApi struct {
}

var hosUsersService = service.ServiceGroupApp.HosServiceGroup.HosUsersService

// CreateHosUsers 创建hosUsers表
// @Tags HosUsers
// @Summary 创建hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUsers true "创建hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosUsers/createHosUsers [post]
func (hosUsersApi *HosUsersApi) CreateHosUsers(c *gin.Context) {
	var hosUsers hos.HosUsers
	err := c.ShouldBindJSON(&hosUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, d := hosUsersService.CreateHosUsers(&hosUsers, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteHosUsers 删除hosUsers表
// @Tags HosUsers
// @Summary 删除hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUsers true "删除hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosUsers/deleteHosUsers [delete]
func (hosUsersApi *HosUsersApi) DeleteHosUsers(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := hosUsersService.DeleteHosUsers(ID, userID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosUsersByIds 批量删除hosUsers表
// @Tags HosUsers
// @Summary 批量删除hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosUsers/deleteHosUsersByIds [delete]
func (hosUsersApi *HosUsersApi) DeleteHosUsersByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := hosUsersService.DeleteHosUsersByIds(IDs, userID, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosUsers 更新hosUsers表
// @Tags HosUsers
// @Summary 更新hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosUsers true "更新hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosUsers/updateHosUsers [put]
func (hosUsersApi *HosUsersApi) UpdateHosUsers(c *gin.Context) {
	var hosUsers hos.HosUsers
	err := c.ShouldBindJSON(&hosUsers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := hosUsersService.UpdateHosUsers(hosUsers, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosUsers 用id查询hosUsers表
// @Tags HosUsers
// @Summary 用id查询hosUsers表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosUsers true "用id查询hosUsers表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosUsers/findHosUsers [get]
func (hosUsersApi *HosUsersApi) FindHosUsers(c *gin.Context) {
	ID := c.Query("ID")
	if rehosUsers, err := hosUsersService.GetHosUsers(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosUsers": rehosUsers}, c)
	}
}

// GetHosUsersList 分页获取hosUsers表列表
// @Tags HosUsers
// @Summary 分页获取hosUsers表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosUsersSearch true "分页获取hosUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUsers/getHosUsersList [get]
func (hosUsersApi *HosUsersApi) GetHosUsersList(c *gin.Context) {
	var pageInfo hosReq.HosUsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosUsersService.GetHosUsersInfoList(pageInfo, c); err != nil {
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

// GetHosUsersDataSource 获取HosUsers的数据源
// @Tags HosUsers
// @Summary 获取HosUsers的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUsers/getHosUsersDataSource [get]
func (hosUsersApi *HosUsersApi) GetHosUsersDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	if dataSource, err := hosUsersService.GetHosUsersDataSource(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(dataSource, c)
	}
}

// GetHosUsersPublic 不需要鉴权的hosUsers表接口
// @Tags HosUsers
// @Summary 不需要鉴权的hosUsers表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosUsersSearch true "分页获取hosUsers表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosUsers/getHosUsersPublic [get]
func (hosUsersApi *HosUsersApi) GetHosUsersPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的hosUsers表接口信息",
	}, "获取成功", c)
}
