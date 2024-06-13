// 自动生成模板HosSportClockCommit
package hos

import (
	"devops-manage/global"
)

// hosSportClockCommit表 结构体  HosSportClockCommit
type HosSportClockCommit struct {
	global.GVA_MODEL
	FlowId    string `json:"flowId" form:"flowId" gorm:"column:flow_id;comment:关联的流程id;size:500;"`     //关联的流程id
	AdviceId  string `json:"adviceId" form:"adviceId" gorm:"column:advice_id;comment:建议id;size:500;"`  //建议id
	Uid       string `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:500;"`                  //用户id
	Commit    string `json:"commit" form:"commit" gorm:"column:commit;comment:评论;size:500;"`           //评论
	Desc      string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                 //描述
	Enable    *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"`            //状态
	Sort      *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                  //排序
	TenantId  *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`   //租户编号
	CreatedBy *uint  `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"` //创建者
	UpdatedBy *uint  `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"` //更新者
	DeletedBy *uint  `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"` //删除者
}

// TableName hosSportClockCommit表 HosSportClockCommit自定义表名 hos_sport_clock_commit
func (HosSportClockCommit) TableName() string {
	return "hos_sport_clock_commit"
}
