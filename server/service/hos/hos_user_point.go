package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type HosUserPointService struct {
}

// CreateHosUserPoint 创建hosUserPoint表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUserPointService *HosUserPointService) CreateHosUserPoint(hosUserPoint *hos.HosUserPoint, ctx *gin.Context) (err error, d *hos.HosUserPoint) {
	hosUserPoint.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosUserPoint).Error
	return err, hosUserPoint
}

// DeleteHosUserPoint 删除hosUserPoint表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUserPointService *HosUserPointService) DeleteHosUserPoint(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosUserPoint{}, "id = ?", ID).Error
	return err
}

// DeleteHosUserPointByIds 批量删除hosUserPoint表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUserPointService *HosUserPointService) DeleteHosUserPointByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosUserPoint{}, "id in ?", IDs).Error
	return err
}

// UpdateHosUserPoint 更新hosUserPoint表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUserPointService *HosUserPointService) UpdateHosUserPoint(hosUserPoint hos.HosUserPoint, ctx *gin.Context) (err error) {
	hosUserPoint.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosUserPoint{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosUserPoint.ID).Updates(&hosUserPoint).Error
	return err
}

// GetHosUserPoint 根据ID获取hosUserPoint表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUserPointService *HosUserPointService) GetHosUserPoint(ID string, ctx *gin.Context) (hosUserPoint hos.HosUserPoint, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosUserPoint).Error
	return
}

// GetHosUserPointInfoList 分页获取hosUserPoint表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUserPointService *HosUserPointService) GetHosUserPointInfoList(info hosReq.HosUserPointSearch, ctx *gin.Context) (list []hos.HosUserPoint, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosUserPoint{}).Scopes(scope.TenantScope(ctx))
	var hosUserPoints []hos.HosUserPoint
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
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosUserPoints).Error
	return hosUserPoints, total, err
}
