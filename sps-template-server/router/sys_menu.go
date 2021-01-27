package router

import (
	"github.com/gin-gonic/gin"
	"sps-template-server/api/v1"
)

func InitMenuRouter(Router *gin.RouterGroup)  {
	MenuRouter := Router.Group("menu")
	{
		MenuRouter.POST("getMenu", v1.GetMenu)					// 获取菜单树
		MenuRouter.POST("getMenuList", v1.GetMenuList)			// 分页获取基础menu列表
		MenuRouter.POST("getBaseMenuById", v1.GetBaseMenuById)	// 根据id获取菜单
		MenuRouter.POST("addBaseMenu", v1.AddBaseMenu)			// 新增菜单
		MenuRouter.POST("updateBaseMenu", v1.UpdateBaseMenu)	// 更新菜单
		MenuRouter.POST("deleteBaseMenu", v1.DeleteBaseMenu)	// 删除菜单
	}
}
