// 自动生成模板HosScale
package hos

import (
	"devops-manage/global"
)

// hosScale表 结构体  HosScale
type HosScale struct {
	global.GVA_MODEL
	Name           string       `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                                  //名称
	HosUserId      *int         `json:"hos_user_id" form:"hos_user_id" gorm:"column:hos_user_id;comment:关联用户id;size:10;"`          //关联用户id
	BirthStatus    string       `json:"birthStatus" form:"birthStatus" gorm:"column:birth_status;comment:出生状态;size:500;"`          //出生状态
	BirthMode      string       `json:"birthMode" form:"birthMode" gorm:"column:birth_mode;comment:出生方式;size:500;"`                //出生方式
	GrowStatus     string       `json:"growStatus" form:"growStatus" gorm:"column:grow_status;comment:发育状态;size:500;"`             //发育状态
	FamilyMedical  string       `json:"familyMedical" form:"familyMedical" gorm:"column:family_medical;comment:家族病史;size:500;"`    //家族病史
	MedicalHistory string       `json:"medicalHistory" form:"medicalHistory" gorm:"column:medical_history;comment:历史病史;size:500;"` //历史病史
	LesionHistory  string       `json:"lesionHistory" form:"lesionHistory" gorm:"column:lesion_history;comment:病灶记录;size:500;"`    //病灶记录
	Other          string       `json:"other" form:"other" gorm:"column:other;comment:其他说明;size:500;"`                             //其他说明
	RelationPhotos string       `json:"relationPhotos" form:"relationPhotos" gorm:"column:relation_photos;comment:关联图片;size:255;"` //关联图片
	Desc           string       `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                                  //描述
	Enable         *int         `json:"enable" form:"enable" gorm:"default:1; column:enable;comment:状态;size:10;"`                  //状态
	Sort           *int         `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                   //排序
	TenantId       *int         `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`                    //租户编号
	CreatedBy      uint         `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                  //创建者
	UpdatedBy      uint         `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                  //更新者
	DeletedBy      *uint        `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                  //删除者
	SysUser        SysUsersInfo `json:"createdByInfo" form:"sysUsers" gorm:"foreignKey:id;references:CreatedBy"`                   //创建者
}

// TableName hosScale表 HosScale自定义表名 hos_scale
func (HosScale) TableName() string {
	return "hos_scale"
}
