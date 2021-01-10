package router

import (
	"github.com/gin-gonic/gin"
	v1 "sps-template-server/api/v1"
)

func InitUserRouter(Router *gin.RouterGroup)  {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", v1.Register)
	}
}
