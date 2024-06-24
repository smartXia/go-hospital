package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type HosUsersDianzanService struct {
}

// CreateHosUsersDianzan 创建hosUsersDianzan表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersDianzanService *HosUsersDianzanService) CreateHosUsersDianzan(hosUsersDianzan *hos.HosUsersDianzan, ctx *gin.Context) (err error, d *hos.HosUsersDianzan) {
	hosUsersDianzan.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantSaveScope(ctx)).Create(hosUsersDianzan).Error
	return err, hosUsersDianzan
}

// DeleteHosUsersDianzan 删除hosUsersDianzan表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersDianzanService *HosUsersDianzanService) DeleteHosUsersDianzan(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosUsersDianzan{}, "id = ?", ID).Error
	return err
}

// DeleteHosUsersDianzanByIds 批量删除hosUsersDianzan表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersDianzanService *HosUsersDianzanService) DeleteHosUsersDianzanByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosUsersDianzan{}, "id in ?", IDs).Error
	return err
}

// UpdateHosUsersDianzan 更新hosUsersDianzan表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersDianzanService *HosUsersDianzanService) UpdateHosUsersDianzan(hosUsersDianzan hos.HosUsersDianzan, ctx *gin.Context) (err error) {
	hosUsersDianzan.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosUsersDianzan{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosUsersDianzan.ID).Updates(&hosUsersDianzan).Error
	return err
}

// GetHosUsersDianzan 根据ID获取hosUsersDianzan表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersDianzanService *HosUsersDianzanService) GetHosUsersDianzan(ID string, ctx *gin.Context) (hosUsersDianzan hos.HosUsersDianzan, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosUsersDianzan).Error
	return
}

// GetHosUsersDianzanInfoList 分页获取hosUsersDianzan表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersDianzanService *HosUsersDianzanService) GetHosUsersDianzanInfoList(info hosReq.HosUsersDianzanSearch, ctx *gin.Context) (list []hos.HosUsersDianzan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosUsersDianzan{}).Scopes(scope.TenantScope(ctx))
	var hosUsersDianzans []hos.HosUsersDianzan
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.HosUserId != nil {
		db = db.Where("hos_user_id = ?", info.HosUserId)
	}
	if info.CommitId != "" {
		db = db.Where("commit_id = ?", info.CommitId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosUsersDianzans).Error
	return hosUsersDianzans, total, err
}
