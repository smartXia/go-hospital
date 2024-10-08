// 自动生成模板SysUsers
package hos

import (
	"devops-manage/global"
)

type SysUsersInfo struct {
	ID       uint   `gorm:"primarykey" json:"ID"` // 主键ID
	Username string `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:191;"`
	NickName string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:191;"` //用户昵称
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:191;"`          //用户手机号
}

// sysUsers表 结构体  SysUsers
type SysUsers struct {
	global.GVA_MODEL
	Uuid           string       `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;size:191;"`                              //用户UUID
	Username       string       `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:191;"`                   //用户登录名
	Password       string       `json:"password" form:"password" gorm:"column:password;comment:用户登录密码;size:191;"`                  //用户登录密码
	NickName       string       `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:191;"`                   //用户昵称
	SideMode       string       `json:"sideMode" form:"sideMode" gorm:"column:side_mode;comment:用户侧边主题;size:191;"`                 //用户侧边主题
	HeaderImg      string       `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;size:191;"`                //用户头像
	BaseColor      string       `json:"baseColor" form:"baseColor" gorm:"column:base_color;comment:基础颜色;size:191;"`                //基础颜色
	AuthorityId    uint         `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:用户角色ID;size:20;"`         //用户角色ID
	Phone          string       `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:191;"`                            //用户手机号
	Email          string       `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:191;"`                             //用户邮箱
	Enable         *int         `json:"enable" form:"enable" gorm:"default:1; column:enable;comment:状态;size:10;"`                  //状态
	Hospital       *int         `json:"hospital" form:"hospital" gorm:"column:hospital;comment:归属医院;size:255;"`                    //归属医院
	Dept           *int         `json:"dept" form:"dept" gorm:"column:dept;comment:部门;size:255;"`                                  //部门
	DeptInfo       SysDept      `json:"deptInfo" form:"deptInfo"  gorm:"foreignKey:id;references:Dept"`                            //部门
	Post           *int         `json:"post" form:"post" gorm:"column:post;comment:职务;size:255;"`                                  //职务
	PostInfo       SysPost      `json:"sysPost" form:"sysPost" gorm:"foreignKey:id;references:Post"`                               //职务
	Birthday       string       `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日;size:255;"`                      //生日
	Sex            string       `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:255;"`                                     //性别
	Address        string       `json:"address" form:"address" gorm:"column:address;comment:住址;size:255;"`                         //住址
	Hometown       string       `json:"hometown" form:"hometown" gorm:"column:hometown;comment:籍贯;size:255;"`                      //籍贯
	Education      string       `json:"education" form:"education" gorm:"column:education;comment:学历;size:255;"`                   //学历
	Profession     string       `json:"profession" form:"profession" gorm:"column:profession;comment:专业;size:500;"`                //专业
	School         string       `json:"school" form:"school" gorm:"column:school;comment:毕业院校;size:500;"`                          //毕业院校
	GraduationTime string       `json:"graduationTime" form:"graduationTime" gorm:"column:graduation_time;comment:毕业时间;size:255;"` //毕业时间
	Desc           string       `json:"desc" form:"desc" gorm:"column:desc;comment:个人简介;size:2000;"`                               //个人简介
	TenantId       *int         `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户id;"`                            //租户id
	CreatedBy      uint         `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                  //创建者
	UpdatedBy      uint         `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                  //更新者
	DeletedBy      *uint        `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                  //删除者
	SysUser        SysUsersInfo `json:"createdByInfo" form:"createdBy" gorm:"foreignKey:id;references:CreatedBy"`                  //创建者

	OrgInfo SysOrg `json:"orgInfo" form:"-" gorm:"foreignKey:id;references:hospital"`
}

// TableName sysUsers表 SysUsers自定义表名 sys_users
func (SysUsers) TableName() string {
	return "sys_users"
}
func (SysUsersInfo) TableName() string {
	return "sys_users"
}
