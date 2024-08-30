package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"sort"
)

type SysOrgService struct {
}

var hosDashboardService HosDashboardService

// CreateSysOrg 创建sysOrg表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysOrgService *SysOrgService) CreateSysOrg(sysOrg *hos.SysOrg, ctx *gin.Context) (err error, d *hos.SysOrg) {
	sysOrg.CreatedBy = utils.GetUserID(ctx)

	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(sysOrg).Error

	var hosDashboard hos.HosDashboard
	hosDashboard.OrgId = sysOrg.ID
	a := int(sysOrg.ID)
	hosDashboard.TenantId = &a
	hosDashboard.Name = fmt.Sprintf("%s-医院大屏", sysOrg.Name)
	if err, _ := hosDashboardService.CreateHosDashboard(&hosDashboard, ctx, 1); err != nil {

	}
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
	db := global.GVA_DB.Model(&hos.SysOrg{})
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
	var list2 []hos.SysOrg
	if len(sysOrgs) != 0 {
		var g errgroup.Group
		for _, org := range sysOrgs {
			db1 := global.GVA_DB.Model(&hos.HosFlow{}).Scopes(scope.TenantScope(ctx))
			db2 := global.GVA_DB.Model(&hos.HosUsers{}).Scopes(scope.TenantScope(ctx))
			db3 := global.GVA_DB.Model(&hos.SysUsers{}).Where("hospital = ?", org.ID)
			g.Go(func() error {
				var uhC int64
				var hC int64
				var fC int64
				err = db1.Count(&fC).Error
				err = db2.Count(&hC).Error
				err = db3.Count(&uhC).Error
				if err != nil {
					return err
				}
				org.HosUserNum = hC
				org.SysUserNum = uhC
				org.FlowNum = fC
				list2 = append(list2, org)
				return err
			})
		}
		if err = g.Wait(); err != nil {
			return
		}
		// Sort the children by ID
		sort.Slice(list2, func(i, j int) bool {
			return list2[i].ID < list2[j].ID
		})
		return list2, total, err
	}

	return sysOrgs, total, err
}
