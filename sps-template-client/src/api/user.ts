import request from '@/utils/request'

// 用户登录请求接口
export interface UserLogin {
  username: string;
  password: string;
  captcha: string;
  captchaId: string;
}

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data: UserLogin) => {
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
