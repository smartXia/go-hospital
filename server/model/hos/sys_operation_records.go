// 自动生成模板SysOperationRecords
package hos

import (
	"devops-manage/global"
)

// sysOperationRecords表 结构体  SysOperationRecords
type SysOperationRecords struct {
	global.GVA_MODEL
	Ip           string `json:"ip" form:"ip" gorm:"column:ip;comment:请求ip;size:191;"`                                //请求ip
	Method       string `json:"method" form:"method" gorm:"column:method;comment:请求方法;size:191;"`                    //请求方法
	Path         string `json:"path" form:"path" gorm:"column:path;comment:请求路径;size:191;"`                          //请求路径
	Status       *int   `json:"status" form:"status" gorm:"column:status;comment:请求状态;size:19;"`                     //请求状态
	Latency      *int   `json:"latency" form:"latency" gorm:"column:latency;comment:延迟;size:19;"`                    //延迟
	Agent        string `json:"agent" form:"agent" gorm:"column:agent;comment:代理;"`                                  //代理
	ErrorMessage string `json:"errorMessage" form:"errorMessage" gorm:"column:error_message;comment:错误信息;size:191;"` //错误信息
	Body         string `json:"body" form:"body" gorm:"column:body;comment:请求Body;"`                                 //请求Body
	Resp         string `json:"resp" form:"resp" gorm:"column:resp;comment:响应Body;"`                                 //响应Body
	UserId       *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:20;"`                    //用户id
	Action       string `json:"action" form:"action" gorm:"column:action;comment:操作;size:255;"`                      //操作
	Browser      string `json:"browser" form:"browser" gorm:"column:browser;comment:浏览器;size:255;"`                  //浏览器
	TenantId     *int   `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户id;size:10;"`              //租户id
	Os           string `json:"os" form:"os" gorm:"column:os;comment:系统;size:255;"`                                  //系统
}

// TableName sysOperationRecords表 SysOperationRecords自定义表名 sys_operation_records
func (SysOperationRecords) TableName() string {
	return "sys_operation_records"
}
