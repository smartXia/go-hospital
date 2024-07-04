// 自动生成模板HosSportClock
package hos

import (
	"devops-manage/global"
)

// hosSportClock表 结构体  HosSportClock
type HosSportClock struct {
	global.GVA_MODEL
	HosUserId      *int           `json:"hosUserId" form:"hosUserId" gorm:"column:hos_user_id;comment:患者id;size:10;"`                    //患者id
	FlowId         *int           `json:"flowId" form:"flowId" gorm:"column:flow_id;comment:流程id;size:10;"`                              //流程id
	AdviceId       uint           `json:"adviceId" form:"adviceId" gorm:"column:advice_id;comment:建议ID;size:10;"`                        //建议ID
	Cishu          int            `json:"cishu" form:"cishu" gorm:"column:cishu;comment:次数;size:10;"`                                    //建议ID
	Name           string         `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                                      //名称
	RelationPhotos string         `json:"relationPhotos" form:"relationPhotos" gorm:"column:relation_photos;comment:打卡图片;size:255;"`     //打卡图片
	Desc           string         `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                                      //描述
	ClockStartTime string         `json:"clockStartTime" form:"clockStartTime" gorm:"column:clock_start_time;comment:可以开始打卡时间;size:50;"` //可以开始打卡时间
	ClockEndTime   string         `json:"clockEndTime" form:"clockEndTime" gorm:"column:clock_end_time;comment:打卡截至时间;size:50;"`         //打卡截至时间
	Enable         *int           `json:"enable" form:"enable" gorm:"default:1; column:enable;comment:状态;size:10;"`                      //状态
	Sort           *int           `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                       //排序
	TenantId       *int           `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`                        //租户编号
	CreatedBy      uint           `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                      //创建者
	UpdatedBy      uint           `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                      //更新者
	DeletedBy      *int           `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                      //删除者
	HosSportAdvice HosSportAdvice `json:"hosSportAdviceInfo" form:"hosSportAdvice" gorm:"foreignKey:id;references:AdviceId"`             //运动建议id
	SysUser        SysUsersInfo   `json:"createdByInfo" form:"sysUsers" gorm:"foreignKey:id;references:CreatedBy"`                       //创建者
	HosFlow        HosFlow        `json:"hosFlowInfo" form:"hosFlowInfo" gorm:"foreignKey:id;references:FlowId"`                         //创建者
}

// TableName hosSportClock表 HosSportClock自定义表名 hos_sport_clock
func (HosSportClock) TableName() string {
	return "hos_sport_clock"
}
