package system

import (
	"devops-manage/global"
	"devops-manage/model/hos"
	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.GVA_MODEL
	UUID           uuid.UUID      `json:"uuid" form:"-"  gorm:"index;comment:用户UUID"`                                           // 用户UUID
	Username       string         `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password       string         `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName       string         `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode       string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg      string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor      string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	AuthorityId    uint           `json:"authorityId" gorm:"default:0;comment:用户角色ID"`                                          // 用户角色ID
	Authority      SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities    []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone          string         `json:"phone"  gorm:"comment:用户手机号"`                        // 用户手机号
	Email          string         `json:"email"  gorm:"comment:用户邮箱"`                         // 用户邮箱
	Hospital       int            `json:"hospital"  gorm:"comment:hospital;column:hospital" ` // 用户邮箱
	Enable         int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`    //用户是否被冻结 1正常 2冻结
	SysOrg         hos.SysOrg     `json:"orgInfo"  gorm:"foreignKey:id;references:Hospital"`
	Dept           *int           `json:"dept" form:"dept" gorm:"column:dept;comment:部门;size:255;"`                                  //部门
	DeptInfo       hos.SysDept    `json:"deptInfo" form:"deptInfo"  gorm:"foreignKey:id;references:Dept"`                            //部门
	Post           *int           `json:"post" form:"post" gorm:"column:post;comment:职务;size:255;"`                                  //职务
	PostInfo       hos.SysPost    `json:"postInfo" form:"sysPost" gorm:"foreignKey:id;references:Post"`                              //职务
	Sex            string         `json:"sex" form:"sex" gorm:"column:sex;comment:性别;size:255;"`                                     //性别
	Address        string         `json:"address" form:"address" gorm:"column:address;comment:住址;size:255;"`                         //住址
	Hometown       string         `json:"hometown" form:"hometown" gorm:"column:hometown;comment:籍贯;size:255;"`                      //籍贯
	Education      string         `json:"education" form:"education" gorm:"column:education;comment:学历;size:255;"`                   //学历
	Profession     string         `json:"profession" form:"profession" gorm:"column:profession;comment:专业;size:500;"`                //专业
	School         string         `json:"school" form:"school" gorm:"column:school;comment:毕业院校;size:500;"`                          //毕业院校
	GraduationTime string         `json:"graduationTime" form:"graduationTime" gorm:"column:graduation_time;comment:毕业时间;size:255;"` //毕业时间
	Desc           string         `json:"desc" form:"desc" gorm:"column:desc;comment:个人简介;size:2000;"`                               //个人简介
	CreatedBy      uint           `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                  //创建者
}

func (SysUser) TableName() string {
	return "sys_users"
}
