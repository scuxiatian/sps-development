// 用户
export interface UserParams {
  ID?: string;
  uuid?: string;
  userName: string;
  password: string;
  nickName: string;
  headerImg?: string;
  authorityId: string;
}

// 用户登录请求接口
export interface LoginParams {
  userName: string;
  password: string;
  captcha: string;
  captchaId: string;
}

// 修改密码
export interface ChangePasswordParams {
  userName: string;
  password: string;
  newPassword: string;
}

// 修改权限
export interface SetUserAuthParams {
  uuid: string;
  authorityId: string;
}
