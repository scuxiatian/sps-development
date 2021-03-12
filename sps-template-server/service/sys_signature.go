package service

import (
	"errors"
	"gorm.io/gorm"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/utils"
)

//@function: CreateSignature
//@description: 新增签章
//@param: signature model.SysSignature
//@return: err error

func CreateSignature(signature model.SysSignature) (err error) {
	if !errors.Is(global.SdDB.Where("name = ? OR url = ?", signature.Name, signature.Url).First(&model.SysSignature{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在同名签章或同图片签章")
	}
	// 密码md5简单加密
	signature.Password = utils.MD5V([]byte(signature.Password))
	return global.SdDB.Create(&signature).Error
}