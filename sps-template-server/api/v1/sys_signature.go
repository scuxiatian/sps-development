package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/model/request"
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
// @Router /signature/createSignature [post]
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

// @Tags SysSignature
// @Summary 分页获取签章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取签章列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /signature/getSignatureList [post]
func GetSignatureList(c *gin.Context)  {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetSignatureList(pageInfo); err != nil {
		global.SdLog.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags SysSignature
// @Summary 根据id获取签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取签章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /signature/getSignatureById [post]
func GetSignatureById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, signature := service.FindSignatureById(idInfo.Id)
	if err != nil {
		global.SdLog.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(signature, c)
	}
}

// @Tags SysSignature
// @Summary 修改签章密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangeSignaturePasswordStruct true "签章id, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /signature/changePassword [put]
func ChangePassword(c *gin.Context) {
	var params request.ChangeSignaturePasswordStruct
	_ = c.ShouldBindJSON(&params)
	if err := utils.Verify(params, utils.ChangeSignaturePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.ChangeSignaturePassword(params.Id, params.Password, params.NewPassword); err != nil {
		global.SdLog.Error("修改失败", zap.Any("err", err))
		response.FailWithMessage("修改失败, 原密码与当前签章不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysSignature
// @Summary 更新签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysSignature true "ID, 签章名, 签章图片地址, 签章描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /signature/updateSignature [put]
func UpdateSignature(c *gin.Context)  {
	var signature model.SysSignature
	_ = c.ShouldBindJSON(&signature)
	if err := utils.Verify(signature, utils.UpdateSignature); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.UpdateSignature(signature); err != nil {
		global.SdLog.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysSignature
// @Summary 删除签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /signature/deleteSignature [post]
func DeleteSignature(c *gin.Context)  {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteSignature(idInfo.Id); err != nil {
		global.SdLog.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysSignature
// @Summary 分页获取签章记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchSignatureParams true "分页获取签章记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /signature/getSignatureRecordList [post]
func GetSignatureRecordList(c *gin.Context)  {
	var pageInfo request.SearchSignatureParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := service.GetSignatureRecordList(pageInfo); err != nil {
		global.SdLog.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags SysSignature
// @Summary 验证签章信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ValidateSignatureStruct true "验证签章信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /signature/validateSignature [post]
func ValidateSignature(c *gin.Context) {
	var info request.ValidateSignatureStruct
	_ = c.ShouldBindJSON(&info)
	if err := utils.Verify(info, utils.ValidateSignatureVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.ValidateSignature(info.Id, info.Password); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("验证成功", c)
	}
}

// @Tags SysSignature
// @Summary 根据id获取签章记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取签章记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /signature/getSignatureRecordById [post]
func GetSignatureRecordById(c *gin.Context) {
	var info request.GetById
	_ = c.ShouldBindJSON(&info)
	if err := utils.Verify(info, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, record := service.GetSignatureRecordById(info.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(record, c)
	}
}

// @Tags SysSignature
// @Summary 使用签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UseSignatureStruct true "使用签章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /signature/useSignature [post]
func UseSignature(c *gin.Context)  {
	var info request.UseSignatureStruct
	_ = c.ShouldBindJSON(&info)
	if err := utils.Verify(info, utils.UseSignatureVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, record, signature := service.UseSignature(info.RecordId, info.SignatureId, getUserID(c), info.Description)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithDetailed(response.SignatureRecordResponse{
			Record:    record,
			Signature: signature,
		}, "签章成功", c)
	}
}

// @Tags SysSignature
// @Summary 保存签章位置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SignatureRecord true "保存签章位置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /signature/saveSignaturePosition [post]
func SaveSignaturePosition(c *gin.Context)  {
	var record model.SignatureRecord
	_ = c.ShouldBindJSON(&record)
	if err := service.SaveSignaturePosition(record); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("保存签章位置成功", c)
	}
}

// @Tags SysSignature
// @Summary 取消签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SignatureUse true "取消签章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /signature/cancelSignature [post]
func CancelSignature(c *gin.Context) {
	var signature model.SignatureUse
	_ = c.ShouldBindJSON(&signature)
	if err := utils.Verify(signature, utils.CancelSignatureVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, record := service.CancelSignature(signature); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(record, c)
	}
}