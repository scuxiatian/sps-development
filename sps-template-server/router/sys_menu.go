package router

import (
	"github.com/gin-gonic/gin"
	"sps-template-server/api/v1"
)

func InitMenuRouter(Router *gin.RouterGroup)  {
	MenuRouter := Router.Group("menu")
	{
		MenuRouter.POST("getMenu", v1.GetMenu)
	}
}
