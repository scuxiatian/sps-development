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
//@return: err error, signature *model.SysSignature

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

//@function: GetSignatureRecordList
//@description: 分页获取数据
//@param: info request.SearchSignatureParams
//@return: err error, list interface{}, total int64

func GetSignatureRecordList(info request.SearchSignatureParams) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.SdDB.Model(&model.SignatureRecord{})
	var signatureRecordList []model.SignatureRecord
	if info.Id != 0 {
		db = db.Where("id = ?", info.Id)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("SignatureUses").Preload("SignatureUses.SignatureBase").Preload("SignatureUses.Operator").Find(&signatureRecordList).Error
	return err, signatureRecordList, total
}

//@function: ValidateSignature
//@description: 验证签章密码
//@param: id float64, password string
//@return: err error

func ValidateSignature(id float64, password string) (err error) {
	var s model.SysSignature
	if errors.Is(global.SdDB.Where("id = ? AND password = ?", id, utils.MD5V([]byte(password))).First(&s).Error, gorm.ErrRecordNotFound) {
		return errors.New("签章密码错误")
	}
	return err
}

//@function: UseSignature
//@description: 根据id获取签章记录
//@param: recordId float64
//@return: err error, record model.SignatureRecord
func GetSignatureRecordById(id float64) (err error, record model.SignatureRecord) {
	err = global.SdDB.Preload("SignatureUses").Preload("SignatureUses.SignatureBase").Where("id = ?", id).First(&record).Error
	return
}

//@function: UseSignature
//@description: 使用签章
//@param: recordId float64, signatureId float64
//@return: err error, record model.SignatureRecord, signatureUse model.SignatureUse

func UseSignature(recordId float64, signatureId float64, userId uint, description string) (err error, record model.SignatureRecord, signatureUse model.SignatureUse) {
	err = global.SdDB.Transaction(func(tx *gorm.DB) error {
		if recordId == 0 {
			err, record = CreateSignatureRecord(tx)
			if err != nil {
				return err
			}
		} else {
			if err = tx.Where("id = ?", recordId).First(&record).Error; err != nil {
				return err
			}
		}
		var signature model.SysSignature
		if err = tx.Where("id = ?", signatureId).First(&signature).Error; err != nil {
			return errors.New("签章不存在")
		}
		s := model.SignatureUse{
			SignatureId: signature.ID,
			SignatureRecordID: record.ID,
			OperatorID: userId,
			Description: description,
		}
		if err = tx.Create(&s).Error; err != nil {
			return err
		}
		err = tx.Preload("SignatureUses").Preload("SignatureUses.SignatureBase").Where("id = ?", record.ID).First(&record).Error
		return err
	})
	return err, record, signatureUse
}

//@function: SaveSignaturePosition
//@description: 保存签章位置
//@param: recordId float64, signatureId float64
//@return: err error, record model.SignatureRecord

func SaveSignaturePosition(record model.SignatureRecord) (err error) {
	return global.SdDB.Transaction(func(tx *gorm.DB) error {
		//var signatures []model.SignatureUse
		var txErr error
		//if txErr = tx.Where("signature_record_id = ?", record.ID).Find(&signatures).Error; txErr != nil {
		//	return txErr
		//}
		signatures := record.SignatureUses
		if len(signatures) == 0{
			return nil
		}
		for _, signature := range signatures {
			if txErr = tx.Save(&signature).Error; txErr != nil {
				return txErr
			}
		}
		return nil
	})
}

//@function: SaveSignaturePosition
//@description: 取消签章
//@param: recordId float64
//@return: err error

func CancelSignature(signature model.SignatureUse) (err error, record model.SignatureRecord) {
	recordId := signature.SignatureRecordID
	if err = global.SdDB.Delete(&signature).Error; err != nil {
		return
	}
	err = global.SdDB.Preload("SignatureUses").Preload("SignatureUses.SignatureBase").Where("id = ?", recordId).First(&record).Error
	return
}


//@function: CreateSignatureRecord
//@description: 创建签章记录
//@return: err error, record model.SignatureRecord

func CreateSignatureRecord(tx *gorm.DB) (err error, record model.SignatureRecord) {
	err = tx.Create(&record).Error
	return
}
