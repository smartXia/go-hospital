// 自动生成模板HosSportClock
package hos

import (
	"devops-manage/global"
)

// hosSportClock表 结构体  HosSportClock
type HosSportClock struct {
	global.GVA_MODEL
	Uid            string `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:500;"`                                   //用户id
	FlowId         string `json:"flowId" form:"flowId" gorm:"column:flow_id;comment:流程id;size:500;"`                         //流程id
	Name           string `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                                  //名称
	RelationPhotos string `json:"relationPhotos" form:"relationPhotos" gorm:"column:relation_photos;comment:打卡图片;size:255;"` //打卡图片
	Desc           string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                                  //描述
	Enable         *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"`                             //状态
	Sort           *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                   //排序
	TenantId       *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`                    //租户编号
	CreatedBy      *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                  //创建者
	UpdatedBy      *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                  //更新者
	DeletedBy      *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                  //删除者
}

// TableName hosSportClock表 HosSportClock自定义表名 hos_sport_clock
func (HosSportClock) TableName() string {
	return "hos_sport_clock"
}
