package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"github.com/gin-gonic/gin"
)

type HosUsersService struct {
}

// CreateHosUsers 创建hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) CreateHosUsers(hosUsers *hos.HosUsers, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosUsers).Error
	return err
}

// DeleteHosUsers 删除hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) DeleteHosUsers(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosUsers{}, "id = ?", ID).Error
	return err
}

// DeleteHosUsersByIds 批量删除hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) DeleteHosUsersByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosUsers{}, "id in ?", IDs).Error
	return err
}

// UpdateHosUsers 更新hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) UpdateHosUsers(hosUsers hos.HosUsers, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Model(&hos.HosUsers{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosUsers.ID).Updates(&hosUsers).Error
	return err
}

// GetHosUsers 根据ID获取hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) GetHosUsers(ID string, ctx *gin.Context) (hosUsers hos.HosUsers, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosUsers).Error
	return
}

// GetHosUsersInfoList 分页获取hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) GetHosUsersInfoList(info hosReq.HosUsersSearch, ctx *gin.Context) (list []hos.HosUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosUsers{}).Scopes(scope.TenantScope(ctx))
	var hosUserss []hos.HosUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Uuid != "" {
		db = db.Where("uuid = ?", info.Uuid)
	}
	if info.Username != "" {
		db = db.Where("username = ?", info.Username)
	}
	if info.NickName != "" {
		db = db.Where("nick_name = ?", info.NickName)
	}
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	if info.JianhuPhone != "" {
		db = db.Where("jianhu_phone = ?", info.JianhuPhone)
	}
	if info.Jianhuren != "" {
		db = db.Where("jianhuren = ?", info.Jianhuren)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&hosUserss).Error
	return hosUserss, total, err
}
func (hosUsersService *HosUsersService) GetHosUsersDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	latelyHos := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_org").Select("name as label,id as value").Where("deleted_at IS  NULL").Find(&latelyHos)
	res["latelyHos"] = latelyHos
	registerHos := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_org").Select("name as label,id as value").Where("deleted_at IS  NULL").Scan(&registerHos)
	res["registerHos"] = registerHos
	return
}
