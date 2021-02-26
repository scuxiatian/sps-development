import request from '@/utils/request'
import { PageInfoParams, IdParams } from './model/common'
import { ApiParams } from './model/api'

// @Tags api
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const getApiList = (data: PageInfoParams) => {
  return request({
    url: '/api/getApiList',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
export const createApi = (data: ApiParams) => {
  return request({
    url: '/api/createApi',
    method: 'post',
    data
  })
}

// @Tags menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.GetById true "根据id获取菜单"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getApiById [post]
export const getApiById = (data: IdParams) => {
  return request({
    url: '/api/getApiById',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary 更新api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "更新api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/updateApi [post]
export const updateApi = (data: ApiParams) => {
  return request({
    url: '/api/updateApi',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary 更新api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "更新api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /api/setAuthApi [post]
export const setAuthApi = (data: ApiParams) => {
  return request({
    url: '/api/setAuthApi',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
export const getAllApis = (data: any) => {
  return request({
    url: '/api/getAllApis',
    method: 'post',
    data
  })
}

// @Tags Api
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Api true "删除api"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
export const deleteApi = (data: ApiParams) => {
  return request({
    url: '/api/deleteApi',
    method: 'post',
    data
  })
}
