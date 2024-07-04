// 自动生成模板HosUsersDianzan
package hos

import (
	"devops-manage/global"
)

// hosUsersDianzan表 结构体  HosUsersDianzan
type HosUsersDianzan struct {
	global.GVA_MODEL
	HosUserId *int   `json:"hosUserId" form:"hosUserId" gorm:"column:hos_user_id;comment:用户id;size:10;"` //用户id
	CommitId  string `json:"commitId" form:"commitId" gorm:"column:commit_id;comment:评论id;size:1000;"`   //评论id
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                   //描述
	Enable    *int   `json:"enable" form:"enable" gorm:"default:1; column:enable;comment:状态;size:10;"`   //状态
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                    //排序
	TenantId  int    `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`     //租户编号
	CreatedBy uint   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`   //创建者
	UpdatedBy uint   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`   //更新者
	DeletedBy uint   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`   //删除者
}

// TableName hosUsersDianzan表 HosUsersDianzan自定义表名 hos_users_dianzan
func (HosUsersDianzan) TableName() string {
	return "hos_users_dianzan"
}
