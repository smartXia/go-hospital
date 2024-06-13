// 自动生成模板HosUserPoint
package hos

import (
	"devops-manage/global"
)

// hosUserPoint表 结构体  HosUserPoint
type HosUserPoint struct {
	global.GVA_MODEL
	Uid       string `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:500;"`                  //用户id
	FlowId    string `json:"flowId" form:"flowId" gorm:"column:flow_id;comment:流程id;size:500;"`        //流程id
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                 //名称
	Event     string `json:"event" form:"event" gorm:"column:event;comment:事件;size:500;"`              //事件
	Change    string `json:"change" form:"change" gorm:"column:change;comment:积分变化;size:500;"`         //积分变化
	Total     string `json:"total" form:"total" gorm:"column:total;comment:共计积分;size:500;"`            //共计积分
	TotalUse  string `json:"totalUse" form:"totalUse" gorm:"column:total_use;comment:可用积分;size:500;"`  //可用积分
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                 //描述
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"`            //状态
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                  //排序
	TenantId  *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`   //租户编号
	CreatedBy *uint  `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"` //创建者
	UpdatedBy *uint  `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"` //更新者
	DeletedBy *uint  `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"` //删除者
}

// TableName hosUserPoint表 HosUserPoint自定义表名 hos_user_point
func (HosUserPoint) TableName() string {
	return "hos_user_point"
}
