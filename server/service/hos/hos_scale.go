package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type HosScaleService struct {
}

// CreateHosScale 创建hosScale表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosScaleService *HosScaleService) CreateHosScale(hosScale *hos.HosScale, ctx *gin.Context) (err error, h *hos.HosScale) {
	hosScale.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantSaveScope(ctx)).Create(hosScale).Error
	return err, hosScale
}

// DeleteHosScale 删除hosScale表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosScaleService *HosScaleService) DeleteHosScale(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosScale{}, "id = ?", ID).Error
	return err
}

// DeleteHosScaleByIds 批量删除hosScale表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosScaleService *HosScaleService) DeleteHosScaleByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosScale{}, "id in ?", IDs).Error
	return err
}

// UpdateHosScale 更新hosScale表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosScaleService *HosScaleService) UpdateHosScale(hosScale hos.HosScale, ctx *gin.Context) (err error) {
	hosScale.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosScale{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosScale.ID).Updates(&hosScale).Error
	return err
}

// GetHosScale 根据ID获取hosScale表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosScaleService *HosScaleService) GetHosScale(ID string, ctx *gin.Context) (hosScale hos.HosScale, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosScale).Error
	return
}

// GetHosScaleInfoList 分页获取hosScale表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosScaleService *HosScaleService) GetHosScaleInfoList(info hosReq.HosScaleSearch, ctx *gin.Context) (list []hos.HosScale, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosScale{}).Scopes(scope.TenantScope(ctx))
	var hosScales []hos.HosScale
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("id desc").Find(&hosScales).Error
	return hosScales, total, err
}

// GetCurrentUserHosScaleInfoList 分页获取hosScale表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosScaleService *HosScaleService) GetCurrentUserHosScaleInfoList(info hosReq.HosScaleSearch, ctx *gin.Context) (list []hos.HosScale, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosScale{}).Scopes(scope.TenantScope(ctx))
	var hosScales []hos.HosScale
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	uid := utils.GetUserID(ctx)

	db = db.Where("hos_user_id = ?", uid)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosScales).Error
	return hosScales, total, err
}
