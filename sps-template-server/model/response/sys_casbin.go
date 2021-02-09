package response

import "sps-template-server/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}

