// 自动生成模板HosFlow
package hos

import (
     "devops-manage/global"
)

// 会诊流程 结构体  HosFlow
type HosFlow struct {
     global.GVA_MODEL
     Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                   //名称
     Uuid      string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:uuid;size:500;"`                   //uuid
     Uid       string `json:"uid" form:"uid" gorm:"column:uid;comment:关联用户id;size:255;"`                //关联用户id
     Type      string `json:"type" form:"type" gorm:"column:type;comment:类型（快速/普通）;size:500;"`        //类型（快速/普通）
     Step      string `json:"step" form:"step" gorm:"column:step;comment:所属状态;size:255;"`               //所属状态
     ScaleId   *int   `json:"scaleId" form:"scaleId" gorm:"column:scale_id;comment:量表录入id;size:10;"`    //量表录入id
     AskId     *int   `json:"askId" form:"askId" gorm:"column:ask_id;comment:现场关联id;size:10;"`          //现场关联id
     AdviceId  *int   `json:"adviceId" form:"adviceId" gorm:"column:advice_id;comment:运动建议id;size:10;"` //运动建议id
     Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                    //排序
     TenantId  *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`   //租户编号
     CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
     UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
     DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 会诊流程 HosFlow自定义表名 hos_flow
func (HosFlow) TableName() string {
     return "hos_flow"
}

