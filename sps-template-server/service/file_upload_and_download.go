package service

import (
	"errors"
	"mime/multipart"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/model/request"
	"sps-template-server/utils/upload"
	"strings"
)

//@function: Upload
//@description: 创建文件上传记录
//@param: file model.FileUploadAndDownload
//@return: error

func Upload(file model.FileUploadAndDownload) error {
	return global.SdDB.Create(&file).Error
}

//@function: FindFile
//@description: 查找文件
//@param: id uint
//@return: error, model.FileUploadAndDownload

func FindFile(id uint) (error, model.FileUploadAndDownload) {
	var file model.FileUploadAndDownload
	err := global.SdDB.Where("id = ?", id).First(&file).Error
	return err, file
}

//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.FileUploadAndDownload
//@return: err error

func DeleteFile(file model.FileUploadAndDownload) (err error) {
	var fileFromDb model.FileUploadAndDownload
	err, fileFromDb = FindFile(file.ID)
	oss := upload.NewOSS()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.SdDB.Where("id = ?", file.ID).Unscoped().Delete(file).Error
	return err
}

//@function: GetFileRecordInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetFileRecordInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.SdDB
	var fileLists []model.FileUploadAndDownload
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return err, fileLists, total
}

//@function: UploadFile
//@description: 文件上传到本地
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.FileUploadAndDownload

func UploadFile(header *multipart.FileHeader, noSave string) (err error, file model.FileUploadAndDownload) {
	oss := upload.NewOSS()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := model.FileUploadAndDownload{
			Name:    header.Filename,
			Url:     filePath,
			Tag:     s[len(s) - 1],
			Key:     key,
		}
		return Upload(f), f
	}
	return
}
