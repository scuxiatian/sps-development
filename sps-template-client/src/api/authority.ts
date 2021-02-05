import request from '@/utils/request'
import { PageInfoParams } from './model/common'
import { AuthorityParams } from './model/authority'

// @Router /authority/getAuthorityList [post]
export const getAuthorityList = (data: PageInfoParams) => {
  return request({
    url: '/authority/getAuthorityList',
    method: 'post',
    data
  })
}

// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "创建角色"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/createAuthority [post]
export const createAuthority = (data: AuthorityParams) => {
  return request({
    url: '/authority/createAuthority',
    method: 'post',
    data
  })
}

// @Summary 修改角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "修改角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
export const updateAuthority = (data: AuthorityParams) => {
  return request({
    url: '/authority/updateAuthority',
    method: 'put',
    data
  })
}

// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body {authorityId uint} true "删除角色"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/deleteAuthority [post]
export const deleteAuthority = (data: AuthorityParams) => {
  return request({
    url: '/authority/deleteAuthority',
    method: 'post',
    data
  })
}
