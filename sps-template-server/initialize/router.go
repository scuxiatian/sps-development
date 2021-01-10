package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "sps-template-server/docs"
	"sps-template-server/global"
	"sps-template-server/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	global.SdLog.Info("use middleware logger")

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.SdLog.Info("register swagger handler")
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("")
	{
		router.InitUserRouter(PrivateGroup)
	}
	global.SdLog.Info("router register success")
	return Router
}
