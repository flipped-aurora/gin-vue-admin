package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
)

func getTable(businessType string) interface{} {
	return model.WorkflowBusinessTable[businessType]()
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateWorkflowProcess
//@description: 创建工作流相关信息
//@param: workflowProcess model.WorkflowProcess
//@return: err error

func CreateWorkflowProcess(workflowProcess model.WorkflowProcess) (err error) {
	err = global.GVA_DB.Create(&workflowProcess).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteWorkflowProcess
//@description: 删除工作流相关信息
//@param: workflowProcess model.WorkflowProcess
//@return: err error

func DeleteWorkflowProcess(workflowProcess model.WorkflowProcess) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var txErr error
		txErr = tx.Delete(workflowProcess).Error
		if txErr != nil {
			return txErr
		}
		var edges []model.WorkflowEdge
		txErr = tx.Delete(model.WorkflowNode{}, "workflow_process_id = ?", workflowProcess.ID).Error
		if txErr != nil {
			return txErr
		}
		txErr = tx.Find(&edges, "workflow_process_id = ?", workflowProcess.ID).Error
		if txErr != nil {
			return txErr
		}
		if len(edges) > 0 {
			txErr = tx.Select("StartPoint", "EndPoint").Delete(&edges).Error
		}
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateWorkflowProcess
//@description: 批量删除工作流信息（暂未启用）
//@param: ids request.IdsReq
//@return: err error

func DeleteWorkflowProcessByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.WorkflowProcess{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateWorkflowProcess
//@description: 更新工作流相关信息
//@param: workflowProcess *model.WorkflowProcess
//@return: err error

func UpdateWorkflowProcess(workflowProcess *model.WorkflowProcess) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
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
		txErr = tx.Create(&workflowProcess).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetWorkflowProcess
//@description: 获取工作流相关信息
//@param: id string
//@return: err error,workflowProcess model.WorkflowProcess

func GetWorkflowProcess(id string) (err error, workflowProcess model.WorkflowProcess) {
	err = global.GVA_DB.Preload("Nodes").Preload("Edges").Where("id = ?", id).First(&workflowProcess).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetWorkflowCreateStep
//@description: 获取工作流步骤信息
//@param: id string
//@return: err error, workflowNodes []model.WorkflowNode

func FindWorkflowStep(id string) (err error, workflowNode model.WorkflowProcess) {
	err = global.GVA_DB.Preload("Nodes", "clazz = ?", model.START).Where("id = ?", id).First(&workflowNode).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetWorkflowProcessInfoList
//@description: 获取工作流列表
//@param: info request.WorkflowProcessSearch
//@return: err error, list interface{}, total int64

func GetWorkflowProcessInfoList(info request.WorkflowProcessSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.WorkflowProcess{})
	var workflowProcesss []model.WorkflowProcess
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Label != "" {
		db = db.Where("`label` LIKE ?", "%"+info.Label+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&workflowProcesss).Error
	return err, workflowProcesss, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StartWorkflow
//@description: 开启一个工作流
//@param: wfInterface model.GVA_Workflow
//@return: err error

func StartWorkflow(wfInterface model.GVA_Workflow) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
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

func CompleteWorkflowMove(wfInterface model.GVA_Workflow) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
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

func complete(tx *gorm.DB, wfm *model.WorkflowMove) (err error) {
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

	if nodeInfo.Clazz == model.START || nodeInfo.Clazz == model.USER_TASK {
		txErr = tx.Find(&Edges, "workflow_process_id = ? and source = ?", wfm.WorkflowProcessID, wfm.WorkflowNodeID).Error
		if txErr != nil {
			return txErr
		}
		if len(Edges) == 0 {
			return errors.New("不存在当前节点为起点的后续流程")
		}
		if len(Edges) == 1 {
			txErr = tx.Model(&returnWfm).Update("param", wfm.Param).Update("is_active", false).Update("action", wfm.Action).Update("operator_id", wfm.OperatorID).Error

			newWfm := createNewWorkflowMove(&returnWfm, Edges[0].Target)
			txErr = tx.Create(newWfm).Error
			if txErr != nil {
				return txErr
			}
			//	当target为自动节点时候 需要做一些事情 这里暂时先不处理 后续慢慢完善
		}
		if len(Edges) > 1 {
			var needUseTargetNodeID string
			if txErr != nil {
				return txErr
			}
			txErr = tx.Model(&returnWfm).Update("param", wfm.Param).Update("is_active", false).Update("action", wfm.Action).Update("operator_id", wfm.OperatorID).Error

			for _, v := range Edges {
				if v.ConditionExpression == wfm.Param {
					needUseTargetNodeID = v.Target
					break
				}
			}
			if needUseTargetNodeID == "" {
				return errors.New("未发现流转参数，流转失败")
			}
			newWfm := createNewWorkflowMove(&returnWfm, needUseTargetNodeID)
			txErr = tx.Create(newWfm).Error
			if txErr != nil {
				return txErr
			}
			//	当target为自动节点时候 需要做一些事情 这里暂时先不处理 后续慢慢完善
		}
	} else if nodeInfo.Clazz == model.EXCLUSIVE_GATEWAY {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	} else if nodeInfo.Clazz == model.INCLUSIVE_GATEWAY {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	} else if nodeInfo.Clazz == model.PARELLEL_GATEWAY {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	} else {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	}
	return nil
}

func createNewWorkflowMove(oldWfm *model.WorkflowMove, targetNodeID string) (newWfm *model.WorkflowMove) {

	return &model.WorkflowMove{
		BusinessID:        oldWfm.BusinessID,
		BusinessType:      oldWfm.BusinessType,
		PromoterID:        oldWfm.PromoterID,
		OperatorID:        0,
		WorkflowNodeID:    targetNodeID,
		WorkflowProcessID: oldWfm.WorkflowProcessID,
		Param:             "",
		Action:            "",
		IsActive:          true,
	}
}

func GetMyStated(userID uint) (err error, wfms []model.WorkflowMove) {
	err = global.GVA_DB.Preload("Promoter").Preload("Operator").Preload("WorkflowNode").Preload("WorkflowProcess").Joins("INNER JOIN workflow_nodes as node ON workflow_moves.workflow_node_id = node.id").Find(&wfms, "promoter_id = ? and ( is_active = ? OR node.clazz = ?)", userID, true, "end").Error
	return err, wfms
}

func GetMyNeed(userID uint, AuthorityID string) (err error, wfms []model.WorkflowMove) {
	user := "%," + strconv.Itoa(int(userID)) + ",%"
	auth := "%," + AuthorityID + ",%"
	err = global.GVA_DB.Preload("Promoter").Preload("Operator").Preload("WorkflowNode").Preload("WorkflowProcess").Joins("INNER JOIN workflow_nodes as node ON workflow_moves.workflow_node_id = node.id").Where("is_active = ? AND (node.assign_type = ? AND node.assign_value LIKE ? ) OR (node.assign_type = ? AND node.assign_value LIKE ? )", true, "user", user, "authority", auth).Find(&wfms).Error
	return err, wfms
}

func GetWorkflowMoveByID(id float64) (err error, move model.WorkflowMove, moves []model.WorkflowMove, business interface{}) {
	var result interface{}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
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
		fmt.Println(result)
		txErr = tx.First(result, "id = ?", move.BusinessID).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err, move, moves, result
}
