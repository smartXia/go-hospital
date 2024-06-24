package request

import (
	"devops-manage/model/common/request"
	"time"
	
)

type HosSportClockCommitSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      FlowId  string `json:"flowId" form:"flowId" `
                      HosUserId  *int `json:"hosUserId" form:"hosUserId" `
                      AdviceId  *int `json:"adviceId" form:"adviceId" `
                      ClockId  *int `json:"clockId" form:"clockId" `
                      CreatedBy  *int `json:"createdBy" form:"createdBy" `
    request.PageInfo
}
