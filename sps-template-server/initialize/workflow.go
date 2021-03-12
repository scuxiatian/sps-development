package initialize

import (
	"sps-template-server/model"
	"sps-template-server/model/workflow"
)

func initWorkflowModel()  {
	model.WorkflowBusinessStruct = make(map[string]func() model.SdWorkflow)
	model.WorkflowBusinessStruct["leave"] = func() model.SdWorkflow {
		return new(workflow.ExaWfLeaveWorkflow)
	}
}

func initWorkflowTable()  {
	model.WorkflowBusinessTable = make(map[string]func() interface{})
	model.WorkflowBusinessTable["leave"] = func() interface{} {
		return new(workflow.ExaWfLeave)
	}
}

func InitWorkflowMode()  {
	initWorkflowModel()
	initWorkflowTable()
}