package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/model/request"
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

//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = global.SdDB.Where("id = ?", id).First(&user).Error
	return err, &u
}

//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetUserInfoList(info request.PageInfo, user model.SysUser) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.SdDB.Model(&model.SysUser{})
	var userList []model.SysUser

	if user.Username != "" {
		db = db.Where("username LIKE ?", "%" + user.Username + "%")
	}

	if user.NickName != "" {
		db = db.Where("nick_name LIKE ?", "%" + user.NickName + "%")
	}

	if user.AuthorityId != "" {
		db = db.Where("authority_id = ?", user.AuthorityId)
	}

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.SdDB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func DeleteUser(id float64) (err error) {
	var user model.SysUser
	err = global.SdDB.Where("id = ?", id).Delete(&user).Error
	return err
}

//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func SetUserInfo(reqUser model.SysUser) (err error, user model.SysUser) {
	err = global.SdDB.Updates(&reqUser).Error
	return err, reqUser
}
