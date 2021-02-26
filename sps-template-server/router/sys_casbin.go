package router

import (
	"github.com/gin-gonic/gin"
	v1 "sps-template-server/api/v1"
)

func InitCasbinRouter(Router *gin.RouterGroup)  {
	CasbinRouter := Router.Group("casbin")
	{
		CasbinRouter.POST("getPolicyPathByAuthorityId", v1.GetPolicyPathByAuthorityId)
		CasbinRouter.POST("updateCasbin", v1.UpdateCasbin)
	}
}