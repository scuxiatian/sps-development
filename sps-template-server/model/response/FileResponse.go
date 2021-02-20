package response

import "sps-template-server/model"

type FileResponse struct {
	File model.FileUploadAndDownload `json:"file"`
}
