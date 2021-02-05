import request from '@/utils/request'

import { LoginParams } from './model/user'

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data: LoginParams) => {
  return request({
    url: '/base/login',
    method: 'post',
    data: data
  })
}

// @Summary 获取验证码
// @Produce application/json
// @Router /base/captcha [post]
export const captcha = () => {
  return request({
    url: '/base/captcha',
    method: 'post'
  })
}
