package hos

import (
	"devops-manage/global"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
)

type SysDeptService struct {
}

// CreateSysDept 创建sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) CreateSysDept(sysDept *hos.SysDept) (err error) {
	err = global.GVA_DB.Create(sysDept).Error
	return err
}

// DeleteSysDept 删除sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) DeleteSysDept(ID string) (err error) {
	err = global.GVA_DB.Delete(&hos.SysDept{}, "id = ?", ID).Error
	return err
}

// DeleteSysDeptByIds 批量删除sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) DeleteSysDeptByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]hos.SysDept{}, "id in ?", IDs).Error
	return err
}

// UpdateSysDept 更新sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) UpdateSysDept(sysDept hos.SysDept) (err error) {
	err = global.GVA_DB.Model(&hos.SysDept{}).Where("id = ?", sysDept.ID).Updates(&sysDept).Error
	return err
}

// GetSysDept 根据ID获取sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) GetSysDept(ID string) (sysDept hos.SysDept, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysDept).Error
	return
}

// GetSysDeptInfoList 分页获取sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) GetSysDeptInfoList(info hosReq.SysDeptSearch) (list []hos.SysDept, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysDept{})
	var sysDepts []hos.SysDept
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

	err = db.Find(&sysDepts).Error
	return sysDepts, total, err
}
