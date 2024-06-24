// 自动生成模板HosSportAdvice
package hos

import (
	"devops-manage/global"
)

// hosSportAdvice表 结构体  HosSportAdvice
type HosSportAdvice struct {
	global.GVA_MODEL
	FlowId       *int         `json:"flowId" form:"flowId" gorm:"column:flow_id;comment:流程id;size:10;"`                   //流程id
	HosUserId    *int         `json:"hosUserId" form:"hosUserId" gorm:"column:hos_user_id;comment:患者id;size:11;"`         //患者id
	SportModeId  *int         `json:"sportModeId" form:"sportModeId" gorm:"column:sport_mode_id;comment:模型 id;size:10;"`  //模型 id
	Name         string       `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                           //名称
	Jianyitime   string       `json:"jianyitime" form:"jianyitime" gorm:"column:jianyitime;comment:建议起止时间;size:500;"`     //建议起止时间
	Jianyizhouqi int          `json:"jianyizhouqi" form:"jianyizhouqi" gorm:"column:jianyizhouqi;comment:建议周期;size:500;"` //建议周期
	Fuzhenriqi   string       `json:"fuzhenriqi" form:"fuzhenriqi" gorm:"column:fuzhenriqi;comment:复诊时间;size:500;"`       //复诊时间
	Detail       string       `json:"detail" form:"detail" gorm:"column:detail;comment:建议详情;size:500;"`                   //建议详情
	Images       string       `json:"images" form:"images" gorm:"column:images;comment:建议图;size:500;"`                    //建议图
	Source       string       `json:"source" form:"source" gorm:"column:source;comment:来源（self-system）;size:500;"`        //来源（self-system）
	SportMode    string       `json:"sportMode" form:"sportMode" gorm:"column:sport_mode;comment:建议;size:500;"`           //建议
	Period       string       `json:"period" form:"period" gorm:"column:period;comment:周期;size:500;"`                     //周期
	Desc         string       `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                           //描述
	Enable       *int         `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"`                      //状态
	Sort         *int         `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                            //排序
	SyncWx       *int         `json:"syncWx" form:"syncWx" gorm:"column:sync_wx;comment:是否同步微信;size:10;"`                 //是否同步微信
	TenantId     *int         `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`             //租户编号
	CreatedBy    uint         `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`           //创建者
	UpdatedBy    uint         `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`           //更新者
	DeletedBy    *int         `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`           //删除者
	SysUser      SysUsersInfo `json:"createdByInfo" form:"sysUsers" gorm:"foreignKey:id;references:CreatedBy"`            //创建者
	HosSportMode HosSportMode `json:"HosSportModInfo" form:"HosSportModInfo" gorm:"foreignKey:id;references:SportModeId"` //运动建议id
}

// TableName hosSportAdvice表 HosSportAdvice自定义表名 hos_sport_advice
func (HosSportAdvice) TableName() string {
	return "hos_sport_advice"
}
