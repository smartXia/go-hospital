// 自动生成模板SysDept
package hos

import (
	"devops-manage/global"
)

// sysDept表 结构体  SysDept
type SysDept struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:部门名称;" binding:"required"`                   //部门名称
	Address   string `json:"address" form:"address" gorm:"column:address;comment:详细地址;size:255;" binding:"required"` //详细地址
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;"`                                  //状态
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                //排序
	Parent    *int   `json:"parent" form:"parent" gorm:"column:parent;comment:父级部门;size:10;"`                        //父级部门
	TenantId  *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`                 //租户编号
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName sysDept表 SysDept自定义表名 sys_dept
func (SysDept) TableName() string {
	return "sys_dept"
}
