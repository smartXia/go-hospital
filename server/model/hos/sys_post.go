// 自动生成模板SysPost
package hos

import (
	"devops-manage/global"
)

// sysPost表 结构体  SysPost
type SysPost struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:岗位名称;size:255;"`             //岗位名称
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`               //描述
	Code      string `json:"code" form:"code" gorm:"column:code;comment:岗位标识;size:255;"`             //岗位标识
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`                        //排序
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:启用;"`                  //启用
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`         //备注
	CreatedBy *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;"`       //创建者
	UpdatedBy *uint  `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;"`       //更新者
	DeletedBy *uint  `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;"`       //删除者
	TenantId  *uint  `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户id;size:19;"` //租户id
}

// TableName sysPost表 SysPost自定义表名 sys_post
func (SysPost) TableName() string {
	return "sys_post"
}
