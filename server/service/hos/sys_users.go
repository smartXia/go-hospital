package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

type SysUsersService struct {
}

// CreateSysUsers 创建sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) CreateSysUsers(sysUsers *hos.SysUsers, ctx *gin.Context) (err error, d *hos.SysUsers) {
	if sysUsers.Phone != "" {
		su := hos.SysUsers{}
		global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("phone = ?", sysUsers.Phone).Where("deleted_at != ?", nil).First(&su)
		if su.ID != 0 {
			return errors.New("手机号已经存在"), d
		}
	}
	if sysUsers.Username != "" && sysUsers.NickName == "" {
		sysUsers.NickName = sysUsers.Username
	}
	sysUsers.Uuid = utils.UniqueId()
	sysUsers.CreatedBy = utils.GetUserID(ctx)
	sysUsers.Password = utils.BcryptHash(sysUsers.Password)
	err = global.GVA_DB.Scopes().Create(sysUsers).Error
	if sysUsers.AuthorityId != 0 {
		a := hos.SysUserAuthority{
			SysUserId:               sysUsers.ID,
			SysAuthorityAuthorityId: sysUsers.AuthorityId,
		}
		db := global.GVA_DB.Model(&hos.SysUserAuthority{})
		_ = db.Create(&a).Error
	}
	return err, sysUsers
}

// DeleteSysUsers 删除sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) DeleteSysUsers(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Delete(&hos.SysUsers{}, "id = ?", ID).Error
	return err
}

// DeleteSysUsersByIds 批量删除sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) DeleteSysUsersByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.SysUsers{}, "id in ?", IDs).Error
	return err
}

// UpdateSysUsers 更新sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) UpdateSysUsers(sysUsers hos.SysUsers, ctx *gin.Context) (err error) {
	sysUsers.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.SysUsers{}).Where("id = ?", sysUsers.ID).Updates(&sysUsers).Error
	return err
}

// GetSysUsers 根据ID获取sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) GetSysUsers(ID string, ctx *gin.Context) (sysUsers hos.SysUsers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&sysUsers).Error
	return
}

// GetSysUsersInfoList 分页获取sysUsers表记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysUsersService *SysUsersService) GetSysUsersInfoList(info hosReq.SysUsersSearch, ctx *gin.Context) (list []hos.SysUsers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.SysUsers{})
	if info.Hospital != 0 {
		db = db.Where("hospital = ?", info.Hospital)
	} else {
		tid := utils.GetTenantId(ctx)
		if tid != 0 {
			db = db.Where("tenant_id = ?", tid)
		}
	}
	var sysUserss []hos.SysUsers
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	if info.Dept != "" {
		db = db.Where("dept = ?", info.Dept)
	}
	if info.NickName != "" {
		db = db.Where("username = ?", info.NickName)
	}
	if info.Post != "" {
		db = db.Where("post = ?", info.Post)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}
	db.Preload("DeptInfo")
	db.Preload("PostInfo")
	err = db.Find(&sysUserss).Error
	return sysUserss, total, err
}
func (sysUsersService *SysUsersService) GetSysUsersDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	dept := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_dept").Select("name as label,id as value").Scan(&dept)
	res["dept"] = dept
	hospital := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_org").Select("name as label,id as value").Scan(&hospital)
	res["hospital"] = hospital
	post := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_post").Select("name as label,id as value").Scan(&post)
	res["post"] = post
	return
}
