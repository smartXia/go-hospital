package hos

import (
	"devops-manage/global"
    "devops-manage/model/hos"
    hosReq "devops-manage/model/hos/request"
    "devops-manage/model/common/response"
    "devops-manage/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type HosSportClockCommitApi struct {
}

var hosSportClockCommitService = service.ServiceGroupApp.HosServiceGroup.HosSportClockCommitService


// CreateHosSportClockCommit 创建hosSportClockCommit表
// @Tags HosSportClockCommit
// @Summary 创建hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportClockCommit true "创建hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /hosSportClockCommit/createHosSportClockCommit [post]
func (hosSportClockCommitApi *HosSportClockCommitApi) CreateHosSportClockCommit(c *gin.Context) {
	var hosSportClockCommit hos.HosSportClockCommit
	err := c.ShouldBindJSON(&hosSportClockCommit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err,d := hosSportClockCommitService.CreateHosSportClockCommit(&hosSportClockCommit,c); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithData(d, c)
	}
}

// DeleteHosSportClockCommit 删除hosSportClockCommit表
// @Tags HosSportClockCommit
// @Summary 删除hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportClockCommit true "删除hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /hosSportClockCommit/deleteHosSportClockCommit [delete]
func (hosSportClockCommitApi *HosSportClockCommitApi) DeleteHosSportClockCommit(c *gin.Context) {
	ID := c.Query("ID")
	if err := hosSportClockCommitService.DeleteHosSportClockCommit(ID,c); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteHosSportClockCommitByIds 批量删除hosSportClockCommit表
// @Tags HosSportClockCommit
// @Summary 批量删除hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /hosSportClockCommit/deleteHosSportClockCommitByIds [delete]
func (hosSportClockCommitApi *HosSportClockCommitApi) DeleteHosSportClockCommitByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := hosSportClockCommitService.DeleteHosSportClockCommitByIds(IDs,c); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateHosSportClockCommit 更新hosSportClockCommit表
// @Tags HosSportClockCommit
// @Summary 更新hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body hos.HosSportClockCommit true "更新hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /hosSportClockCommit/updateHosSportClockCommit [put]
func (hosSportClockCommitApi *HosSportClockCommitApi) UpdateHosSportClockCommit(c *gin.Context) {
	var hosSportClockCommit hos.HosSportClockCommit
	err := c.ShouldBindJSON(&hosSportClockCommit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := hosSportClockCommitService.UpdateHosSportClockCommit(hosSportClockCommit,c); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindHosSportClockCommit 用id查询hosSportClockCommit表
// @Tags HosSportClockCommit
// @Summary 用id查询hosSportClockCommit表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hos.HosSportClockCommit true "用id查询hosSportClockCommit表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /hosSportClockCommit/findHosSportClockCommit [get]
func (hosSportClockCommitApi *HosSportClockCommitApi) FindHosSportClockCommit(c *gin.Context) {
	ID := c.Query("ID")
	if rehosSportClockCommit, err := hosSportClockCommitService.GetHosSportClockCommit(ID,c); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rehosSportClockCommit": rehosSportClockCommit}, c)
	}
}

// GetHosSportClockCommitList 分页获取hosSportClockCommit表列表
// @Tags HosSportClockCommit
// @Summary 分页获取hosSportClockCommit表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportClockCommitSearch true "分页获取hosSportClockCommit表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportClockCommit/getHosSportClockCommitList [get]
func (hosSportClockCommitApi *HosSportClockCommitApi) GetHosSportClockCommitList(c *gin.Context) {
	var pageInfo hosReq.HosSportClockCommitSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := hosSportClockCommitService.GetHosSportClockCommitInfoList(pageInfo,c); err != nil {
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

// GetHosSportClockCommitPublic 不需要鉴权的hosSportClockCommit表接口
// @Tags HosSportClockCommit
// @Summary 不需要鉴权的hosSportClockCommit表接口
// @accept application/json
// @Produce application/json
// @Param data query hosReq.HosSportClockCommitSearch true "分页获取hosSportClockCommit表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /hosSportClockCommit/getHosSportClockCommitPublic [get]
func (hosSportClockCommitApi *HosSportClockCommitApi) GetHosSportClockCommitPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的hosSportClockCommit表接口信息",
    }, "获取成功", c)
}
