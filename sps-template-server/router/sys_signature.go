package router

import (
	"github.com/gin-gonic/gin"
	v1 "sps-template-server/api/v1"
)

func InitSignatureRouter(Router *gin.RouterGroup)  {
	SignatureRouter := Router.Group("signature")
	{
		SignatureRouter.POST("createSignature", v1.CreateSignature)					// 创建签章
		SignatureRouter.POST("getSignatureList", v1.GetSignatureList)				// 分页获取签章列表
		SignatureRouter.POST("getSignatureById", v1.GetSignatureById)				// 根据id获取签章
		SignatureRouter.PUT("changePassword", v1.ChangePassword)					// 修改签章密码
		SignatureRouter.PUT("updateSignature", v1.UpdateSignature)					// 更新签章
		SignatureRouter.POST("deleteSignature", v1.DeleteSignature)					// 删除签章
		SignatureRouter.POST("getSignatureRecordList", v1.GetSignatureRecordList)	// 分页获取签章记录列表
		SignatureRouter.POST("validateSignature", v1.ValidateSignature)				// 验证签章
		SignatureRouter.POST("getSignatureRecordById", v1.GetSignatureRecordById)	// 根据id获取签章记录
		SignatureRouter.POST("useSignature", v1.UseSignature)						// 使用签章
		SignatureRouter.POST("saveSignaturePosition", v1.SaveSignaturePosition)		// 保存签章位置成功
		SignatureRouter.POST("cancelSignature", v1.CancelSignature)					// 取消签章
	}
}