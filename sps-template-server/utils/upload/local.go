package upload

import (
	"errors"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"sps-template-server/global"
	"sps-template-server/utils"
	"strings"
	"time"
)

type Local struct{}

//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.SdConfig.Local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.SdLog.Error("创建上传文件夹失败", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("创建上传文件夹失败, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.SdConfig.Local.Path + "/" + filename

	f, openError := file.Open()
	if openError != nil {
		global.SdLog.Error("读取文件失败", zap.Any("err", openError.Error()))
		return "", "", errors.New("读取文件失败, err:" + openError.Error())
	}
	defer f.Close()

	out, createErr := os.Create(p)
	if createErr != nil {
		global.SdLog.Error("创建文件失败", zap.Any("err", createErr.Error()))

		return "", "", errors.New("创建文件失败, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.SdLog.Error("传输（拷贝）文件失败", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("传输（拷贝）文件失败, err:" + copyErr.Error())
	}
	return p, filename, nil
}

//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	p := global.SdConfig.Local.Path + "/" + key
	if strings.Contains(p, global.SdConfig.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}