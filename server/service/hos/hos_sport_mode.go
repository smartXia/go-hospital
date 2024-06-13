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
	id := utils.GetUserID(ctx)
	hosSportMode.CreatedBy = &id
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosSportMode).Error
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
	err = global.GVA_DB.Model(&hos.HosSportMode{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosSportMode.ID).Updates(&hosSportMode).Error
	return err
}

// GetHosSportMode 根据ID获取hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) GetHosSportMode(ID string, ctx *gin.Context) (hosSportMode hos.HosSportMode, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosSportMode).Error
	return
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

// GetHosSportModeInfoList 分页获取hosSportMode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportModeService *HosSportModeService) GetHosSportModeInfoList(info hosReq.HosSportModeSearch, ctx *gin.Context) (list []hos.HosSportMode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosSportMode{}).Scopes(scope.TenantScope(ctx))
	var hosSportModes []hos.HosSportMode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Category != "" {
		category := strings.Split(info.Category, ",")
		db = db.Where("category in ?", category)
	}
	if info.Buwei != "" {
		buwei := strings.Split(info.Buwei, ",")
		db = db.Where("buwei in ?", buwei)
	}
	if info.Tiduan != "" {
		tidaun := strings.Split(info.Tiduan, ",")
		db = db.Where("tiduan in ?", tidaun)
	}
	if info.Fangxiang != "" {
		fangxiang := strings.Split(info.Fangxiang, ",")
		db = db.Where("fangxiang in ?", fangxiang)
	}
	if info.Weizhi != "" {
		weizhi := strings.Split(info.Weizhi, ",")

		db = db.Where("weizhi in ?", weizhi)
	}

	if info.Source != "" {
		db = db.Where("source = ?", info.Source)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Order("id desc").Limit(limit).Offset(offset).Preload("User")
	}

	err = db.Find(&hosSportModes).Debug().Error
	return hosSportModes, total, err
}
