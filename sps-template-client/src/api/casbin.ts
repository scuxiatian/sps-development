import request from '@/utils/request'
import type { CasbinInReceiveParams } from './model/casbin'
import type { AuthorityIdParams } from './model/common'

// @Tags casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "获取权限列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
export const getPolicyPathByAuthorityId = (data: AuthorityIdParams) => {
  return request({
    url: '/casbin/getPolicyPathByAuthorityId',
    method: 'post',
    data
  })
}

// @Tags authority
// @Summary 更改角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateAuthorityPatams true "更改角色api权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/UpdateCasbin [post]

export const updateCasbin = (data: CasbinInReceiveParams) => {
  return request({
    url: '/casbin/updateCasbin',
    method: 'post',
    data
  })
}
