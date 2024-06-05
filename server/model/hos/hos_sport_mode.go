// 自动生成模板HosSportMode
package hos

import (
	"devops-manage/global"
)

// hosSportMode表 结构体  HosSportMode
type HosSportMode struct {
	global.GVA_MODEL
	Name           string `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                                  //名称
	Category       string `json:"category" form:"category" gorm:"column:category;comment:分类;size:255;"`                      //分类
	Buwei          string `json:"buwei" form:"buwei" gorm:"column:buwei;comment:部位;size:500;"`                               //部位
	Tiduan         string `json:"tiduan" form:"tiduan" gorm:"column:tiduan;comment:体段;size:500;"`                            //体段
	Xingdong       string `json:"xingdong" form:"xingdong" gorm:"column:xingdong;comment:行动;size:500;"`                      //行动
	Fangxiang      string `json:"fangxiang" form:"fangxiang" gorm:"column:fangxiang;comment:方向;size:500;"`                   //方向
	Weizhi         string `json:"weizhi" form:"weizhi" gorm:"column:weizhi;comment:位置;size:500;"`                            //位置
	Detail         string `json:"detail" form:"detail" gorm:"column:detail;comment:详情;"`                                     //详情
	RelationPhotos string `json:"relationPhotos" form:"relationPhotos" gorm:"column:relation_photos;comment:关联图片;size:255;"` //关联图片
	Source         string `json:"source" form:"source" gorm:"column:source;comment:来源;size:255;"`                            //来源
	Desc           string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                                  //描述
	Enable         *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"`                             //状态
	Sort           *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                   //排序
	TenantId       *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`                    //租户编号
	CreatedBy      *int   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                  //创建者
	UpdatedBy      *int   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                  //更新者
	DeletedBy      *int   `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                  //删除者
}

// TableName hosSportMode表 HosSportMode自定义表名 hos_sport_mode
func (HosSportMode) TableName() string {
	return "hos_sport_mode"
}
