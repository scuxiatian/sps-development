package service

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sps-template-server/constant"
	"sps-template-server/global"
	"sps-template-server/model"
	"sps-template-server/model/request"
	"strconv"
)

//@function: CreateWorkflowProcess
//@description: 创建工作流相关信息
//@param: workflowProcess model.WorkflowProcess
//@return: err error

func CreateWorkflowProcess(workflowProcess model.WorkflowProcess) (err error) {
	err = global.SdDB.Create(&workflowProcess).Error
	return err
}

//@function: GetWorkflowProcessInfoList
//@description: 获取工作流列表
//@param: info request.WorkflowProcessSearch
//@return: err error, list interface{}, total int64

func GetWorkflowProcessInfoList(info request.WorkflowProcessSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.SdDB.Model(&model.WorkflowProcess{})
	var workflowProcess []model.WorkflowProcess
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%" + info.Name + "%")
	}
	if info.Label != "" {
		db = db.Where("`label` LIKE ?", "%" + info.Label + "%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&workflowProcess).Error
	return err, workflowProcess, total
}

//@function: GetWorkflowProcess
//@description: 获取工作流相关信息
//@param: id string
//@return: err error,workflowProcess model.WorkflowProcess

func GetWorkflowProcess(id string) (err error, workflowProcess model.WorkflowProcess) {
	err = global.SdDB.Preload("Nodes").Preload("Edges").Where("id = ?", id).First(&workflowProcess).Error
	return
}

//@function: UpdateWorkflowProcess
//@description: 更新工作流相关信息
//@param: workflowProcess *model.WorkflowProcess
//@return: err error

func UpdateWorkflowProcess(workflowProcess *model.WorkflowProcess) (err error) {
	return global.SdDB.Transaction(func(tx *gorm.DB) error {
		var txErr error
		var edges []model.WorkflowEdge
		var edgesIds []string
		txErr = tx.Unscoped().Delete(workflowProcess).Error
		if txErr != nil {
			return txErr
		}
		txErr = tx.Unscoped().Delete(model.WorkflowNode{}, "workflow_process_id = ?", workflowProcess.ID).Error
		if txErr != nil {
			return txErr
		}
		txErr = tx.Unscoped().Find(&edges, "workflow_process_id = ?", workflowProcess.ID).Error
		if txErr != nil {
			return txErr
		}
		if len(edges) > 0 {
			txErr = tx.Unscoped().Delete(&edges).Error
			if txErr != nil {
				return txErr
			}
			for _, v := range edges {
				edgesIds = append(edgesIds, v.ID)
			}
			txErr = tx.Unscoped().Delete(model.WorkflowStartPoint{}, "workflow_edge_id in ?", edgesIds).Error
			if txErr != nil {
				return txErr
			}
			txErr = tx.Unscoped().Delete(model.WorkflowEndPoint{}, "workflow_edge_id in ?", edgesIds).Error
			if txErr != nil {
				return txErr
			}
		}
		txErr = tx.Create(&workflowProcess).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})
}

//@function: DeleteWorkflowProcess
//@description: 删除工作流相关信息
//@param: workflowProcess model.WorkflowProcess
//@return: err error

func DeleteWorkflowProcess(workflowProcess model.WorkflowProcess) (err error) {
	err = global.SdDB.Transaction(func(tx *gorm.DB) error {
		var txErr error
		var edges []model.WorkflowEdge
		var edgesIds []string
		txErr = tx.Unscoped().Delete(workflowProcess).Error
		if txErr != nil {
			return txErr
		}
		txErr = tx.Unscoped().Delete(model.WorkflowNode{}, "workflow_process_id = ?", workflowProcess.ID).Error
		if txErr != nil {
			return txErr
		}
		txErr = tx.Unscoped().Find(&edges, "workflow_process_id = ?", workflowProcess.ID).Error
		if txErr != nil {
			return txErr
		}
		if len(edges) > 0 {
			txErr = tx.Unscoped().Delete(&edges).Error
			if txErr != nil {
				return txErr
			}
			for _, v := range edges {
				edgesIds = append(edgesIds, v.ID)
			}
			txErr = tx.Unscoped().Delete(model.WorkflowStartPoint{}, "workflow_edge_id in ?", edgesIds).Error
			if txErr != nil {
				return txErr
			}
			txErr = tx.Unscoped().Delete(model.WorkflowEndPoint{}, "workflow_edge_id in ?", edgesIds).Error
			if txErr != nil {
				return txErr
			}
		}
		return nil
	})
	return err
}

//@function: GetWorkflowCreateStep
//@description: 获取工作流步骤信息
//@param: id string
//@return: err error, workflowNodes []model.WorkflowNode

func FindWorkflowStep(id string) (err error, workflowNode model.WorkflowProcess) {
	err = global.SdDB.Preload("Nodes", "clazz = ?", constant.START).Where("id = ?", id).First(&workflowNode).Error
	return
}

//@function: StartWorkflow
//@description: 开启一个工作流
//@param: wfInterface model.SdWorkflow
//@return: err error

func StartWorkflow(wfInterface model.SdWorkflow) (err error) {
	err = global.SdDB.Transaction(func(tx *gorm.DB) error {
		var txErr error
		tableName := getTable(wfInterface.GetBusinessType()).(schema.Tabler).TableName()
		txErr = tx.Table(tableName).Create(wfInterface).Error
		if txErr != nil {
			return txErr
		}
		wfm := wfInterface.CreateWorkflowMove()
		txErr = tx.Create(wfm).Error
		if txErr != nil {
			return txErr
		}
		txErr = complete(tx, wfm)
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err
}

//@function: CompleteWorkflowMove
//@description: 执行工作流
//@param: wfInterface model.SdWorkflow
//@return: err error

func CompleteWorkflowMove(wfInterface model.SdWorkflow) (err error) {
	err = global.SdDB.Transaction(func(tx *gorm.DB) error {
		var txErr error
		tableName := getTable(wfInterface.GetBusinessType()).(schema.Tabler).TableName()
		txErr = tx.Table(tableName).Where("id = ?", wfInterface.GetBusinessID()).Updates(wfInterface).Error
		if txErr != nil {
			return txErr
		}
		nowWorkflowMove := wfInterface.CreateWorkflowMove()
		txErr = complete(tx, nowWorkflowMove)
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err
}

//@function: GetMyStated
//@description: 获取用户发起的流程列表
//@param: userID uint
//@return: err error, wfms []model.WorkflowMove

func GetMyStated(userID uint) (err error, wfms []model.WorkflowMove) {
	err = global.SdDB.Preload("Promoter").Preload("Operator").Preload("WorkflowNode").Preload("WorkflowProcess").Joins("INNER JOIN workflow_nodes as node ON workflow_moves.workflow_node_id = node.id").Find(&wfms, "promoter_id = ? and ( is_active = ? OR node.clazz = ?)", userID, true, "end").Error
	return
}

//@function: GetMyNeed
//@description: 获取用户待办的流程列表
//@param: userID uint, AuthorityID string
//@return: err error, wfms []model.WorkflowMove

func GetMyNeed(userID uint, AuthorityID string) (err error, wfms []model.WorkflowMove) {
	user := "%," + strconv.Itoa(int(userID)) + ",%"
	auth := "%," + AuthorityID + ",%"
	err = global.SdDB.Preload("Promoter").Preload("Operator").Preload("WorkflowNode").Preload("WorkflowProcess").Joins("INNER JOIN workflow_nodes as node ON workflow_moves.workflow_node_id = node.id").Where("is_active = ? AND ((node.assign_type = ? AND node.assign_value LIKE ? ) OR (node.assign_type = ? AND node.assign_value LIKE ? ) OR (node.assign_type = ? AND promoter_id = ? ))", true, "user", user, "authority", auth, "self", userID).Find(&wfms).Error
	return err, wfms
}

//@function: GetWorkflowMoveByID
//@description: 获取工作流步骤信息
//@param: id float64
//@return: err error, move model.WorkflowMove, moves []model.WorkflowMove, business interface{}

func GetWorkflowMoveByID(id float64) (err error, move model.WorkflowMove, moves []model.WorkflowMove, business interface{}) {
	var result interface{}
	err = global.SdDB.Transaction(func(tx *gorm.DB) error {
		var txErr error
		txErr = tx.Preload("Promoter").Preload("Operator").Preload("WorkflowNode").Preload("WorkflowProcess").First(&move, "id = ?", id).Error
		if txErr != nil {
			return txErr
		}
		txErr = tx.Preload("Promoter").Preload("Operator").Preload("WorkflowNode").Preload("WorkflowProcess").Find(&moves, "business_id = ? AND business_type = ?", move.BusinessID, move.BusinessType).Error
		if txErr != nil {
			return txErr
		}
		result = getTable(move.BusinessType)
		//fmt.Println(result)
		txErr = tx.First(result, "id = ?", move.BusinessID).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err, move, moves, result
}

func getTable(businessType string) interface{} {
	return model.WorkflowBusinessTable[businessType]()
}

func complete(tx *gorm.DB, wfm *model.WorkflowMove) (err error)  {
	var returnWfm model.WorkflowMove
	var nodeInfo model.WorkflowNode
	var Edges []model.WorkflowEdge
	txErr := tx.First(&returnWfm, "id = ? AND is_active = ?", wfm.ID, true).Error
	if txErr != nil {
		return txErr
	}
	txErr = tx.First(&nodeInfo, "id = ?", wfm.WorkflowNodeID).Error
	if txErr != nil {
		return txErr
	}
	
	if nodeInfo.Clazz == constant.START || nodeInfo.Clazz == constant.USER_TASK {
		txErr = tx.Find(&Edges, "workflow_process_id = ? and source = ?", wfm.WorkflowProcessID, wfm.WorkflowNodeID).Error
		if txErr != nil {
			return txErr
		}
		if len(Edges) == 0 {
			return errors.New("不存在当前节点为起点的后续流程")
		}
		if len(Edges) == 1 {
			txErr = tx.Model(&returnWfm).Update("param", wfm.Param).Update("is_active", false).Update("action", wfm.Action).Update("operator_id", wfm.OperatorID).Error
			if txErr != nil {
				return txErr
			}
			txErr, newWfm := createNewWorkflowMove(tx, &returnWfm, Edges[0].Target)
			if txErr != nil {
				return txErr
			}
			if len(newWfm) > 0 {
				txErr = tx.Create(&newWfm).Error
				if txErr != nil {
					return txErr
				}
			}
		}
		if len(Edges) > 1 {
			var needUseTargetNodeID string

			txErr = tx.Model(&returnWfm).Update("param", wfm.Param).Update("is_active", false).Update("action", wfm.Action).Update("operator_id", wfm.OperatorID).Error
			if txErr != nil {
				return txErr
			}
			for _, v := range Edges {
				if v.ConditionExpression == wfm.Param {
					needUseTargetNodeID = v.Target
					break
				}
			}
			if needUseTargetNodeID == "" {
				return errors.New("未发现流转参数，流转失败")
			}
			txErr, newWfm := createNewWorkflowMove(tx, &returnWfm, needUseTargetNodeID)
			if txErr != nil {
				return txErr
			}
			if len(newWfm) > 0 {
				txErr = tx.Create(&newWfm).Error
				if txErr != nil {
					return txErr
				}
			}
		}
	} else {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	}
	return nil
}

func createNewWorkflowMove(tx *gorm.DB, oldWfm *model.WorkflowMove, targetNodeID string) (err error, newWfm []model.WorkflowMove) {
	// 以下所有非 default的节点的下一步流转均应该处理为递归形式
	var nodeInfo model.WorkflowNode
	//var edge model.WorkflowEdge
	//var edges []model.WorkflowEdge
	//var wfms []model.WorkflowMove
	txErr := tx.First(&nodeInfo, "id = ?", targetNodeID).Error
	if txErr != nil {
		return txErr, []model.WorkflowMove{}
	}
	switch nodeInfo.Clazz {
	case constant.END:
		newWfm = append(newWfm, model.WorkflowMove{
			BusinessID:        oldWfm.BusinessID,
			BusinessType:      oldWfm.BusinessType,
			PromoterID:        oldWfm.PromoterID,
			OperatorID:        oldWfm.OperatorID,
			WorkflowProcessID: oldWfm.WorkflowProcessID,
			WorkflowNodeID:    targetNodeID,
			Param:             "",
			Action:            "",
			IsActive:          false,
		})
		return nil, newWfm
	default:
		newWfm = append(newWfm, model.WorkflowMove{
			BusinessID:        oldWfm.BusinessID,
			BusinessType:      oldWfm.BusinessType,
			PromoterID:        oldWfm.PromoterID,
			OperatorID:        0,
			WorkflowProcessID: oldWfm.WorkflowProcessID,
			WorkflowNodeID:    targetNodeID,
			Param:             "",
			Action:            "",
			IsActive:          true,
		})
		return nil, newWfm
	}
}
