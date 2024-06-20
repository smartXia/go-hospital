package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type SysOrgService struct {
}

// CreateSysOrg 创建sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService) CreateSysOrg(sysOrg *hos.SysOrg, ctx *gin.Context) (err error, d *hos.SysOrg) {
	sysOrg.CreatedBy = utils.GetUserID(ctx)

	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(sysOrg).Error
	return err, sysOrg
}

// DeleteSysOrg 删除sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService) DeleteSysOrg(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.SysOrg{}, "id = ?", ID).Error
	return err
}

// DeleteSysOrgByIds 批量删除sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService) DeleteSysOrgByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.SysOrg{}, "id in ?", IDs).Error
	return err
}

// UpdateSysOrg 更新sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService) UpdateSysOrg(sysOrg hos.SysOrg, ctx *gin.Context) (err error) {
	sysOrg.UpdatedBy = utils.GetUserID(ctx)

	err = global.GVA_DB.Model(&hos.SysOrg{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", sysOrg.ID).Updates(&sysOrg).Error
	return err
}

// GetSysOrg 根据ID获取sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService) GetSysOrg(ID string, ctx *gin.Context) (sysOrg hos.SysOrg, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&sysOrg).Error
	return
}

// GetSysOrgInfoList 分页获取sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService) GetSysOrgInfoList(info hosReq.SysOrgSearch, ctx *gin.Context) (list []hos.SysOrg, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysOrg{}).Scopes(scope.TenantScope(ctx))
	var sysOrgs []hos.SysOrg
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Desc != "" {
		db = db.Where("desc LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&sysOrgs).Error
	return sysOrgs, total, err
}
