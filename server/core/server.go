package core

import (
	"devops-manage/global"
	"devops-manage/initialize"
	"devops-manage/service/system"
	"fmt"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 gin-vue
	当前版本:v2.6.4
	文档地址:http://127.0.0.1%s/swagger/index.html
	运行地址:http://127.0.0.1:8080
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
