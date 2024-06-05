package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"github.com/gin-gonic/gin"
)

type SysUsersService struct {
}

// CreateSysUsers 创建sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) CreateSysUsers(sysUsers *hos.SysUsers, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(sysUsers).Error
	return err
}

// DeleteSysUsers 删除sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) DeleteSysUsers(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.SysUsers{}, "id = ?", ID).Error
	return err
}

// DeleteSysUsersByIds 批量删除sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) DeleteSysUsersByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.SysUsers{}, "id in ?", IDs).Error
	return err
}

// UpdateSysUsers 更新sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) UpdateSysUsers(sysUsers hos.SysUsers, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Model(&hos.SysUsers{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", sysUsers.ID).Updates(&sysUsers).Error
	return err
}

// GetSysUsers 根据ID获取sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) GetSysUsers(ID string, ctx *gin.Context) (sysUsers hos.SysUsers, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&sysUsers).Error
	return
}

// GetSysUsersInfoList 分页获取sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) GetSysUsersInfoList(info hosReq.SysUsersSearch, ctx *gin.Context) (list []hos.SysUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysUsers{}).Scopes(scope.TenantScope(ctx))
	var sysUserss []hos.SysUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	db = db.Where("authority_id  = ?", 0)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysUserss).Error
	return sysUserss, total, err
}
