package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "sps-template-server/docs"
	"sps-template-server/global"
	"sps-template-server/middleware"
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
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitUserRouter(PrivateGroup)			// 注册用户路由
		router.InitMenuRouter(PrivateGroup)			// 注册menu路由
		router.InitJwtRouter(PrivateGroup)			// jwt相关路由
		router.InitAuthorityRouter(PrivateGroup)	// 注册角色路由
		router.InitApiRouter(PrivateGroup)			// 注册功能api路由
	}
	global.SdLog.Info("router register success")
	return Router
}
