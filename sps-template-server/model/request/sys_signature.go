package request

type ChangeSignaturePasswordStruct struct {
	Id			float64	`json:"id"`
	Password	string	`json:"password"`
	NewPassword string	`json:"newPassword"`
}
