package hos

import (
	"devops-manage/global"
	"devops-manage/model/hos"
    hosReq "devops-manage/model/hos/request"
    "gorm.io/gorm"
)

type SysOrgService struct {
}

// CreateSysOrg 创建sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService) CreateSysOrg(sysOrg *hos.SysOrg) (err error) {
	err = global.GVA_DB.Create(sysOrg).Error
	return err
}

// DeleteSysOrg 删除sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService)DeleteSysOrg(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hos.SysOrg{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&hos.SysOrg{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteSysOrgByIds 批量删除sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService)DeleteSysOrgByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hos.SysOrg{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&hos.SysOrg{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateSysOrg 更新sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService)UpdateSysOrg(sysOrg hos.SysOrg) (err error) {
	err = global.GVA_DB.Model(&hos.SysOrg{}).Where("id = ?",sysOrg.ID).Updates(&sysOrg).Error
	return err
}

// GetSysOrg 根据ID获取sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService)GetSysOrg(ID string) (sysOrg hos.SysOrg, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysOrg).Error
	return
}

// GetSysOrgInfoList 分页获取sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService)GetSysOrgInfoList(info hosReq.SysOrgSearch) (list []hos.SysOrg, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hos.SysOrg{})
    var sysOrgs []hos.SysOrg
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&sysOrgs).Error
	return  sysOrgs, total, err
}