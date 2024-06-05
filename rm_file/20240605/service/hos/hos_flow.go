package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
    hosReq "devops-manage/model/hos/request"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

type HosFlowService struct {
}

// CreateHosFlow 创建会诊流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService) CreateHosFlow(hosFlow *hos.HosFlow, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Create(hosFlow).Error
	return err
}

// DeleteHosFlow 删除会诊流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService)DeleteHosFlow(ID string,userID uint, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hos.HosFlow{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&hos.HosFlow{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteHosFlowByIds 批量删除会诊流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService)DeleteHosFlowByIds(IDs []string,deleted_by uint, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&hos.HosFlow{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&hos.HosFlow{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateHosFlow 更新会诊流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService)UpdateHosFlow(hosFlow hos.HosFlow, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Model(&hos.HosFlow{}).Scopes(scope.TenantScope(ctx)).Where("id = ?",hosFlow.ID).Updates(&hosFlow).Error
	return err
}

// GetHosFlow 根据ID获取会诊流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService)GetHosFlow(ID string, ctx *gin.Context) (hosFlow hos.HosFlow, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosFlow).Error
	return
}

// GetHosFlowInfoList 分页获取会诊流程记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosFlowService *HosFlowService)GetHosFlowInfoList(info hosReq.HosFlowSearch, ctx *gin.Context) (list []hos.HosFlow, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&hos.HosFlow{}).Scopes(scope.TenantScope(ctx))
    var hosFlows []hos.HosFlow
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Uuid != "" {
        db = db.Where("uuid = ?",info.Uuid)
    }
    if info.ScaleId != nil {
        db = db.Where("scale_id = ?",info.ScaleId)
    }
    if info.AskId != nil {
        db = db.Where("ask_id = ?",info.AskId)
    }
    if info.AdviceId != nil {
        db = db.Where("advice_id = ?",info.AdviceId)
    }
    if info.TenantId != nil {
        db = db.Where("tenant_id = ?",info.TenantId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&hosFlows).Error
	return  hosFlows, total, err
}
func (hosFlowService *HosFlowService)GetHosFlowDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   uid := make([]map[string]any, 0)
       global.GVA_DB.Table("sys_users").Select("nick_name as label,id as value").Scan(&uid)
	   res["uid"] = uid
	return
}