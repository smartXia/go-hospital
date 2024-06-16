package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"github.com/gin-gonic/gin"
)

type HosSportClockCommitService struct {
}

// CreateHosSportClockCommit 创建hosSportClockCommit表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockCommitService *HosSportClockCommitService) CreateHosSportClockCommit(hosSportClockCommit *hos.HosSportClockCommit, ctx *gin.Context) (err error, h *hos.HosSportClockCommit) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosSportClockCommit).Error
	return err, hosSportClockCommit
}

// DeleteHosSportClockCommit 删除hosSportClockCommit表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockCommitService *HosSportClockCommitService) DeleteHosSportClockCommit(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosSportClockCommit{}, "id = ?", ID).Error
	return err
}

// DeleteHosSportClockCommitByIds 批量删除hosSportClockCommit表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockCommitService *HosSportClockCommitService) DeleteHosSportClockCommitByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosSportClockCommit{}, "id in ?", IDs).Error
	return err
}

// UpdateHosSportClockCommit 更新hosSportClockCommit表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockCommitService *HosSportClockCommitService) UpdateHosSportClockCommit(hosSportClockCommit hos.HosSportClockCommit, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Model(&hos.HosSportClockCommit{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosSportClockCommit.ID).Updates(&hosSportClockCommit).Error
	return err
}

// GetHosSportClockCommit 根据ID获取hosSportClockCommit表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockCommitService *HosSportClockCommitService) GetHosSportClockCommit(ID string, ctx *gin.Context) (hosSportClockCommit hos.HosSportClockCommit, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosSportClockCommit).Error
	return
}

// GetHosSportClockCommitInfoList 分页获取hosSportClockCommit表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosSportClockCommitService *HosSportClockCommitService) GetHosSportClockCommitInfoList(info hosReq.HosSportClockCommitSearch, ctx *gin.Context) (list []hos.HosSportClockCommit, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosSportClockCommit{}).Scopes(scope.TenantScope(ctx))
	var hosSportClockCommits []hos.HosSportClockCommit
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FlowId != "" {
		db = db.Where("flow_id = ?", info.FlowId)
	}
	if info.AdviceId != "" {
		db = db.Where("advice_id = ?", info.AdviceId)
	}
	if info.HosUserId != "" {
		db = db.Where("hos_user_id = ?", info.HosUserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&hosSportClockCommits).Error
	return hosSportClockCommits, total, err
}
