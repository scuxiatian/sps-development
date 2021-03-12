package request

import "sps-template-server/model"

type WorkflowProcessSearch struct {
	model.WorkflowProcess
	PageInfo
}
