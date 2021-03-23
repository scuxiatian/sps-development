package request

type ChangeSignaturePasswordStruct struct {
	Id			float64	`json:"id"`
	Password	string	`json:"password"`
	NewPassword string	`json:"newPassword"`
}

type ValidateSignatureStruct struct {
	Id			float64	`json:"id"`
	Password	string	`json:"password"`
}

type UseSignatureStruct struct {
	RecordId	float64	`json:"recordId"`
	SignatureId	float64	`json:"signatureId"`
	Description string	`json:"description"`
}

type SearchSignatureParams struct {
	GetById
	PageInfo
}
