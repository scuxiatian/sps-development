package upload

import "mime/multipart"

//@interface_name: OSS
//@description: OSS接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOSS() OSS {
	return &Local{}
}