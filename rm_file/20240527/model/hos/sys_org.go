// 自动生成模板SysOrg
package hos

import (
	"devops-manage/global"
)

// sysOrg表 结构体  SysOrg
type SysOrg struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:名称;size:191;" binding:"required"` //名称
	Address   string `json:"address" form:"address" gorm:"column:address;comment:详细地址;size:255;"`         //详细地址
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:19;"`               //状态
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`                             //排序
	ManageId  *int   `json:"manageId" form:"manageId" gorm:"column:manage_id;comment:管理人id;"`             //管理人id
	ParentId  *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:parentId;"`          //parentId
	TenantId  *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;"`              //租户编号
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName sysOrg表 SysOrg自定义表名 sys_org
func (SysOrg) TableName() string {
	return "sys_org"
}
