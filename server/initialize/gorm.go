package initialize

import (
	"gorm.io/gorm"
)

// MultiTenancy entity

func Gorm() *gorm.DB {

	return GormMysql()
	//switch global.GVA_CONFIG.System.DbType {
	//case "mysql":
	//	return GormMysql()
	//default:
	//	return GormMysql()
	//}
}

//
//func RegisterTables() {
//	db := global.GVA_DB
//	err := db.AutoMigrate(
//
//		system.SysApi{},
//		system.SysUser{},
//		system.SysBaseMenu{},
//		system.JwtBlacklist{},
//		system.SysAuthority{},
//		system.SysDictionary{},
//		system.SysOperationRecord{},
//		system.SysAutoCodeHistory{},
//		system.SysDictionaryDetail{},
//		system.SysBaseMenuParameter{},
//		system.SysBaseMenuBtn{},
//		system.SysAuthorityBtn{},
//		system.SysAutoCode{},
//		system.SysExportTemplate{},
//		system.Condition{},
//		system.JoinTemplate{},
//
//		example.ExaFile{},
//		example.ExaCustomer{},
//		example.ExaFileChunk{},
//		example.ExaFileUploadAndDownload{}, hos.SysOperationRecords{}, hos.HosScale{}, hos.HosLocalAsk{}, hos.HosFlow{}, hos.SysUsers{}, hos.HosSportMode{}, hos.HosSportAdvice{}, hos.HosUsers{}, hos.SysDept{}, hos.SysOrg{}, hos.SysPost{}, hos.HosSportClock{}, hos.HosUserPoint{}, hos.HosSportClockCommit{}, hos.HosUsersDianzan{}, hos.HosDashboard{},
//	)
//	if err != nil {
//		global.GVA_LOG.Error("register table failed", zap.Error(err))
//		os.Exit(0)
//	}
//	global.GVA_LOG.Info("register table success")
//}
