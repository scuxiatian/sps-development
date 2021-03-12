package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/model/response"
	"sps-template-server/service"
	"sps-template-server/utils"
)

// @Tags SysSignature
// @Summary 创建签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysSignature true "签章名, 签章密码, 签章图片地址, 签章描述, 是否为公章, 所属人(仅私章)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /api/createSignature [post]
func CreateSignature(c *gin.Context)  {
	var signature model.SysSignature
	_ = c.ShouldBindJSON(&signature)
	if err := utils.Verify(signature, utils.SignatureVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.CreateSignature(signature); err != nil {
		global.SdLog.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
