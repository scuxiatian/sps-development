import request from '@/utils/request'
import { IdParams, PageInfoParams } from './model/common'
import { ChangeSignaturePasswordParams, SignatureParams, SignatureRecordParams, UseSignatureParams, ValidateSignaturePasswordParams } from './model/signature'

// @Tags SysSignature
// @Summary 创建签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysSignature true "签章名, 签章密码, 签章图片地址, 签章描述, 是否为公章, 所属人(仅私章)"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /signature/createSignature [post]
export const createSignature = (data: SignatureParams) => {
  return request({
    url: '/signature/createSignature',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 分页获取签章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取签章列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /signature/getSignatureList [post]
export const getSignatureList = (data: PageInfoParams) => {
  return request({
    url: '/signature/getSignatureList',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 根据id获取签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取签章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /signature/getSignatureById [post]
export const getSignatureById = (data: IdParams) => {
  return request({
    url: '/signature/getSignatureById',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 修改签章密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangeSignaturePasswordStruct true "签章id, 原密码, 新密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /signature/changePassword [put]
export const changeSignaturePassword = (data: ChangeSignaturePasswordParams) => {
  return request({
    url: '/signature/changePassword',
    method: 'put',
    data
  })
}

// @Tags SysSignature
// @Summary 更新签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysSignature true "ID, 签章名, 签章图片地址, 签章描述"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /signature/updateSignature [put]
export const updateSignature = (data: SignatureParams) => {
  return request({
    url: '/signature/updateSignature',
    method: 'put',
    data
  })
}

// @Tags SysSignature
// @Summary 删除签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /signature/deleteSignature [post]
export const deleteSignature = (data: IdParams) => {
  return request({
    url: '/signature/deleteSignature',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 分页获取签章记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchSignatureParams true "分页获取签章记录列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /signature/getSignatureRecordList [post]
export const getSignatureRecordList = (data: PageInfoParams) => {
  return request({
    url: '/signature/getSignatureRecordList',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 验证签章信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ValidateSignatureStruct true "验证签章信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /signature/validateSignature [post]
export const validateSignature = (data: ValidateSignaturePasswordParams) => {
  return request({
    url: '/signature/validateSignature',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 使用签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UseSignatureStruct true "使用签章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /signature/useSignature [post]
export const useSignature = (data: UseSignatureParams) => {
  return request({
    url: '/signature/useSignature',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 根据id获取签章记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取签章记录"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /signature/getSignatureRecordById [post]
export const getSignatureRecordById = (data: IdParams) => {
  return request({
    url: '/signature/getSignatureRecordById',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 保存签章位置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SignatureRecord true "保存签章位置"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /signature/saveSignaturePosition [post]
export const saveSignaturePosition = (data: SignatureRecordParams) => {
  return request({
    url: '/signature/saveSignaturePosition',
    method: 'post',
    data
  })
}

// @Tags SysSignature
// @Summary 取消签章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SignatureUse true "取消签章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证成功"}"
// @Router /signature/cancelSignature [post]
export const CancelSignature = (data: UseSignatureParams) => {
  return request({
    url: '/signature/cancelSignature',
    method: 'post',
    data
  })
}
