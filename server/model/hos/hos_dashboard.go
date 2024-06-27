// 自动生成模板HosDashboard
package hos

import (
	"devops-manage/core/constants"
	"devops-manage/global"
	"github.com/goccy/go-json"
)

// hosDashboard表 结构体  HosDashboard
type HosDashboard struct {
	global.GVA_MODEL
	Name                   string       `json:"name" form:"name" gorm:"column:name;comment:名称;size:500;"`                                                           //名称
	OrgId                  uint         `json:"orgId" form:"orgId" gorm:"column:org_id;comment:机构id;size:10;"`                                                      //机构id
	JizhuduilieTotal       string       `json:"jizhuduilieTotal" form:"jizhuduilieTotal" gorm:"column:jizhuduilie_total;comment:脊柱队列数量;size:500;"`                  //脊柱队列数量
	XinxishagnbaoTotal     string       `json:"xinxishagnbaoTotal" form:"xinxishagnbaoTotal" gorm:"column:xinxishagnbao_total;comment:信息上报数;size:500;"`             //信息上报数
	XianchagnzhenliaoTotal string       `json:"xianchagnzhenliaoTotal" form:"xianchagnzhenliaoTotal" gorm:"column:xianchagnzhenliao_total;comment:运动建议数;size:500;"` //运动建议数
	YundongjianyiTotal     string       `json:"yundongjianyiTotal" form:"yundongjianyiTotal" gorm:"column:yundongjianyi_total;comment:线上打卡数;size:500;"`             //线上打卡数
	XianshagndakaTotal     string       `json:"xianshagndakaTotal" form:"xianshagndakaTotal" gorm:"column:xianshagndaka_total;comment:位置;size:500;"`                //位置
	Diqupaihang            string       `json:"diqupaihang" form:"diqupaihang" gorm:"column:diqupaihang;comment:地区排行;"`                                             //地区排行
	Duilienianling         string       `json:"duilienianling" form:"duilienianling" gorm:"column:duilienianling;comment:队列年龄;"`                                    //队列年龄
	Duilieyanzhongxing     string       `json:"duilieyanzhongxing" form:"duilieyanzhongxing" gorm:"column:duilieyanzhongxing;comment:队列严重性;"`                       //队列严重性
	Duiliefenlei           string       `json:"duiliefenlei" form:"duiliefenlei" gorm:"column:duiliefenlei;comment:队列分类;"`                                          //队列分类
	Desc                   string       `json:"desc" form:"desc" gorm:"column:desc;comment:描述;size:500;"`                                                           //描述
	Enable                 *int         `json:"enable" form:"enable" gorm:"column:enable;comment:状态;size:10;"`                                                      //状态
	Sort                   *int         `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`                                                            //排序
	TenantId               *int         `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户编号;size:10;"`                                             //租户编号
	CreatedBy              uint         `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`                                           //创建者
	UpdatedBy              uint         `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`                                           //更新者
	DeletedBy              *int         `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`                                           //删除者
	SysUsersInfo           SysUsersInfo `json:"sysUsersInfo" form:"sysUsersInfo" gorm:"foreignKey:id;references:CreatedBy"`
	SysOrg                 SysOrg       `json:"orgInfo" form:"orgInfo" gorm:"foreignKey:id;references:OrgId"`
}

type Diqupaihang struct {
	Province      string `json:"province"`
	WeekIncrease  int    `json:"week"`
	MonthIncrease int    `json:"month"`
	YearIncrease  int    `json:"year"`
	Total         int    `json:"total"`
}
type Duilienianling struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}
type Duilieyanzhongxing struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}
type Duiliefenlei struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func (hosd *HosDashboard) BuildInitDashBoardDate(hosDashboard *HosDashboard) (err error) {
	if hosDashboard.JizhuduilieTotal == "" {
		hosd.JizhuduilieTotal = "1000"
	}
	if hosDashboard.XinxishagnbaoTotal == "" {
		hosd.XinxishagnbaoTotal = "1000"
	}
	if hosDashboard.XianchagnzhenliaoTotal == "" {
		hosd.XianchagnzhenliaoTotal = "1000"
	}
	if hosDashboard.YundongjianyiTotal == "" {
		hosd.YundongjianyiTotal = "1000"
	}
	if hosDashboard.XianshagndakaTotal == "" {
		hosd.XianshagndakaTotal = "1000"
	}

	var a []Duilienianling
	var b []Duilieyanzhongxing
	var c []Duiliefenlei
	var d []Diqupaihang
	for _, s := range constants.Duilienianling {
		a = append(a, Duilienianling{
			Name:   s,
			Number: 0,
		})
	}
	for _, s := range constants.Duilieyanzhongxing {
		b = append(b, Duilieyanzhongxing{
			Name:   s,
			Number: 0,
		})
	}
	for _, s := range constants.Duiliefenlei {
		c = append(c, Duiliefenlei{
			Name:   s,
			Number: 0,
		})
	}

	for _, s := range constants.Diqupaihang {
		d = append(d, Diqupaihang{
			Province:      s,
			WeekIncrease:  0,
			MonthIncrease: 0,
			YearIncrease:  0,
			Total:         0,
		})
	}
	if hosDashboard.Duilienianling == "" {
		aa, _ := json.Marshal(a)
		hosd.Duilienianling = string(aa)
	}
	if hosDashboard.Duilieyanzhongxing == "" {
		bb, _ := json.Marshal(b)
		hosd.Duilieyanzhongxing = string(bb)
	}

	if hosDashboard.Duiliefenlei == "" {
		cc, _ := json.Marshal(c)
		hosd.Duiliefenlei = string(cc)

	}
	if hosDashboard.Diqupaihang == "" {
		dd, _ := json.Marshal(d)
		hosd.Diqupaihang = string(dd)
	}

	if err != nil {
		return err
	}
	return nil
}

// TableName hosDashboard表 HosDashboard自定义表名 hos_dashboard
func (HosDashboard) TableName() string {
	return "hos_dashboard"
}
