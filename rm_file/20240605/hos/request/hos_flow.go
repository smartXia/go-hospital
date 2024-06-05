package request

import (
	"devops-manage/model/common/request"
	"time"
	
)

type HosFlowSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Uuid  string `json:"uuid" form:"uuid" `
                      ScaleId  *int `json:"scaleId" form:"scaleId" `
                      AskId  *int `json:"askId" form:"askId" `
                      AdviceId  *int `json:"adviceId" form:"adviceId" `
                      TenantId  *int `json:"tenantId" form:"tenantId" `
    request.PageInfo
}
