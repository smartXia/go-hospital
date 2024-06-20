package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type HosLoaclAskService struct {
}

// CreateHosLoaclAsk 创建hosLoaclAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLoaclAskService *HosLoaclAskService) CreateHosLoaclAsk(hosLoaclAsk *hos.HosLoaclAsk, ctx *gin.Context) (err error, h *hos.HosLoaclAsk) {
	hosLoaclAsk.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosLoaclAsk).Error
	return err, hosLoaclAsk
}

// DeleteHosLoaclAsk 删除hosLoaclAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLoaclAskService *HosLoaclAskService) DeleteHosLoaclAsk(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosLoaclAsk{}, "id = ?", ID).Error
	return err
}

// DeleteHosLoaclAskByIds 批量删除hosLoaclAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLoaclAskService *HosLoaclAskService) DeleteHosLoaclAskByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosLoaclAsk{}, "id in ?", IDs).Error
	return err
}

// UpdateHosLoaclAsk 更新hosLoaclAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLoaclAskService *HosLoaclAskService) UpdateHosLoaclAsk(hosLoaclAsk hos.HosLoaclAsk, ctx *gin.Context) (err error) {
	hosLoaclAsk.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosLoaclAsk{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosLoaclAsk.ID).Updates(&hosLoaclAsk).Error
	return err
}

// GetHosLoaclAsk 根据ID获取hosLoaclAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLoaclAskService *HosLoaclAskService) GetHosLoaclAsk(ID string, ctx *gin.Context) (hosLoaclAsk hos.HosLoaclAsk, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosLoaclAsk).Error
	return
}

// GetHosLoaclAskInfoList 分页获取hosLoaclAsk表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosLoaclAskService *HosLoaclAskService) GetHosLoaclAskInfoList(info hosReq.HosLoaclAskSearch, ctx *gin.Context) (list []hos.HosLoaclAsk, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosLoaclAsk{}).Scopes(scope.TenantScope(ctx))
	var hosLoaclAsks []hos.HosLoaclAsk
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
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&hosLoaclAsks).Error
	return hosLoaclAsks, total, err
}
