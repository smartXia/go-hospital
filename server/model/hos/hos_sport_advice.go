// 自动生成模板HosSportAdvice
package hos

import (
	"devops-manage/global"
)

// hosSportAdvice表 结构体  HosSportAdvice
type HosSportAdvice struct {
	global.GVA_MODEL
	FlowId    *int   `json:"flowId" form:"flowId" gorm:"column:flow_id;comment:流程id;size:10;"`            //流程id
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                    //名称
	Period    string `json:"period" form:"period" gorm:"column:period;comment:周期;size:500;"`              //周期
	SportMode string `json:"sportMode" form:"sportMode" gorm:"column:sport_mode;comment:建议;size:500;"`    //建议
	Detail    string `json:"detail" form:"detail" gorm:"column:detail;comment:建议详情;size:500;"`            //建议详情
	Images    string `json:"images" form:"images" gorm:"column:images;comment:建议图;size:500;"`             //建议图
	Source    string `json:"source" form:"source" gorm:"column:source;comment:来源（self-system）;size:500;"` //来源（self-system）
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                    //描述
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"`               //状态
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                     //排序
	TenantId  *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`      //租户编号
	CreatedBy *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`    //创建者
	UpdatedBy *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`    //更新者
	DeletedBy *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`    //删除者
}

// TableName hosSportAdvice表 HosSportAdvice自定义表名 hos_sport_advice
func (HosSportAdvice) TableName() string {
	return "hos_sport_advice"
}
