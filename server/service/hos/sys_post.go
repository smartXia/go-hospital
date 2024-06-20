package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type SysPostService struct {
}

// CreateSysPost 创建sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) CreateSysPost(sysPost *hos.SysPost, ctx *gin.Context) (err error, d *hos.SysPost) {
	sysPost.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(sysPost).Error
	return err, sysPost
}

// DeleteSysPost 删除sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) DeleteSysPost(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.SysPost{}, "id = ?", ID).Error
	return err
}

// DeleteSysPostByIds 批量删除sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) DeleteSysPostByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.SysPost{}, "id in ?", IDs).Error
	return err
}

// UpdateSysPost 更新sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) UpdateSysPost(sysPost hos.SysPost, ctx *gin.Context) (err error) {
	sysPost.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.SysPost{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", sysPost.ID).Updates(&sysPost).Error
	return err
}

// GetSysPost 根据ID获取sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) GetSysPost(ID string, ctx *gin.Context) (sysPost hos.SysPost, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&sysPost).Error
	return
}

// GetSysPostInfoList 分页获取sysPost表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysPostService *SysPostService) GetSysPostInfoList(info hosReq.SysPostSearch, ctx *gin.Context) (list []hos.SysPost, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysPost{}).Scopes(scope.TenantScope(ctx))
	var sysPosts []hos.SysPost
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

	err = db.Find(&sysPosts).Error
	return sysPosts, total, err
}
