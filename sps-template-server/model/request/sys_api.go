package request

import "sps-template-server/model"

type SearchApiParams struct {
	model.SysApi
	PageInfo
	OrderKey string `json:"orderKey"`
	Desc     bool   `json:"desc"`
}
