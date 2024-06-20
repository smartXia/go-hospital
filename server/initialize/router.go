package initialize

import (
	"devops-manage/global"
	"devops-manage/middleware"
	"devops-manage/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(middleware.Cors())
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	global.GVA_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{

		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)
		systemRouter.InitInitRouter(PublicGroup)
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.TenantMiddleware())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)
		systemRouter.InitJwtRouter(PrivateGroup)
		systemRouter.InitUserRouter(PrivateGroup)
		systemRouter.InitMenuRouter(PrivateGroup)
		systemRouter.InitSystemRouter(PrivateGroup)
		systemRouter.InitCasbinRouter(PrivateGroup)
		systemRouter.InitAutoCodeRouter(PrivateGroup)
		systemRouter.InitAuthorityRouter(PrivateGroup)
		systemRouter.InitSysDictionaryRouter(PrivateGroup)
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)
		exampleRouter.InitCustomerRouter(PrivateGroup)
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)

	}
	{
		hosRouter := router.RouterGroupApp.Hos

		hosRouter.InitSysOperationRecordsRouter(PrivateGroup, PublicGroup)

		hosRouter.InitHosScaleRouter(PrivateGroup, PublicGroup)
		hosRouter.InitHosLocalAskRouter(PrivateGroup, PublicGroup)

		hosRouter.InitHosSportClockRouter(PrivateGroup, PublicGroup)
		hosRouter.InitHosSportClockCommitRouter(PrivateGroup, PublicGroup)
		hosRouter.InitHosUserPointRouter(PrivateGroup, PublicGroup)
		hosRouter.InitHosFlowRouter(PrivateGroup, PublicGroup)
		hosRouter.InitSysUsersRouter(PrivateGroup, PublicGroup)

		hosRouter.InitHosSportModeRouter(PrivateGroup, PublicGroup)
		hosRouter.InitHosSportAdviceRouter(PrivateGroup, PublicGroup)
		hosRouter.InitHosUsersRouter(PrivateGroup, PublicGroup)
		hosRouter.InitSysDeptRouter(PrivateGroup, PublicGroup)
		hosRouter.InitSysOrgRouter(PrivateGroup, PublicGroup)
		hosRouter.InitSysPostRouter(PrivateGroup, PublicGroup)

	}

	global.GVA_LOG.Info("router register success")
	return Router
}
