// 自动生成模板HosUsers
package hos

import (
	"devops-manage/global"
)

// hosUsers表 结构体  HosUsers
type HosUsers struct {
	global.GVA_MODEL
	Uuid            string       `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;size:255;"`                                    //用户UUID
	Username        string       `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:255;"`                         //用户登录名
	Password        string       `json:"password" form:"password" gorm:"column:password;comment:用户登录密码;size:255;"`                        //用户登录密码
	NickName        string       `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:255;"`                         //用户昵称
	HeaderImg       string       `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;size:255;"`                      //用户头像
	Phone           string       `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:20;"`                                   //用户手机号
	JianhuPhone     string       `json:"jianhuPhone" form:"jianhuPhone" gorm:"column:jianhu_phone;comment:监护人手机号;size:255;"`              //监护人手机号
	Jianhuren       string       `json:"jianhuren" form:"jianhuren" gorm:"column:jianhuren;comment:监护人;size:255;"`                        //监护人
	WxUuid          string       `json:"wxUuid" form:"wxUuid" gorm:"column:wx_uuid;comment:微信uuid;size:255;"`                             //微信uuid
	Email           string       `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:255;"`                                   //用户邮箱
	Enable          *int         `json:"enable" form:"enable" gorm:"default:1; column:enable;comment:状态;size:10;"`                        //状态
	YuyueType       string       `json:"yuyueType" form:"yuyueType" gorm:"column:yuyue_type;comment:预约类型;size:255;"`                      //预约类型
	YuyueTime       string       `json:"yuyueTime" form:"yuyueTime" gorm:"column:yuyue_time;comment:预约时间;size:255;"`                      //预约时间
	Birthday        string       `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;size:255;"`                            //生日
	Sex             string       `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:255;"`                                           //性别
	Address         string       `json:"address" form:"address" gorm:"column:address;comment:住址;size:255;"`                               //住址
	Province        string       `json:"province" form:"province" gorm:"column:province;comment:省;size:255;"`                             //省
	City            string       `json:"city" form:"city" gorm:"column:city;comment:市;size:255;"`                                         //市
	Hometown        string       `json:"hometown" form:"hometown" gorm:"column:hometown;comment:籍贯;size:255;"`                            //籍贯
	Education       string       `json:"education" form:"education" gorm:"column:education;comment:学历;size:255;"`                         //学历
	CardNo          string       `json:"cardNo" form:"cardNo" gorm:"column:card_no;comment:身份证号码;size:255;"`                              //身份证号码
	Age             *int         `json:"age" form:"age" gorm:"column:age;comment:年龄;size:255;"`                                           //年龄
	WomanPeriodDate string       `json:"womanPeriodDate" form:"womanPeriodDate" gorm:"column:woman_period_date;comment:女性初潮日期;size:255;"` //女性初潮日期
	Height          *int         `json:"height" form:"height" gorm:"column:height;comment:身高;size:255;"`                                  //身高
	Weight          *int         `json:"weight" form:"weight" gorm:"column:weight;comment:体重;size:255;"`                                  //体重
	RegisterHos     string       `json:"registerHos" form:"registerHos" gorm:"column:register_hos;comment:第一次登记医院;size:255;"`             //第一次登记医院
	LatelyHos       string       `json:"latelyHos" form:"latelyHos" gorm:"column:lately_hos;comment:近一次登记医院;size:255;"`                   //近一次登记医院
	TenantId        *int         `json:"tenantId" form:"tenantId" gorm:"default:0;column:tenant_id;comment:租户id;"`                        //租户id
	CreatedBy       uint         `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                        //创建者
	UpdatedBy       uint         `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                        //更新者
	DeletedBy       *int         `json:"-" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                                //删除者
	SysUser         SysUsersInfo `json:"createdByInfo" form:"createdBy" gorm:"foreignKey:id;references:CreatedBy"`                        //创建者
}

// TableName hosUsers表 HosUsers自定义表名 hos_users
func (HosUsers) TableName() string {
	return "hos_users"
}
