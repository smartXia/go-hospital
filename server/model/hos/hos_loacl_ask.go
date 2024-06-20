// 自动生成模板HosLoaclAsk
package hos

import (
	"devops-manage/global"
)

// hosLoaclAsk表 结构体  HosLoaclAsk
type HosLoaclAsk struct {
	global.GVA_MODEL
	Name                 string `json:"name" form:"name" gorm:"column:name;comment:名称;size:255;"`                                           //名称
	HosUserId            *int   `json:"hos_user_id" form:"hos_user_id" gorm:"column:hos_user_id;comment:关联用户id;size:10;"`                   //关联用户id
	Jizhucewan           string `json:"jizhucewan" form:"jizhucewan" gorm:"column:jizhucewan;comment:脊柱侧弯;"`                                //脊柱侧弯
	Shizhuangmianzhangai string `json:"shizhuangmianzhangai" form:"shizhuangmianzhangai" gorm:"column:shizhuangmianzhangai;comment:矢状面障碍;"` //矢状面障碍
	Jizhuihuatuo         string `json:"jizhuihuatuo" form:"jizhuihuatuo" gorm:"column:jizhuihuatuo;comment:脊椎滑脱;"`                          //脊椎滑脱
	Changduanjiao        string `json:"changduanjiao" form:"changduanjiao" gorm:"column:changduanjiao;comment:长短脚;"`                        //长短脚
	Beitong              string `json:"beitong" form:"beitong" gorm:"column:beitong;comment:背痛;"`                                           //背痛
	Other                string `json:"other" form:"other" gorm:"column:other;comment:其他说明;size:500;"`                                      //其他说明
	Desc                 string `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                                           //描述
	Enable               *int   `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"`                                      //状态
	Sort                 *int   `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                            //排序
	TenantId             *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`                             //租户编号
	CreatedBy            uint   `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                           //创建者
	UpdatedBy            uint   `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                           //更新者
	DeletedBy            *uint  `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                           //删除者
}

// TableName hosLoaclAsk表 HosLoaclAsk自定义表名 hos_loacl_ask
func (HosLoaclAsk) TableName() string {
	return "hos_loacl_ask"
}

//func (h *HosLoaclAsk) BeforeCreate(tx *gorm.DB) (err error) {
//	ctx := tx.Statement.Context
//	userID := utils.GetUserID(ctx)
//	h.CreatedBy = userID
//	return
//}
//func getUserIDFromContext(ctx context.Context) uint {
//	userID, ok := ctx.Value("userID").(uint)
//	if !ok {
//		// 处理获取用户ID失败的情况，这里假设返回0
//		return 0
//	}
//	return userID
//}
