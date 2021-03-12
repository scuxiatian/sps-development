package router

import (
	"github.com/gin-gonic/gin"
	v1 "sps-template-server/api/v1"
)

func InitSignatureRouter(Router *gin.RouterGroup)  {
	SignatureRouter := Router.Group("signature")
	{
		SignatureRouter.POST("createSignature", v1.CreateSignature)	// 创建签章
	}
}