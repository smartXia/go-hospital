package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type HosSportModeService struct {
}

// CreateHosSportMode 创建hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) CreateHosSportMode(hosSportMode *hos.HosSportMode, ctx *gin.Context) (err error, d *hos.HosSportMode) {
	hosSportMode.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantSaveScope(ctx)).Create(hosSportMode).Error
	return err, hosSportMode
}

// DeleteHosSportMode 删除hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) DeleteHosSportMode(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosSportMode{}, "id = ?", ID).Error
	return err
}

// DeleteHosSportModeByIds 批量删除hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) DeleteHosSportModeByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosSportMode{}, "id in ?", IDs).Error
	return err
}

// UpdateHosSportMode 更新hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) UpdateHosSportMode(hosSportMode hos.HosSportMode, ctx *gin.Context) (err error) {
	hosSportMode.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosSportMode{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosSportMode.ID).Updates(&hosSportMode).Error
	return err
}

// GetHosSportMode 根据ID获取hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) GetHosSportMode(ID string, ctx *gin.Context) (hosSportMode hos.HosSportMode, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosSportMode).Error
	return
}

// GetHosSportModeInfoList 分页获取hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) GetHosSportModeInfoList(info hosReq.HosSportModeSearch, ctx *gin.Context) (list []hos.HosSportMode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&hos.HosSportMode{})

	// 创建db 这个接口没有scope 不加where 条件
	var hosSportModes []hos.HosSportMode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.Enable != nil {
		db.Where("enable = ?", info.Enable)
	}
	if info.Hospital != 0 {
		db = db.Where("hospital = ?", info.Hospital)
	}
	if info.Category != "" {
		tmp := strings.Split(info.Category, ",")
		db = db.Where("category in ?", tmp)
	}
	if info.Buwei != "" {
		tmp := strings.Split(info.Buwei, ",")
		db = db.Where("buwei in ?", tmp)
	}
	if info.Tiduan != "" {
		tmp := strings.Split(info.Tiduan, ",")
		db = db.Where("tiduan in ?", tmp)
	}
	if info.Xingdong != "" {
		tmp := strings.Split(info.Xingdong, ",")
		db = db.Where("xingdong in ?", tmp)
	}
	if info.Fangxiang != "" {
		tmp := strings.Split(info.Fangxiang, ",")
		db = db.Where("fangxiang in ?", tmp)
	}
	if info.Weizhi != "" {
		tmp := strings.Split(info.Weizhi, ",")
		db = db.Where("weizhi = ?", tmp)
	}
	id := utils.GetUserAuthorityId(ctx)
	if id == 100 {
		//管理员 那么去掉tenantId查询 获取所有的
	} else {
		//医院和医生 的管理需要加 tenant = 1
		if info.All == 1 {
			db.Where("enable = ?", 1)
			db.Where("tenant_id in ?", []int{0, utils.GetTenantId(ctx)})
		} else {
			db.Scopes(scope.TenantScope(ctx))
		}
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}
	db = db.Preload("SysUsersInfo")
	err = db.Find(&hosSportModes).Error
	return hosSportModes, total, err
}

var Field = map[int]string{
	1: "category",
	2: "buwei",
	3: "tiduan",
	4: "xingdong",
	5: "fangxiang",
	6: "weizhi",
}

var FieldSort = []string{
	"category",
	"buwei",
	"tiduan",
	"xingdong",
	"fangxiang",
	"weizhi",
}

func (hosSportModeService *HosSportModeService) GetHosSportModeMatrix(info hosReq.HosSportModeSearch, ctx *gin.Context) (res map[string][]string, total int64, err error) {
	//limit := info.PageSize
	//offset := info.PageSize * (info.Page - 1)
	// 创建db
	tmp := make(map[string][]string, 0)
	for _, s := range Field {
		field, _ := getAttrByField(s)
		tmp[s] = field
	}
	res = make(map[string][]string, 6)

	for _, k := range FieldSort {
		res[k] = tmp[k]
	}
	return res, total, err
}

func getAttrByField(field string) (res []string, count int64) {
	db := global.GVA_DB.Model(&hos.HosSportMode{})
	db = db.Distinct(field)
	db.Count(&count)
	db.Find(&res)

	return res, count
}
