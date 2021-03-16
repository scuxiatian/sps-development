package service

import (
	"errors"
	"gorm.io/gorm"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/model/request"
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

//@function: GetSignatureList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetSignatureList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.SdDB.Model(&model.SysSignature{})
	var signatureList []model.SysSignature
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Owner").Find(&signatureList).Error
	return err, signatureList, total
}

//@function: FindSignatureById
//@description: 通过id获取签章信息
//@param: id float64
//@return: err error, user *model.SysSignature

func FindSignatureById(id float64) (err error, signature *model.SysSignature) {
	var s model.SysSignature
	err = global.SdDB.Where("id = ?", id).First(&s).Error
	return err, &s
}

//@function: ChangeSignaturePassword
//@description: 修改签章密码
//@param: id string, oldPassword string, newPassword string
//@return: err error

func ChangeSignaturePassword(id float64, oldPassword string, newPassword string) (err error) {
	var signature model.SysSignature
	err = global.SdDB.Where("id = ? AND password = ?", id, utils.MD5V([]byte(oldPassword))).First(&signature).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err
}

//@function: UpdateSignature
//@description: 根据id更新签章
//@param: api model.SysSignature
//@return: err error

func UpdateSignature(signature model.SysSignature) (err error) {
	var oldSignature model.SysSignature
	err = global.SdDB.Where("id = ?", signature.ID).First(&oldSignature).Error
	if oldSignature.Name != signature.Name || oldSignature.Url != signature.Url {
		if !errors.Is(global.SdDB.Where("name = ? OR url = ? AND id != ?", signature.Name, signature.Url, signature.ID).First(&model.SysSignature{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在同名签章或同图片签章")
		}
	}
	if err != nil {
		return err
	} else {
		err = global.SdDB.Save(&signature).Error
	}
	return err
}

//@function: DeleteSignature
//@description: 删除签章
//@param: id float64
//@return: err error

func DeleteSignature (id float64) (err error)  {
	var signature model.SysSignature
	err = global.SdDB.Where("id = ?", id).Delete(&signature).Error
	return err
}


