package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sps-template-server/global"
	"sps-template-server/middleware"
	"sps-template-server/model"
	"sps-template-server/model/request"
	"sps-template-server/model/response"
	"sps-template-server/service"
	"sps-template-server/utils"
	"time"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context)  {
	var L request.Login
	_ = c.ShouldBindJSON(&L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	U := &model.SysUser{Username: L.Username, Password: L.Password}
	if err, user := service.Login(U); err != nil {
		global.SdLog.Error("登录失败!用户名不存在或者密码错误", zap.Any("err", err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
	} else {
		tokenNext(c, *user)
	}
}

func tokenNext(c *gin.Context, user model.SysUser)  {
	j := &middleware.JWT{SigningKey: []byte(global.SdConfig.JWT.SigningKey)}
	claims := request.CustomClaims{
		UUID:           user.UUID,
		ID:             user.ID,
		Username:       user.Username,
		NickName:       user.NickName,
		AuthorityId:    user.AuthorityId,
		BufferTime:     global.SdConfig.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore:	time.Now().Unix() - 1000,
			ExpiresAt:	time.Now().Unix() + global.SdConfig.JWT.ExpiresTime,
			Issuer:		"spsDev",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.SdLog.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func Register(c *gin.Context)  {
	var R request.Register
	_ = c.ShouldBindJSON(&R)
	if err := utils.Verify(R, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &model.SysUser{
		Username:  R.Username,
		Password:  R.Password,
		NickName:  R.NickName,
		HeaderImg: R.HeaderImg,
	}
	err, userReturn := service.Register(*user)
	if err != nil {
		global.SdLog.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(response.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(response.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}

func getUserAuthorityId(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.SdLog.Error("从Gin的Context中获取从jwt解析出来的用户UUID失败, 请检查路由是否使用jwt中间件")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
	}
}