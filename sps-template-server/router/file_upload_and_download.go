package router

import (
	"github.com/gin-gonic/gin"
	"sps-template-server/api/v1"
)

func InitFileUploadAndDownloadRouter(Router *gin.RouterGroup)  {
	FileUploadAndDownloadGroup := Router.Group("fileUploadAndDownload")
	{
		FileUploadAndDownloadGroup.POST("/upload", v1.UploadFile)                                 // 上传文件
		FileUploadAndDownloadGroup.POST("/getFileList", v1.GetFileList)                           // 获取上传文件列表
		FileUploadAndDownloadGroup.POST("/deleteFile", v1.DeleteFile)                             // 删除指定文件
	}
}
