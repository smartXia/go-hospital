// 自动生成模板SysOrg
package hos

import (
	"devops-manage/global"
)

// sysOrg表 结构体  SysOrg
type SysOrg struct {
	global.GVA_MODEL
	Name      string       `json:"name" form:"name" gorm:"column:name;comment:名称;size:255;"`                   //名称
	Desc      string       `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:255;"`                   //描述
	Img       string       `json:"img" form:"img" gorm:"column:img;comment:图标;size:255;"`                      //图标
	Address   string       `json:"address" form:"address" gorm:"column:address;comment:详细地址;size:255;"`        //详细地址
	Enable    *int         `json:"enable" form:"enable" gorm:"column:enable;comment:状态;"`                      //状态
	Sort      *int         `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:19;"`                    //排序
	ManageId  *int         `json:"manageId" form:"manageId" gorm:"column:manage_id;comment:管理人id;size:19;"`    //管理人id
	ParentId  *int         `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:parentId;size:19;"` //parentId
	TenantId  *int         `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`     //租户编号
	CreatedBy uint         `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:19;"`   //创建者
	UpdatedBy uint         `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:19;"`   //更新者
	DeletedBy uint         `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:19;"`   //删除者
	SysUser   SysUsersInfo `json:"createdByInfo" form:"createdBy" gorm:"foreignKey:id;references:CreatedBy"`   //创建者
}

// TableName sysOrg表 SysOrg自定义表名 sys_org
func (SysOrg) TableName() string {
	return "sys_org"
}
