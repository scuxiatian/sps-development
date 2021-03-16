package router

import (
	"github.com/gin-gonic/gin"
	v1 "sps-template-server/api/v1"
)

func InitSignatureRouter(Router *gin.RouterGroup)  {
	SignatureRouter := Router.Group("signature")
	{
		SignatureRouter.POST("createSignature", v1.CreateSignature)		// 创建签章
		SignatureRouter.POST("getSignatureList", v1.GetSignatureList)	// 分页获取签章列表
		SignatureRouter.POST("getSignatureById", v1.GetSignatureById)	// 根据id获取签章
		SignatureRouter.PUT("changePassword", v1.ChangePassword)		// 修改签章密码
		SignatureRouter.PUT("updateSignature", v1.UpdateSignature)		// 更新签章
		SignatureRouter.POST("deleteSignature", v1.DeleteSignature)		// 删除签章
	}
}