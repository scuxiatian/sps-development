package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sps-template-server/global"
	"sps-template-server/model/response"
	"sps-template-server/service"
)

// @Tags AuthorityMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context)  {
	if err, menus := service.GetMenuTree(getUserAuthorityId(c)); err != nil {
		global.SdLog.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}
