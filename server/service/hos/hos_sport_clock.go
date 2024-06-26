package hos

import (
	"devops-manage/core/constants"
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"sort"
	"time"
)

type HosSportClockService struct {
}

// CreateHosSportClock 创建hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) CreateHosSportClock(hosSportClock *hos.HosSportClock, ctx *gin.Context) (err error, d *hos.HosSportClock) {
	hosSportClock.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosSportClock).Error
	return err, hosSportClock
}

// DeleteHosSportClock 删除hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) DeleteHosSportClock(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosSportClock{}, "id = ?", ID).Error
	return err
}

// DeleteHosSportClockByIds 批量删除hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) DeleteHosSportClockByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosSportClock{}, "id in ?", IDs).Error
	return err
}

// UpdateHosSportClock 更新hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) UpdateHosSportClock(hosSportClock hos.HosSportClock, ctx *gin.Context) (err error) {
	hosSportClock.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosSportClock{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosSportClock.ID).Updates(&hosSportClock).Error
	//埋点
	go UpdateUserPoint(ctx, hosSportClock.ID, constants.POINT)

	return err
}

// GetHosSportClock 根据ID获取hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) GetHosSportClock(ID string, ctx *gin.Context) (hosSportClock hos.HosSportClock, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosSportClock).Error
	return
}

func (hosSportClockService *HosSportClockService) GetHosSportDistinctClockList(info hosReq.HosSportClockSearch, ctx *gin.Context) (list []hos.HosSportClock, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosSportClock{}).Scopes(scope.TenantScope(ctx))
	var hosSportClocks []hos.HosSportClock
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.HosUserId != nil {
		db = db.Where("hos_user_id = ?", info.HosUserId)
	}
	if info.FlowId != nil {
		db = db.Where("flow_id = ?", info.FlowId)
	}
	if info.AdviceId != nil {
		db = db.Where("advice_id = ?", info.AdviceId)
	}
	if info.ClockStartTime != "" {
		db = db.Where("clock_start_time > ?", info.ClockStartTime)
	}
	if info.ClockEndTime != "" {
		db = db.Where("clock_end_time < ?", info.ClockEndTime)
	}
	if err != nil {
		return
	}
	db = db.Preload("HosSportAdvice")
	db = db.Preload("HosFlow")
	db = db.Preload("SysUser")
	db = db.Group("flow_id")
	err = db.Count(&total).Error

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("hos_sport_clock.id desc")
	}

	err = db.Find(&hosSportClocks).Error
	return hosSportClocks, total, err
}

// GetHosSportClockInfoList 分页获取hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) GetHosSportClockInfoList(info hosReq.HosSportClockSearch, ctx *gin.Context) (list []hos.HosSportClock, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosSportClock{}).Scopes(scope.TenantScope(ctx))
	var hosSportClocks []hos.HosSportClock
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.HosUserId != nil {
		db = db.Where("hos_user_id = ?", info.HosUserId)
	}
	if info.FlowId != nil {
		db = db.Where("flow_id = ?", info.FlowId)
	}
	if info.AdviceId != nil {
		db = db.Where("advice_id = ?", info.AdviceId)
	}
	if info.ClockStartTime != "" {
		db = db.Where("clock_start_time > ?", info.ClockStartTime)
	}
	if info.ClockEndTime != "" {
		db = db.Where("clock_end_time < ?", info.ClockEndTime)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}
	db = db.Preload("HosSportAdvice").Preload("HosSportAdvice.HosSportMode")
	err = db.Find(&hosSportClocks).Error
	return hosSportClocks, total, err
}

// GetCurrentUserHosSportClockList 分页获取hosSportClock表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockService *HosSportClockService) GetCurrentUserHosSportClockList(info hosReq.HosSportClockSearch, ctx *gin.Context) (list interface{}, total int64, err error) {
	// 创建db
	//  先找到 自己 建议周期内的所有建议 然后按照建议id 来筛选出来 这个建议这天需要打卡的周期 比如 2
	// 筛选条件 就可以 判断今天属于 哪个打卡 周期
	db := global.GVA_DB.Model(&hos.HosSportClock{}).Scopes(scope.TenantScope(ctx))
	var hosSportClocks []hos.HosSportClock
	start := time.Now().Format("2006-01-02")

	uids, err := GetUserIds(ctx)
	if len(uids) == 0 {
		return list, 0, nil
	}

	db = db.Where("hos_user_id in ?", uids)
	//db = db.Where("hos_user_id = ?", 49)
	db = db.Where("clock_start_time <= ?", start)
	db = db.Preload("HosSportAdvice").Preload("SysUser")
	db = db.Preload("HosSportAdvice").Preload("HosSportAdvice.HosSportMode")
	err = db.Find(&hosSportClocks).Order("id desc").Error

	mapByDate := make(map[int][]hos.HosSportClock, 0)
	if len(hosSportClocks) != 0 {
		for _, v := range hosSportClocks {
			mapByDate[v.Cishu] = append(mapByDate[v.Cishu], v)
		}
	}
	// 提取 map 的所有 keys
	keys := make([]int, 0, len(mapByDate))
	for key := range mapByDate {
		keys = append(keys, key)
	}
	// 对 keys 进行排序
	sort.Ints(keys)
	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}
	var c []ClockPlan
	for _, key := range keys {

		current := ClockPlan{
			Name:              mapByDate[key][0].Name,
			Image:             mapByDate[key][0].HosSportAdvice.Images,
			JieDuan:           fmt.Sprintf("%s~%s", mapByDate[key][0].HosSportAdvice.Jianyitime, mapByDate[key][0].HosSportAdvice.Fuzhenriqi),
			DoctorInfo:        mapByDate[key][0].SysUser,
			SportAdviceDetail: mapByDate[key],
			Jindu:             process(mapByDate[key]),
			Status:            "over",
		}
		parseS, err := time.Parse("2006-01-02", mapByDate[key][0].ClockStartTime)
		parseE, err := time.Parse("2006-01-02", mapByDate[key][0].ClockEndTime)
		if err != nil {
			return nil, 0, err
		}
		if parseS.Before(time.Now()) && parseE.After(time.Now()) {
			current.Status = "process"
		}
		c = append(c, current)
	}

	return c, total, err
}

func process(hs []hos.HosSportClock) int {
	i := 0
	for _, v := range hs {
		if v.RelationPhotos != "" {
			i++
		}
	}
	total := len(hs)
	res := float64(i) / float64(total) * 100
	return int(math.Round(res))
}

type ClockPlan struct {
	Name              string              `json:"name"`
	JieDuan           string              `json:"jie_duan"`
	DoctorInfo        hos.SysUsersInfo    `json:"doctor_info"`
	SportAdviceDetail []hos.HosSportClock `json:"sport_advice_detail"`
	Image             string              `json:"image"`
	Jindu             int                 `json:"process"`
	Status            string              `json:"status"`
}

func (hosSportClockService *HosSportClockService) GetHosSportClockDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	hosUserId := make([]map[string]any, 0)
	global.GVA_DB.Table("hos_users").Select("name as label,username as value").Scan(&hosUserId)
	res["hosUserId"] = hosUserId
	return
}
