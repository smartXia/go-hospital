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

type SysPostApi struct {
}

var sysPostService = service.ServiceGroupApp.HosServiceGroup.SysPostService

// CreateSysPost 创建sysPost表
// @Tags SysPost
// @Summary 创建sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysPost true "创建sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysPost/createSysPost [post]
func (sysPostApi *SysPostApi) CreateSysPost(c *gin.Context) {
	var sysPost hos.SysPost
	err := c.ShouldBindJSON(&sysPost)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysPostService.CreateSysPost(&sysPost, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysPost 删除sysPost表
// @Tags SysPost
// @Summary 删除sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysPost true "删除sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysPost/deleteSysPost [delete]
func (sysPostApi *SysPostApi) DeleteSysPost(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysPostService.DeleteSysPost(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysPostByIds 批量删除sysPost表
// @Tags SysPost
// @Summary 批量删除sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysPost/deleteSysPostByIds [delete]
func (sysPostApi *SysPostApi) DeleteSysPostByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysPostService.DeleteSysPostByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysPost 更新sysPost表
// @Tags SysPost
// @Summary 更新sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysPost true "更新sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysPost/updateSysPost [put]
func (sysPostApi *SysPostApi) UpdateSysPost(c *gin.Context) {
	var sysPost hos.SysPost
	err := c.ShouldBindJSON(&sysPost)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysPostService.UpdateSysPost(sysPost, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysPost 用id查询sysPost表
// @Tags SysPost
// @Summary 用id查询sysPost表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.SysPost true "用id查询sysPost表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysPost/findSysPost [get]
func (sysPostApi *SysPostApi) FindSysPost(c *gin.Context) {
	ID := c.Query("ID")
	if resysPost, err := sysPostService.GetSysPost(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysPost": resysPost}, c)
	}
}

// GetSysPostList 分页获取sysPost表列表
// @Tags SysPost
// @Summary 分页获取sysPost表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysPostSearch true "分页获取sysPost表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysPost/getSysPostList [get]
func (sysPostApi *SysPostApi) GetSysPostList(c *gin.Context) {
	var pageInfo hosReq.SysPostSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysPostService.GetSysPostInfoList(pageInfo, c); err != nil {
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

// GetSysPostPublic 不需要鉴权的sysPost表接口
// @Tags SysPost
// @Summary 不需要鉴权的sysPost表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysPostSearch true "分页获取sysPost表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysPost/getSysPostPublic [get]
func (sysPostApi *SysPostApi) GetSysPostPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的sysPost表接口信息",
	}, "获取成功", c)
}
