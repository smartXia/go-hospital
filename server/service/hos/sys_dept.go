package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type SysDeptService struct {
}

// CreateSysDept 创建sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) CreateSysDept(sysDept *hos.SysDept, ctx *gin.Context) (err error, d *hos.SysDept) {
	sysDept.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(sysDept).Error
	return err, sysDept
}

// DeleteSysDept 删除sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) DeleteSysDept(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.SysDept{}, "id = ?", ID).Error
	return err
}

// DeleteSysDeptByIds 批量删除sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) DeleteSysDeptByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.SysDept{}, "id in ?", IDs).Error
	return err
}

// UpdateSysDept 更新sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) UpdateSysDept(sysDept hos.SysDept, ctx *gin.Context) (err error) {
	sysDept.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.SysDept{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", sysDept.ID).Updates(&sysDept).Error
	return err
}

// GetSysDept 根据ID获取sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) GetSysDept(ID string, ctx *gin.Context) (sysDept hos.SysDept, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&sysDept).Error
	return
}

// GetSysDeptInfoList 分页获取sysDept表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) GetSysDeptInfoList(info hosReq.SysDeptSearch, ctx *gin.Context) (list []hos.SysDept, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysDept{}).Scopes(scope.TenantScope(ctx))
	var sysDepts []hos.SysDept
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.ManageId != "" {
		db = db.Where("manage_id = ?", info.ManageId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&sysDepts).Error
	return sysDepts, total, err
}

// GetSysDeptInfoList 分页获取sysDept表记录

func (sysDeptService *SysDeptService) Tree(info hosReq.SysDeptSearch, ctx *gin.Context) (list []*hos.SysDept, err error) {
	// 创建db
	db := global.GVA_DB.Model(&hos.SysDept{}).Scopes(scope.TenantScope(ctx))
	var sysDepts []*hos.SysDept
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Find(&sysDepts).Error
	sysDepts = hos.BuildDeptTree(sysDepts)
	//tree := utils.BuildTree(sysDepts)
	//println(tree)
	return sysDepts, err
}
