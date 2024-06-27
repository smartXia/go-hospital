package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type HosDashboardService struct {
}

// CreateHosDashboard 创建hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) CreateHosDashboard(hosDashboard *hos.HosDashboard, ctx *gin.Context, init int) (err error, d *hos.HosDashboard) {
	hosDashboard.CreatedBy = utils.GetUserID(ctx)
	if init == 1 {
		hosDashboard.BuildInitDashBoardDate()
	}
	err = global.GVA_DB.Scopes(scope.TenantSaveScope(ctx)).Create(hosDashboard).Error
	return err, hosDashboard
}

// DeleteHosDashboard 删除hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) DeleteHosDashboard(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosDashboard{}, "id = ?", ID).Error
	return err
}

// DeleteHosDashboardByIds 批量删除hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) DeleteHosDashboardByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosDashboard{}, "id in ?", IDs).Error
	return err
}

// UpdateHosDashboard 更新hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) UpdateHosDashboard(hosDashboard hos.HosDashboard, ctx *gin.Context) (err error) {
	hosDashboard.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosDashboard{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosDashboard.ID).Updates(&hosDashboard).Error
	return err
}

// GetHosDashboard 根据ID获取hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) GetHosDashboard(ID string, ctx *gin.Context) (hosDashboard hos.HosDashboard, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosDashboard).Error
	return
}

// GetHosDashboardInfoList 分页获取hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) GetHosDashboardInfoList(info hosReq.HosDashboardSearch, ctx *gin.Context) (list []hos.HosDashboard, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosDashboard{}).Scopes(scope.TenantScope(ctx))
	var hosDashboards []hos.HosDashboard
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrgId != nil {
		db = db.Where("org_id = ?", info.OrgId)
	}
	if info.Enable != nil {
		db = db.Where("enable = ?", info.Enable)
	}
	db = db.Preload("SysOrg")
	db = db.Preload("SysUsersInfo")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosDashboards).Error
	return hosDashboards, total, err
}

func (hosDashboardService *HosDashboardService) GetCurrentDashBoardInfo(info hosReq.HosDashboardSearch, ctx *gin.Context) (hosDashboard hos.HosDashboard, err error) {
	// 创建db
	db := global.GVA_DB.Model(&hos.HosDashboard{}).Scopes(scope.TenantScope(ctx))
	// 如果有条件搜索 下方会自动创建搜索语句
	tid := utils.GetTenantId(ctx)
	if info.OrgId != nil {
		db = db.Where("org_id = ?", info.OrgId)
	} else {
		db = db.Where("org_id = ?", tid)
	}
	db = db.Preload("SysOrg")
	db = db.Preload("SysUsersInfo")
	err = db.First(&hosDashboard).Error
	if hosDashboard.ID == 0 {
		return hos.HosDashboard{}, nil
	}
	if hosDashboard.Enable == nil {
		return hosDashboard, err
	} else {
		//默认处理
		return hosDashboard, err
	}
}

//func BuildDa0ta() {
//	var a []hos.Duilienianling
//	var b []hos.Duilieyanzhongxing
//	var c []hos.Duiliefenlei
//	var d []hos.Diqupaihang
//}

func (hosDashboardService *HosDashboardService) GetHosDashboardDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	orgId := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_org").Select("name as label,id as value").Scan(&orgId)
	res["orgId"] = orgId
	return
}
