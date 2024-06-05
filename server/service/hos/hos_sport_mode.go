package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"github.com/gin-gonic/gin"
)

type HosSportModeService struct {
}

// CreateHosSportMode 创建hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) CreateHosSportMode(hosSportMode *hos.HosSportMode, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosSportMode).Error
	return err
}

// DeleteHosSportMode 删除hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) DeleteHosSportMode(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosSportMode{}, "id = ?", ID).Error
	return err
}

// DeleteHosSportModeByIds 批量删除hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) DeleteHosSportModeByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosSportMode{}, "id in ?", IDs).Error
	return err
}

// UpdateHosSportMode 更新hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) UpdateHosSportMode(hosSportMode hos.HosSportMode, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Model(&hos.HosSportMode{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosSportMode.ID).Updates(&hosSportMode).Error
	return err
}

// GetHosSportMode 根据ID获取hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) GetHosSportMode(ID string, ctx *gin.Context) (hosSportMode hos.HosSportMode, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosSportMode).Error
	return
}

// GetHosSportModeInfoList 分页获取hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) GetHosSportModeInfoList(info hosReq.HosSportModeSearch, ctx *gin.Context) (list []hos.HosSportMode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosSportMode{}).Scopes(scope.TenantScope(ctx))
	var hosSportModes []hos.HosSportMode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Category != "" {
		db = db.Where("category = ?", info.Category)
	}
	if info.Buwei != "" {
		db = db.Where("buwei = ?", info.Buwei)
	}
	if info.Tiduan != "" {
		db = db.Where("tiduan = ?", info.Tiduan)
	}
	if info.Fangxiang != "" {
		db = db.Where("fangxiang = ?", info.Fangxiang)
	}
	if info.Weizhi != "" {
		db = db.Where("weizhi = ?", info.Weizhi)
	}
	if info.Source != "" {
		db = db.Where("source = ?", info.Source)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&hosSportModes).Error
	return hosSportModes, total, err
}