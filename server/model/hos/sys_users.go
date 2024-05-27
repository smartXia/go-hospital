// 自动生成模板SysUsers
package hos

import (
	"devops-manage/global"
)

// sysUsers表 结构体  SysUsers
type SysUsers struct {
	global.GVA_MODEL
	Uuid           string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;size:255;"`                              //用户UUID
	Username       string `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:255;"`                   //用户登录名
	Password       string `json:"password" form:"password" gorm:"column:password;comment:用户登录密码;size:255;"`                  //用户登录密码
	NickName       string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:255;"`                   //用户昵称
	SideMode       string `json:"sideMode" form:"sideMode" gorm:"column:side_mode;comment:用户侧边主题;size:255;"`                 //用户侧边主题
	HeaderImg      string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;size:500;"`                //用户头像
	BaseColor      string `json:"baseColor" form:"baseColor" gorm:"column:base_color;comment:基础颜色;size:255;"`                //基础颜色
	AuthorityId    *int   `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:用户角色ID;size:20;"`         //用户角色ID
	Phone          string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:255;"`                            //用户手机号
	Email          string `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:255;"`                             //用户邮箱
	Enable         *int   `json:"enable" form:"enable" gorm:"column:enable;comment:用户是否被冻结 1正常 2冻结;size:255;"`               //用户是否被冻结 1正常 2冻结
	Hospital       string `json:"hospital" form:"hospital" gorm:"column:hospital;comment:归属医院;size:255;"`                    //归属医院
	Dept           string `json:"dept" form:"dept" gorm:"column:dept;comment:部门;size:255;"`                                  //部门
	Post           string `json:"post" form:"post" gorm:"column:post;comment:职务;size:255;"`                                  //职务
	Birthday       string `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;size:255;"`                      //生日
	Sex            string `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:255;"`                                     //性别
	Address        string `json:"address" form:"address" gorm:"column:address;comment:住址;size:255;"`                         //住址
	Hometown       string `json:"hometown" form:"hometown" gorm:"column:hometown;comment:籍贯;size:255;"`                      //籍贯
	Education      string `json:"education" form:"education" gorm:"column:education;comment:学历;size:255;"`                   //学历
	Profession     string `json:"profession" form:"profession" gorm:"column:profession;comment:专业;size:500;"`                //专业
	School         string `json:"school" form:"school" gorm:"column:school;comment:毕业院校;size:500;"`                          //毕业院校
	GraduationTime string `json:"graduationTime" form:"graduationTime" gorm:"column:graduation_time;comment:毕业时间;size:255;"` //毕业时间
	Desc           string `json:"desc" form:"desc" gorm:"column:desc;comment:个人简介;size:2000;"`                               //个人简介
	TenantId       *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户id;size:10;"`                    //租户id
	CreatedBy      uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy      uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy      uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName sysUsers表 SysUsers自定义表名 sys_users
func (SysUsers) TableName() string {
	return "sys_users"
}
