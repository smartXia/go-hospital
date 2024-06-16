package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"github.com/gin-gonic/gin"
)

type HosSportClockService struct {
}

// CreateHosSportClock 创建hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) CreateHosSportClock(hosSportClock *hos.HosSportClock, ctx *gin.Context) (err error, h *hos.HosSportClock) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosSportClock).Error
	return err, hosSportClock
}

// DeleteHosSportClock 删除hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) DeleteHosSportClock(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosSportClock{}, "id = ?", ID).Error
	return err
}

// DeleteHosSportClockByIds 批量删除hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) DeleteHosSportClockByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosSportClock{}, "id in ?", IDs).Error
	return err
}

// UpdateHosSportClock 更新hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) UpdateHosSportClock(hosSportClock hos.HosSportClock, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Model(&hos.HosSportClock{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosSportClock.ID).Updates(&hosSportClock).Error
	return err
}

// GetHosSportClock 根据ID获取hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) GetHosSportClock(ID string, ctx *gin.Context) (hosSportClock hos.HosSportClock, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosSportClock).Error
	return
}

// GetHosSportClockInfoList 分页获取hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) GetHosSportClockInfoList(info hosReq.HosSportClockSearch, ctx *gin.Context) (list []hos.HosSportClock, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosSportClock{}).Scopes(scope.TenantScope(ctx))
	var hosSportClocks []hos.HosSportClock
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.HosUserId != "" {
		db = db.Where("hos_user_id = ?", info.HosUserId)
	}
	if info.FlowId != "" {
		db = db.Where("flow_id = ?", info.FlowId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&hosSportClocks).Error
	return hosSportClocks, total, err
}
