package router

import (
	"github.com/gin-gonic/gin"
	v1 "sps-template-server/api/v1"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("authority")
	{
		AuthorityRouter.POST("createAuthority", v1.CreateAuthority)		// 创建角色
		AuthorityRouter.PUT("updateAuthority", v1.UpdateAuthority)		// 更新角色
		AuthorityRouter.POST("copyAuthority", v1.CopyAuthority)			// 复制角色
		AuthorityRouter.POST("deleteAuthority", v1.DeleteAuthority)		// 删除角色
		AuthorityRouter.POST("getAuthorityList", v1.GetAuthorityList)	// 获取角色列表
	}
}