package workflow

import (
	"sps-template-server/global"
	"sps-template-server/model"
	"time"
)

// 如果含有time.Time 请自行import time包
type ExaWfLeave struct {
	global.SdModel
	Cause     string    `json:"cause" form:"cause" gorm:"column:cause;comment:"`
	StartTime time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:"`
	EndTime   time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:"`
}

type ExaWfLeaveWorkflow struct {
	// 工作流操作结构体
	model.WorkflowBase `json:"wf"`
	ExaWfLeave   `json:"business"`
}

func (ExaWfLeave) TableName() string {
	return "exa_wf_leaves"
}