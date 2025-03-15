package hos

import (
	"devops-manage/global"
	"devops-manage/model/common/scope"
	"devops-manage/model/hos"
	hosReq "devops-manage/model/hos/request"
	"devops-manage/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"strconv"
)

type HosDashboardService struct {
}

// CreateHosDashboard 创建hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) CreateHosDashboard(hosDashboard *hos.HosDashboard, ctx *gin.Context, init int) (err error, d *hos.HosDashboard) {
	hosDashboard.CreatedBy = utils.GetUserID(ctx)
	//if init == 1 {
	hosDashboard.BuildInitDashBoardDate(hosDashboard)
	//}
	err = global.GVA_DB.Scopes(scope.TenantSaveScope(ctx)).Create(hosDashboard).Error
	return err, hosDashboard
}

// DeleteHosDashboard 删除hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) DeleteHosDashboard(ID string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&hos.HosDashboard{}, "id = ?", ID).Error
	return err
}

// DeleteHosDashboardByIds 批量删除hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) DeleteHosDashboardByIds(IDs []string, ctx *gin.Context) (err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Delete(&[]hos.HosDashboard{}, "id in ?", IDs).Error
	return err
}

// UpdateHosDashboard 更新hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) UpdateHosDashboard(hosDashboard hos.HosDashboard, ctx *gin.Context) (err error) {
	hosDashboard.UpdatedBy = utils.GetUserID(ctx)
	err = global.GVA_DB.Model(&hos.HosDashboard{}).Scopes(scope.TenantScope(ctx)).Where("id = ?", hosDashboard.ID).Updates(&hosDashboard).Error
	return err
}

// GetHosDashboard 根据ID获取hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) GetHosDashboard(ID string, ctx *gin.Context) (hosDashboard hos.HosDashboard, err error) {
	err = global.GVA_DB.Scopes(scope.TenantScope(ctx)).Where("id = ?", ID).First(&hosDashboard).Error
	return
}

// GetHosDashboardInfoList 分页获取hosDashboard表记录
// Author [piexlmax](https://github.com/piexlmax)
func (hosDashboardService *HosDashboardService) GetHosDashboardInfoList(info hosReq.HosDashboardSearch, ctx *gin.Context) (list []hos.HosDashboard, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&hos.HosDashboard{}).Scopes(scope.TenantScope(ctx))
	var hosDashboards []hos.HosDashboard
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.OrgId != nil {
		db = db.Where("org_id = ?", info.OrgId)
	}
	if info.Enable != nil {
		db = db.Where("enable = ?", info.Enable)
	}
	db = db.Preload("SysOrg")
	db = db.Preload("SysUsersInfo")
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset).Order("id desc")
	}

	err = db.Find(&hosDashboards).Error
	return hosDashboards, total, err
}

func (hosDashboardService *HosDashboardService) GetCurrentDashBoardInfo(info hosReq.HosDashboardSearch, ctx *gin.Context) (hosDashboard hos.HosDashboard, err error) {
	// 创建db
	db := global.GVA_DB.Model(&hos.HosDashboard{}).Scopes(scope.TenantScope(ctx))
	// 如果有条件搜索 下方会自动创建搜索语句
	tid := utils.GetTenantId(ctx)
	if info.OrgId != nil {
		db = db.Where("org_id = ?", info.OrgId)
	} else {
		db = db.Where("org_id = ?", tid)
	}
	db = db.Preload("SysOrg")
	db = db.Preload("SysUsersInfo")
	err = db.First(&hosDashboard).Error
	if hosDashboard.ID == 0 {
		return hos.HosDashboard{}, nil
	}
	if *hosDashboard.Enable == 1 {
		return hosDashboard, err
	} else {
		var g errgroup.Group

		//脊椎队列 就是flow数量
		var flowNum int64
		g.Go(func() error {
			global.GVA_DB.Model(&hos.HosFlow{}).Scopes(scope.TenantScope(ctx)).Count(&flowNum)
			hosDashboard.JizhuduilieTotal = strconv.FormatInt(flowNum, 10)
			return nil
		})
		//信息上报就是sacle数量
		var scaleNum int64
		g.Go(func() error {
			global.GVA_DB.Model(&hos.HosScale{}).Scopes(scope.TenantScope(ctx)).Count(&scaleNum)
			hosDashboard.XinxishagnbaoTotal = strconv.FormatInt(scaleNum, 10)
			return nil
		})
		//现场诊疗就是local
		var localNum int64
		g.Go(func() error {
			global.GVA_DB.Model(&hos.HosLocalAsk{}).Scopes(scope.TenantScope(ctx)).Count(&localNum)
			hosDashboard.XianchagnzhenliaoTotal = strconv.FormatInt(localNum, 10)
			return nil
		})
		//运动建议就是建议
		var adviceNum int64
		g.Go(func() error {
			global.GVA_DB.Model(&hos.HosSportAdvice{}).Scopes(scope.TenantScope(ctx)).Count(&adviceNum)
			hosDashboard.YundongjianyiTotal = strconv.FormatInt(adviceNum, 10)
			return nil
		})
		//打卡数量就是总打卡数量
		var clockNum int64
		g.Go(func() error {
			global.GVA_DB.Model(&hos.HosSportClock{}).Scopes(scope.TenantScope(ctx)).Count(&clockNum)
			hosDashboard.XianshagndakaTotal = strconv.FormatInt(clockNum, 10)
			return nil
		})
		//默认处理
		//开始处理 其他的表格数据
		var diqupaihang []hos.Diqupaihang
		//SELECT province, city, COUNT(*) AS user_count
		//FROM hos_users
		//GROUP BY province, city
		//ORDER BY total DESC;
		//1.地区排行
		g.Go(func() error {
			global.GVA_DB.Model(&hos.HosUsers{}).
				Scopes(scope.TenantScope(ctx)).
				Select("province, city, COUNT(*) AS total").
				Where("province != ?", "").Where("city != ?", "").
				Group("province, city").
				Order("total desc").
				Scan(&diqupaihang)
			if len(diqupaihang) != 0 {
				marshal, err := json.Marshal(diqupaihang)
				if err != nil {
					return err
				}
				hosDashboard.Diqupaihang = string(marshal)
			}
			return nil
		})

		//2.年龄范围
		var duilienianling []hos.Duilienianling
		g.Go(func() error {
			//遍历 duilienianling 分类 年龄范围
			//0≦幼儿≦2：2<儿童≦6：6<少年≦14：14<青年≦
			//35、35<中年≦60、老年>60
			global.GVA_DB.
				Table("hos_users").             // 使用 Table 代替 Model，方便后续 Scope
				Scopes(scope.TenantScope(ctx)). // 作用于整个查询
				Select(`
        CASE 
            WHEN age BETWEEN 0 AND 2 THEN '幼儿'
            WHEN age BETWEEN 3 AND 6 THEN '儿童'
            WHEN age BETWEEN 7 AND 14 THEN '少年'
            WHEN age BETWEEN 15 AND 35 THEN '青年'
            WHEN age BETWEEN 36 AND 60 THEN '中年'
            WHEN age > 60 THEN '老年'
        END AS name,
        COUNT(*) AS number
    `).
				Where("age != ?", 0).
				Group(`
        CASE 
            WHEN age BETWEEN 0 AND 2 THEN '幼儿'
            WHEN age BETWEEN 3 AND 6 THEN '儿童'
            WHEN age BETWEEN 7 AND 14 THEN '少年'
            WHEN age BETWEEN 15 AND 35 THEN '青年'
            WHEN age BETWEEN 36 AND 60 THEN '中年'
            WHEN age > 60 THEN '老年'
        END
    `).
				Order("number DESC").
				Scan(&duilienianling)

			if len(duilienianling) != 0 {
				marshal, err := json.Marshal(duilienianling)
				if err != nil {
					return err
				}
				hosDashboard.Duilienianling = string(marshal)
			}
			return nil
		})
		//3.队列严重性
		var duilieyanzhongxing []hos.Duilieyanzhongxing
		g.Go(func() error {
			global.GVA_DB.
				Table("hos_local_ask").         // 使用 Table 代替 Model，方便后续 Scope
				Scopes(scope.TenantScope(ctx)). // 作用于整个查询
				Raw(`
    SELECT 
        '轻' AS name, COUNT(CASE WHEN jizhucewan LIKE '%轻%' THEN 1 END) AS number
    FROM hos_local_ask
    UNION
    SELECT 
        '轻度温和' AS name, COUNT(CASE WHEN jizhucewan LIKE '%轻度温和%' THEN 1 END) AS number
    FROM hos_local_ask
    UNION
    SELECT 
        '温和' AS name, COUNT(CASE WHEN jizhucewan LIKE '%温和%' THEN 1 END) AS number
    FROM hos_local_ask
    UNION
    SELECT 
        '温和严重' AS name, COUNT(CASE WHEN jizhucewan LIKE '%温和严重%' THEN 1 END) AS number
    FROM hos_local_ask
    UNION
    SELECT 
        '建议手术' AS name, COUNT(CASE WHEN jizhucewan LIKE '%建议手术%' THEN 1 END) AS number
    FROM hos_local_ask
    UNION
    SELECT 
        '必须手术' AS name, COUNT(CASE WHEN jizhucewan LIKE '%必须手术%' THEN 1 END) AS number
    FROM hos_local_ask
`).Scan(&duilieyanzhongxing)
			if len(duilieyanzhongxing) != 0 {
				marshal, err := json.Marshal(duilieyanzhongxing)
				if err != nil {
					return err
				}
				hosDashboard.Duilieyanzhongxing = string(marshal)
			}
			return nil
		})
		//4.病形统计
		var duiliefenlei []hos.Duiliefenlei
		g.Go(func() error {
			// Execute the SQL query to fetch the counts for each condition
			global.GVA_DB.
				Table("hos_local_ask"). // 使用 Table 代替 Model，方便后续 Scope
				Scopes(scope.TenantScope(ctx)).
				Raw(`
    SELECT 
        '脊柱侧弯' AS name, 
        COUNT(*) AS number
    FROM hos_local_ask
    WHERE jizhucewan REGEXP '[一-龥]'
    UNION
    SELECT 
        '矢状面障碍' AS name, 
        COUNT(*) AS number
    FROM hos_local_ask
    WHERE shizhuangmianzhangai REGEXP '[一-龥]'
    UNION
    SELECT 
        '脊椎化脱' AS name, 
        COUNT(*) AS number
    FROM hos_local_ask
    WHERE jizhuihuatuo REGEXP '[一-龥]'
    UNION
    SELECT 
        '长短脚' AS name, 
        COUNT(*) AS number
    FROM hos_local_ask
    WHERE changduanjiao REGEXP '[一-龥]'
    UNION
    SELECT 
        '背痛' AS name, 
        COUNT(*) AS number
    FROM hos_local_ask
    WHERE beitong REGEXP '[一-龥]'
    UNION
    SELECT 
        '其他' AS name, 
        COUNT(*) AS number
    FROM hos_local_ask
    WHERE jizhucewan NOT REGEXP '[一-龥]' 
      AND shizhuangmianzhangai NOT REGEXP '[一-龥]' 
      AND jizhuihuatuo NOT REGEXP '[一-龥]' 
      AND changduanjiao NOT REGEXP '[一-龥]' 
      AND beitong NOT REGEXP '[一-龥]'
`).Scan(&duiliefenlei)

			// Handle potential errors
			if len(duiliefenlei) != 0 {
				// Optionally, marshal the result into JSON
				marshal, _ := json.Marshal(duiliefenlei)
				// Assign the result to a field in your dashboard or struct
				hosDashboard.Duiliefenlei = string(marshal)
			}
			// Return nil if no error
			return nil
		})

		if err := g.Wait(); err != nil {
			return hos.HosDashboard{}, err
		}

		return hosDashboard, err

	}
}

//func BuildDa0ta() {
//	var a []hos.Duilienianling
//	var b []hos.Duilieyanzhongxing
//	var c []hos.Duiliefenlei
//	var d []hos.Diqupaihang
//}

func (hosDashboardService *HosDashboardService) GetHosDashboardDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	orgId := make([]map[string]any, 0)
	global.GVA_DB.Table("sys_org").Select("name as label,id as value").Scan(&orgId)
	res["orgId"] = orgId
	return
}
