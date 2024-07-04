// 自动生成模板HosSportClockCommit
package hos

import (
	"devops-manage/global"
)

// hosSportClockCommit表 结构体  HosSportClockCommit
type HosSportClockCommit struct {
	global.GVA_MODEL
	FlowId       int          `json:"flowId" form:"flowId" gorm:"column:flow_id;comment:关联的流程id;size:500;"`        //关联的流程id
	HosUserId    *int         `json:"hosUserId" form:"hosUserId" gorm:"column:hos_user_id;comment:患者id;size:10;"`  //患者id
	AdviceId     *int         `json:"adviceId" form:"adviceId" gorm:"column:advice_id;comment:建议id;size:10;"`      //建议id
	ClockId      *int         `json:"clockId" form:"clockId" gorm:"column:clock_id;comment:打卡id;size:10;"`         //打卡id
	Commit       string       `json:"commit" form:"commit" gorm:"column:commit;comment:评论;size:1000;"`             //评论
	Dianzan      *int         `json:"dianzan" form:"dianzan" gorm:"column:dianzan;comment:点赞数;size:10;"`           //点赞数
	Desc         string       `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                    //描述
	Enable       *int         `json:"enable" form:"enable" gorm:"default:1; column:enable;comment:状态;size:10;"`    //状态
	Sort         *int         `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                     //排序
	TenantId     *int         `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`      //租户编号
	CreatedBy    uint         `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`    //创建者
	UpdatedBy    uint         `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`    //更新者
	DeletedBy    *int         `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`    //删除者
	SysUsersInfo SysUsersInfo `json:"createdByInfo" form:"SysUsersInfo" gorm:"foreignKey:id;references:CreatedBy"` //创建者

}

// TableName hosSportClockCommit表 HosSportClockCommit自定义表名 hos_sport_clock_commit
func (HosSportClockCommit) TableName() string {
	return "hos_sport_clock_commit"
}
