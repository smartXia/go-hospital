package hos

import (
	"devops-manage/core/constants"
	"devops-manage/model/hos"
	"github.com/gin-gonic/gin"
)

func GetUserIds(ctx *gin.Context) (list []uint, err error) {
	service := HosUsersService{}
	uids, err := service.GetHosUsersIdsByPhone(ctx)
	if len(uids) == 0 {
		return list, nil
	}
	return uids, nil
}

func UpdateUserPoint(ctx *gin.Context, huid uint, action string) (err error) {
	service := HosUserPointService{}
	p := hos.HosUserPoint{
		HosUserId: huid,
		Event:     action,
		Change:    constants.ActionEventChange[action],
		Name:      constants.ActionEventChangeCn[action],
	}
	err, _ = service.CreateHosUserPoint(&p, ctx)
	if err != nil {
		return err
	}
	return err
}
