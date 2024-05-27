package hos

import (
	"devops-manage/global"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
)

type SysPostService struct {
}

// CreateSysPost 创建sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) CreateSysPost(sysPost *hos.SysPost) (err error) {
	err = global.GVA_DB.Create(sysPost).Error
	return err
}

// DeleteSysPost 删除sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) DeleteSysPost(ID string) (err error) {
	err = global.GVA_DB.Delete(&hos.SysPost{}, "id = ?", ID).Error
	return err
}

// DeleteSysPostByIds 批量删除sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) DeleteSysPostByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]hos.SysPost{}, "id in ?", IDs).Error
	return err
}

// UpdateSysPost 更新sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) UpdateSysPost(sysPost hos.SysPost) (err error) {
	err = global.GVA_DB.Model(&hos.SysPost{}).Where("id = ?", sysPost.ID).Updates(&sysPost).Error
	return err
}

// GetSysPost 根据ID获取sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) GetSysPost(ID string) (sysPost hos.SysPost, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysPost).Error
	return
}

// GetSysPostInfoList 分页获取sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) GetSysPostInfoList(info hosReq.SysPostSearch) (list []hos.SysPost, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysPost{})
	var sysPosts []hos.SysPost
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysPosts).Error
	return sysPosts, total, err
}
