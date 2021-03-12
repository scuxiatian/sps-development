package router

import (
	"github.com/gin-gonic/gin"
	v1 "sps-template-server/api/v1"
)

func InitWorkflowProcessRouter(Router *gin.RouterGroup) {
	WorkflowProcessRouter := Router.Group("workflowProcess")
	{
		WorkflowProcessRouter.POST("createWorkflowProcess", v1.CreateWorkflowProcess)		// 新建WorkflowProcess
		WorkflowProcessRouter.DELETE("deleteWorkflowProcess", v1.DeleteWorkflowProcess)     // 删除WorkflowProcess
		WorkflowProcessRouter.PUT("updateWorkflowProcess", v1.UpdateWorkflowProcess)    	// 更新WorkflowProcess
		WorkflowProcessRouter.GET("getWorkflowProcessList", v1.GetWorkflowProcessList)		// 获取WorkflowProcess列表
		WorkflowProcessRouter.GET("findWorkflowProcess", v1.FindWorkflowProcess)			// 根据ID获取WorkflowProcess
		WorkflowProcessRouter.GET("findWorkflowStep", v1.FindWorkflowStep)                  // 根据ID获取工作流步骤
		WorkflowProcessRouter.POST("startWorkflow", v1.StartWorkflow)                       // 开启工作流
		WorkflowProcessRouter.POST("completeWorkflowMove", v1.CompleteWorkflowMove)         // 提交工作流
		WorkflowProcessRouter.GET("getMyStated", v1.GetMyStated)                            // 获取我发起的工作流
		WorkflowProcessRouter.GET("getMyNeed", v1.GetMyNeed)                                // 获取我的待办
		WorkflowProcessRouter.GET("getWorkflowMoveByID", v1.GetWorkflowMoveByID)            // 根据ID获取工作流全周期
	}
}