// 自动生成模板SysDept
package hos

import (
	"devops-manage/global"
)

// sysDept表 结构体  SysDept
type SysDept struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:部门名称;size:191;"`          //部门名称
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`            //描述
	Address   string `json:"address" form:"address" gorm:"column:address;comment:详细地址;size:255;"` //详细地址
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:19;"`       //状态
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`                     //排序
	ManageId  *int   `json:"manageId" form:"manageId" gorm:"column:manage_id;comment:管理人id;"`     //管理人id
	ParentId  *int   `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:上级部门;"`      //上级部门
	TenantId  *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;"`      //租户编号
	CreatedBy *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;"`    //创建者
	UpdatedBy *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;"`    //更新者
	DeletedBy *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;"`    //删除者
}

// TableName sysDept表 SysDept自定义表名 sys_dept
func (SysDept) TableName() string {
	return "sys_dept"
}
