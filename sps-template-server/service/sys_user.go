package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/utils"
)

//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

func Register(u model.SysUser) (err error, userInter model.SysUser)  {
	var user model.SysUser
	if !errors.Is(global.SdDB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), userInter
	}
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid.NewV4()
	err = global.SdDB.Create(&u).Error
	return err, u
}

//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser)  {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.SdDB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func FindUserByUUid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err := global.SdDB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
