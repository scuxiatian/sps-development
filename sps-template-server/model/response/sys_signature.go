package response

import "sps-template-server/model"

type SignatureRecordResponse struct {
	Record		model.SignatureRecord	`json:"record"`
	Signature	model.SignatureUse		`json:"signature"`
} 
