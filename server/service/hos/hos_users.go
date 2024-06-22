package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HosUsersService struct {
}

// CreateHosUsers 创建hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) CreateHosUsers(hosUsers *hos.HosUsers, ctx *gin.Context) (err error, d *hos.HosUsers) {
	hosUsers.CreatedBy = utils.GetUserID(ctx)
	if hosUsers.Phone != "" && hosUsers.JianhuPhone == "" {
		hosUsers.JianhuPhone = hosUsers.Phone
	}
	if hosUsers.Username != "" && hosUsers.NickName == "" {
		hosUsers.NickName = hosUsers.Username
	}
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosUsers).Error
	return err, hosUsers
}

// DeleteHosUsers 删除hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) DeleteHosUsers(ID string, userID uint, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&hos.HosUsers{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&hos.HosUsers{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteHosUsersByIds 批量删除hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) DeleteHosUsersByIds(IDs []string, deleted_by uint, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&hos.HosUsers{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&hos.HosUsers{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateHosUsers 更新hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) UpdateHosUsers(hosUsers hos.HosUsers, ctx *gin.Context) (err error) {
	hosUsers.UpdatedBy = utils.GetUserID(ctx)
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
	if info.CardNo != "" {
		db = db.Where("card_no LIKE ?", "%"+info.CardNo+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosUserss).Error
	return hosUserss, total, err
}

// GetCurrentHosUsersList 分页获取hosUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosUsersService *HosUsersService) GetCurrentHosUsersList(info hosReq.HosUsersSearch, ctx *gin.Context) (list []hos.HosUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosUsers{}).Scopes(scope.TenantScope(ctx))
	var hosUserss []hos.HosUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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
	if info.CardNo != "" {
		db = db.Where("card_no LIKE ?", "%"+info.CardNo+"%")
	}
	//UserPhone := utils.GetUserPhone(ctx)
	//db = db.Where("jianhu_phone = ?", UserPhone)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosUserss).Error
	return hosUserss, total, err
}

func (hosUsersService *HosUsersService) GetHosUsersIds(info hosReq.HosUsersSearch, ctx *gin.Context) (list []int, err error) {
	// 创建db
	db := global.GVA_DB.Model(&hos.HosUsers{}).Scopes(scope.TenantScope(ctx))
	var hosUserss []hos.HosUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	// 通过搜索患者id对应的监护人创建人 获取hos_user_id
	if info.CardNo != "" {
		db = db.Where("card_no LIKE ?", "%"+info.CardNo+"%")
	}
	//global.GVA_DB.Where("hos_user_id in (?) ", MenuIds)
	err = db.Find(&hosUserss).Error
	return list, err
}

func (hosUsersService *HosUsersService) GetHosUsersDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	latelyHos := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_org").Select("name as label,id as value").Scan(&latelyHos)
	res["latelyHos"] = latelyHos
	registerHos := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_org").Select("name as label,id as value").Scan(&registerHos)
	res["registerHos"] = registerHos
	return
}
