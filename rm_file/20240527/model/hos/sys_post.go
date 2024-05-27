// 自动生成模板SysPost
package hos

import (
	"devops-manage/global"
)

// sysPost表 结构体  SysPost
type SysPost struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:岗位名称;" binding:"required"` //岗位名称
	Code      string `json:"code" form:"code" gorm:"column:code;comment:岗位标识;" binding:"required"` //岗位标识
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`              //排序
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:启用;size:10;"`        //启用
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`       //备注
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName sysPost表 SysPost自定义表名 sys_post
func (SysPost) TableName() string {
	return "sys_post"
}
