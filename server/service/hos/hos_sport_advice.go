package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type HosSportAdviceService struct {
}

// CreateHosSportAdvice 创建hosSportAdvice表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportAdviceService *HosSportAdviceService) CreateHosSportAdvice(hosSportAdvice *hos.HosSportAdvice, ctx *gin.Context) (err error, d *hos.HosSportAdvice) {
	hosSportAdvice.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosSportAdvice).Error
	return err, hosSportAdvice
}

// DeleteHosSportAdvice 删除hosSportAdvice表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportAdviceService *HosSportAdviceService) DeleteHosSportAdvice(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosSportAdvice{}, "id = ?", ID).Error
	return err
}

// DeleteHosSportAdviceByIds 批量删除hosSportAdvice表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportAdviceService *HosSportAdviceService) DeleteHosSportAdviceByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosSportAdvice{}, "id in ?", IDs).Error
	return err
}

// UpdateHosSportAdvice 更新hosSportAdvice表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportAdviceService *HosSportAdviceService) UpdateHosSportAdvice(hosSportAdvice hos.HosSportAdvice, ctx *gin.Context) (err error) {
	hosSportAdvice.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosSportAdvice{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosSportAdvice.ID).Updates(&hosSportAdvice).Error
	return err
}

// GetHosSportAdvice 根据ID获取hosSportAdvice表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportAdviceService *HosSportAdviceService) GetHosSportAdvice(ID string, ctx *gin.Context) (hosSportAdvice hos.HosSportAdvice, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosSportAdvice).Error
	return
}

// GetHosSportAdviceInfoList 分页获取hosSportAdvice表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportAdviceService *HosSportAdviceService) GetHosSportAdviceInfoList(info hosReq.HosSportAdviceSearch, ctx *gin.Context) (list []hos.HosSportAdvice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosSportAdvice{}).Scopes(scope.TenantScope(ctx))
	var hosSportAdvices []hos.HosSportAdvice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FlowId != nil {
		db = db.Where("flow_id = ?", info.FlowId)
	}
	if info.HosUserId != "" {
		db = db.Where("hos_user_id = ?", info.HosUserId)
	}
	if info.SportModeId != nil {
		db = db.Where("sport_mode_id = ?", info.SportModeId)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.Fuzhenriqi != "" {
		db = db.Where("fuzhenriqi = ?", info.Fuzhenriqi)
	}
	if info.SyncWx != nil {
		db = db.Where("sync_wx = ?", info.SyncWx)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosSportAdvices).Error
	return hosSportAdvices, total, err
}

// GetHosSportAdviceInfoList 分页获取hosSportAdvice表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportAdviceService *HosSportAdviceService) GetCurrentHosSportAdviceList(info hosReq.HosSportAdviceSearch, ctx *gin.Context) (list []hos.HosSportAdvice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosSportAdvice{}).Scopes(scope.TenantScope(ctx))
	var hosSportAdvices []hos.HosSportAdvice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FlowId != nil {
		db = db.Where("flow_id = ?", info.FlowId)
	}
	if info.HosUserId != "" {
		db = db.Where("hos_user_id = ?", info.HosUserId)
	}
	if info.SportModeId != nil {
		db = db.Where("sport_mode_id = ?", info.SportModeId)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.Fuzhenriqi != "" {
		db = db.Where("fuzhenriqi = ?", info.Fuzhenriqi)
	}
	if info.SyncWx != nil {
		db = db.Where("sync_wx = ?", info.SyncWx)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	uid := utils.GetUserID(ctx)
	//获取 自己运动建议
	db = db.Where("hos_user_id = ?", uid)
	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosSportAdvices).Error
	return hosSportAdvices, total, err
}
