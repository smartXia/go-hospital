package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type HosLocalAskService struct {
}

// CreateHosLocalAsk 创建hosLocalAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLocalAskService *HosLocalAskService) CreateHosLocalAsk(hosLocalAsk *hos.HosLocalAsk, ctx *gin.Context) (err error, h *hos.HosLocalAsk) {
	hosLocalAsk.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantSaveScope(ctx)).Create(hosLocalAsk).Error
	return err, hosLocalAsk
}

// DeleteHosLocalAsk 删除hosLocalAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLocalAskService *HosLocalAskService) DeleteHosLocalAsk(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosLocalAsk{}, "id = ?", ID).Error
	return err
}

// DeleteHosLocalAskByIds 批量删除hosLocalAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLocalAskService *HosLocalAskService) DeleteHosLocalAskByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosLocalAsk{}, "id in ?", IDs).Error
	return err
}

// UpdateHosLocalAsk 更新hosLocalAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLocalAskService *HosLocalAskService) UpdateHosLocalAsk(hosLocalAsk hos.HosLocalAsk, ctx *gin.Context) (err error) {
	hosLocalAsk.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosLocalAsk{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosLocalAsk.ID).Updates(&hosLocalAsk).Error
	return err
}

// GetHosLocalAsk 根据ID获取hosLocalAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLocalAskService *HosLocalAskService) GetHosLocalAsk(ID string, ctx *gin.Context) (hosLocalAsk hos.HosLocalAsk, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosLocalAsk).Error
	return
}

// GetHosLocalAskInfoList 分页获取hosLocalAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLocalAskService *HosLocalAskService) GetHosLocalAskInfoList(info hosReq.HosLocalAskSearch, ctx *gin.Context) (list []hos.HosLocalAsk, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosLocalAsk{}).Scopes(scope.TenantScope(ctx))
	var hosLocalAsks []hos.HosLocalAsk
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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

	err = db.Find(&hosLocalAsks).Error
	return hosLocalAsks, total, err
}
