package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"github.com/gin-gonic/gin"
)

type HosFlowService struct {
}

// CreateHosFlow 创建hosFlow表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService) CreateHosFlow(hosFlow *hos.HosFlow, ctx *gin.Context) (err error, f *hos.HosFlow) {
	hosFlow.Uuid = utils.UniqueId()
	hosFlow.CreatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Scopes(scope.TenantSaveScope(ctx)).Create(hosFlow).Error
	return err, hosFlow
}

// DeleteHosFlow 删除hosFlow表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService) DeleteHosFlow(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosFlow{}, "id = ?", ID).Error
	return err
}

// DeleteHosFlowByIds 批量删除hosFlow表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService) DeleteHosFlowByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosFlow{}, "id in ?", IDs).Error
	return err
}

// UpdateHosFlow 更新hosFlow表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService) UpdateHosFlow(hosFlow hos.HosFlow, ctx *gin.Context) (err error) {
	hosFlow.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosFlow{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosFlow.ID).Updates(&hosFlow).Error
	return err
}

// GetHosFlow 根据ID获取hosFlow表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService) GetHosFlow(ID string, ctx *gin.Context) (hosFlow hos.HosFlow, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosFlow).Error
	return
}

// GetHosFlowInfoList 分页获取hosFlow表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService) GetHosFlowInfoList(info hosReq.HosFlowSearch, ctx *gin.Context) (list []hos.HosFlow, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db

	db := global.GVA_DB.Model(&hos.HosFlow{}).Scopes(scope.TableTenantScope(ctx, "hos_flow"))
	var hosFlows []hos.HosFlow
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("hos_flow.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 如果需要联表 hos_users
	if info.Name != "" || info.Phone != "" || info.YuyueType != "" {
		db.Joins("JOIN `hos_users`  ON hos_flow.hos_user_id = hos_users.id")
	}
	if info.YuyueType != "" {
		db = db.Where("hos_users.yuyue_type = ?", info.YuyueType)
	}
	if info.Phone != "" {
		db = db.Where("hos_users.phone = ?", info.Phone)
	}
	if info.Name != "" {
		db = db.Where("hos_users.username = ?", info.Name)
	}
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Type != "" {
		db = db.Where("hos_flow.type = ?", info.Type)
	}
	if info.HosUserId != nil {
		db = db.Where("hos_flow.hos_user_id = ?", info.HosUserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("hos_flow.id desc")
	}
	db.Preload("HosScale").Preload("HosScale.SysUser").
		Preload("HosLocalAsk").Preload("HosLocalAsk.SysUser").
		Preload("HosSportAdvice").Preload("HosSportAdvice.SysUser").
		Preload("SysUser").
		Preload("HosUser").Preload("HosUser.SysUser")

	err = db.Find(&hosFlows).Error
	return hosFlows, total, err
}

func (hosFlowService *HosFlowService) GetCurrentHosFlowInfoList(info hosReq.HosFlowSearch, ctx *gin.Context) (list []hos.HosFlow, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosFlow{}).Scopes(scope.TenantScope(ctx))
	var hosFlows []hos.HosFlow
	// 如果有条件搜索 下方会自动创建搜索语句

	uids, err := GetUserIds(ctx)
	if len(uids) == 0 {
		return list, 0, nil
	}

	db = db.Where("hos_user_id in ?", uids)

	if info.AdviceId != nil {
		db = db.Where("advice_id = ?", info.AdviceId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}
	db.Preload("HosScale").Preload("HosScale.SysUser").
		Preload("HosLocalAsk").Preload("HosLocalAsk.SysUser").
		Preload("HosSportAdvice").Preload("HosSportAdvice.SysUser").
		Preload("SysUser").
		Preload("HosUser").Preload("HosUser.SysUser")

	err = db.Find(&hosFlows).Error
	return hosFlows, total, err
}

func (hosFlowService *HosFlowService) GetHosFlowDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	adviceId := make([]map[string]any, 0)
	global.GVA_DB.Table("hos_sport_advice").Select("id as label,id as value").Scan(&adviceId)
	res["adviceId"] = adviceId
	askId := make([]map[string]any, 0)
	global.GVA_DB.Table("hos_loacl_ask").Select("id as label,id as value").Scan(&askId)
	res["askId"] = askId
	scaleId := make([]map[string]any, 0)
	global.GVA_DB.Table("hos_scale").Select("id as label,id as value").Scan(&scaleId)
	res["scaleId"] = scaleId
	uid := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_users").Select("nick_name as label,id as value").Scan(&uid)
	res["uid"] = uid
	return
}
