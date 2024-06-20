package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"github.com/gin-gonic/gin"
)

type SysOperationRecordsService struct {
}

// CreateSysOperationRecords 创建sysOperationRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOperationRecordsService *SysOperationRecordsService) CreateSysOperationRecords(sysOperationRecords *hos.SysOperationRecords, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(sysOperationRecords).Error
	return err
}

// DeleteSysOperationRecords 删除sysOperationRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOperationRecordsService *SysOperationRecordsService) DeleteSysOperationRecords(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.SysOperationRecords{}, "id = ?", ID).Error
	return err
}

// DeleteSysOperationRecordsByIds 批量删除sysOperationRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOperationRecordsService *SysOperationRecordsService) DeleteSysOperationRecordsByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.SysOperationRecords{}, "id in ?", IDs).Error
	return err
}

// UpdateSysOperationRecords 更新sysOperationRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOperationRecordsService *SysOperationRecordsService) UpdateSysOperationRecords(sysOperationRecords hos.SysOperationRecords, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Model(&hos.SysOperationRecords{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", sysOperationRecords.ID).Updates(&sysOperationRecords).Error
	return err
}

// GetSysOperationRecords 根据ID获取sysOperationRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOperationRecordsService *SysOperationRecordsService) GetSysOperationRecords(ID string, ctx *gin.Context) (sysOperationRecords hos.SysOperationRecords, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&sysOperationRecords).Error
	return
}

// GetSysOperationRecordsInfoList 分页获取sysOperationRecords表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOperationRecordsService *SysOperationRecordsService) GetSysOperationRecordsInfoList(info hosReq.SysOperationRecordsSearch, ctx *gin.Context) (list []hos.SysOperationRecords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysOperationRecords{}).Scopes(scope.TenantScope(ctx))
	var sysOperationRecordss []hos.SysOperationRecords
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Action != "" {
		db = db.Where("action = ?", info.Action)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&sysOperationRecordss).Error
	return sysOperationRecordss, total, err
}
