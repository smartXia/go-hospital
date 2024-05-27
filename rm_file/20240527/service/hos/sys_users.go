package hos

import (
	"devops-manage/global"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"gorm.io/gorm"
)

type SysUsersService struct {
}

// CreateSysUsers 创建sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) CreateSysUsers(sysUsers *hos.SysUsers) (err error) {
	err = global.GVA_DB.Create(sysUsers).Error
	return err
}

// DeleteSysUsers 删除sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) DeleteSysUsers(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&hos.SysUsers{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&hos.SysUsers{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteSysUsersByIds 批量删除sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) DeleteSysUsersByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&hos.SysUsers{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&hos.SysUsers{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateSysUsers 更新sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) UpdateSysUsers(sysUsers hos.SysUsers) (err error) {
	err = global.GVA_DB.Model(&hos.SysUsers{}).Where("id = ?", sysUsers.ID).Updates(&sysUsers).Error
	return err
}

// GetSysUsers 根据ID获取sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) GetSysUsers(ID string) (sysUsers hos.SysUsers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysUsers).Error
	return
}

// GetSysUsersInfoList 分页获取sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) GetSysUsersInfoList(info hosReq.SysUsersSearch) (list []hos.SysUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysUsers{})
	var sysUserss []hos.SysUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Uuid != "" {
		db = db.Where("uuid LIKE ?", "%"+info.Uuid+"%")
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
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
