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

type SysDeptApi struct {
}

var sysDeptService = service.ServiceGroupApp.HosServiceGroup.SysDeptService

// CreateSysDept 创建sysDept表
// @Tags SysDept
// @Summary 创建sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysDept true "创建sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDept/createSysDept [post]
func (sysDeptApi *SysDeptApi) CreateSysDept(c *gin.Context) {
	var sysDept hos.SysDept
	err := c.ShouldBindJSON(&sysDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, d := sysDeptService.CreateSysDept(&sysDept, c); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteSysDept 删除sysDept表
// @Tags SysDept
// @Summary 删除sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysDept true "删除sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDept/deleteSysDept [delete]
func (sysDeptApi *SysDeptApi) DeleteSysDept(c *gin.Context) {
	ID := c.Query("ID")
	if err := sysDeptService.DeleteSysDept(ID, c); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysDeptByIds 批量删除sysDept表
// @Tags SysDept
// @Summary 批量删除sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysDept/deleteSysDeptByIds [delete]
func (sysDeptApi *SysDeptApi) DeleteSysDeptByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := sysDeptService.DeleteSysDeptByIds(IDs, c); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysDept 更新sysDept表
// @Tags SysDept
// @Summary 更新sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.SysDept true "更新sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDept/updateSysDept [put]
func (sysDeptApi *SysDeptApi) UpdateSysDept(c *gin.Context) {
	var sysDept hos.SysDept
	err := c.ShouldBindJSON(&sysDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := sysDeptService.UpdateSysDept(sysDept, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysDept 用id查询sysDept表
// @Tags SysDept
// @Summary 用id查询sysDept表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.SysDept true "用id查询sysDept表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDept/findSysDept [get]
func (sysDeptApi *SysDeptApi) FindSysDept(c *gin.Context) {
	ID := c.Query("ID")
	if resysDept, err := sysDeptService.GetSysDept(ID, c); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysDept": resysDept}, c)
	}
}

// GetSysDeptList 分页获取sysDept表列表
// @Tags SysDept
// @Summary 分页获取sysDept表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysDeptSearch true "分页获取sysDept表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDept/getSysDeptList [get]
func (sysDeptApi *SysDeptApi) GetSysDeptList(c *gin.Context) {
	var pageInfo hosReq.SysDeptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysDeptService.GetSysDeptInfoList(pageInfo, c); err != nil {
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

// GetSysDeptPublic 不需要鉴权的sysDept表接口
// @Tags SysDept
// @Summary 不需要鉴权的sysDept表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysDeptSearch true "分页获取sysDept表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDept/getSysDeptPublic [get]
func (sysDeptApi *SysDeptApi) GetSysDeptPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的sysDept表接口信息",
	}, "获取成功", c)
}

// GetSysDeptList 分页获取sysDept表列表
// @Tags Tree
// @Summary 分页获取sysDept表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.SysDeptSearch true "分页获取sysDept表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDept/getSysDeptList [get]

func (sysDeptApi *SysDeptApi) Tree(c *gin.Context) {
	var pageInfo hosReq.SysDeptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, err := sysDeptService.Tree(pageInfo, c); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
