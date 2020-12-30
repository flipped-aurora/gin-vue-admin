package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/constant"
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
		txErr = tx.Delete(&workflowProcess).Error
		if txErr != nil {
			return txErr
		}
		var edges []model.WorkflowEdge
		txErr = tx.Delete(&model.WorkflowNode{}, "workflow_process_id = ?", workflowProcess.ID).Error
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
		txErr = tx.Unscoped().Delete(&model.WorkflowNode{}, "workflow_process_id = ?", workflowProcess.ID).Error
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
		txErr = tx.Unscoped().Delete(&model.WorkflowStartPoint{}, "workflow_edge_id in ?", edgesIds).Error
		if txErr != nil {
			return txErr
		}
		txErr = tx.Unscoped().Delete(&model.WorkflowEndPoint{}, "workflow_edge_id in ?", edgesIds).Error
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
	err = global.GVA_DB.Preload("Nodes", "clazz = ?", constant.START).Where("id = ?", id).First(&workflowNode).Error
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
			//	当target为自动节点时候 需要做一些事情 这里暂时先不处理 后续慢慢完善
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
	} else if nodeInfo.Clazz == constant.EXCLUSIVE_GATEWAY {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	} else if nodeInfo.Clazz == constant.INCLUSIVE_GATEWAY {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	} else if nodeInfo.Clazz == constant.PARELLEL_GATEWAY {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	} else {
		return errors.New("目前只支持start节点和userTask功能，其他功能正在开发中")
	}
	return nil
}

func createNewWorkflowMove(tx *gorm.DB, oldWfm *model.WorkflowMove, targetNodeID string) (err error, newWfm []model.WorkflowMove) {
	// 以下所有非 default的节点的下一步流转均应该处理为递归形式
	var nodeInfo model.WorkflowNode
	var edge model.WorkflowEdge
	var edges []model.WorkflowEdge
	var wfms []model.WorkflowMove
	txErr := tx.First(&nodeInfo, "id = ?", targetNodeID).Error
	if txErr != nil {
		return txErr, []model.WorkflowMove{}
	}
	switch nodeInfo.Clazz {
	case constant.EXCLUSIVE_GATEWAY:
		// 当为排他网关时候 选择一个参数进行排他线路选择
		txErr := tx.First(&edge, "workflow_process_id = ? and source = ? and condition_expression = ?", oldWfm.WorkflowProcessID, nodeInfo.ID, oldWfm.Param).Error
		if txErr != nil {
			return txErr, []model.WorkflowMove{}
		}
		newWfm = append(newWfm, model.WorkflowMove{
			BusinessID:        oldWfm.BusinessID,
			BusinessType:      oldWfm.BusinessType,
			PromoterID:        oldWfm.PromoterID,
			OperatorID:        0,
			WorkflowNodeID:    edge.Target,
			WorkflowProcessID: oldWfm.WorkflowProcessID,
			Param:             "",
			Action:            "",
			IsActive:          true})
		return nil, newWfm
	case constant.INCLUSIVE_GATEWAY:
		// 当为包容网关时，需要等待其他网关执行结束才进行创建
		txErr := tx.Find(&edges, "workflow_process_id = ? and target = ?", oldWfm.WorkflowProcessID, nodeInfo.ID).Error
		if txErr != nil {
			return txErr, []model.WorkflowMove{}
		}
		var sourceIds []string
		for _, v := range edges {
			sourceIds = append(sourceIds, v.Source)
		}
		txErr = tx.Find(&wfms, "workflow_process_id = ? and business_id = ? and workflow_node_id in (?) and is_active = ?", oldWfm.WorkflowProcessID, oldWfm.BusinessID, sourceIds, false).Error
		if txErr != nil {
			return txErr, []model.WorkflowMove{}
		}
		if len(wfms) != len(edges) {
			return nil, []model.WorkflowMove{}
		}
		if len(wfms) == len(edges) {
			params := make(map[string]int)
			var param string
			var temp int
			for _, v := range wfms {
				params[v.Param]++
			}
			for k, v := range params {
				if temp < v {
					temp = v
					param = k
				}
			}
			//参数携带模式暂时未定 暂时为少数服从多数原则  后续会增加原则配置 （少数服从多数，仅一关键字即为关键字，所有关键字才为关键字 三种方案）
			txErr := tx.First(&edge, "workflow_process_id = ? and source = ? and condition_expression = ?", oldWfm.WorkflowProcessID, nodeInfo.ID, param).Error
			if txErr != nil {
				return txErr, []model.WorkflowMove{}
			}
			newWfm = append(newWfm, model.WorkflowMove{
				BusinessID:        oldWfm.BusinessID,
				BusinessType:      oldWfm.BusinessType,
				PromoterID:        oldWfm.PromoterID,
				OperatorID:        0,
				WorkflowNodeID:    edge.Target,
				WorkflowProcessID: oldWfm.WorkflowProcessID,
				Param:             "",
				Action:            "",
				IsActive:          true})
		}
		return nil, newWfm

	case constant.PARELLEL_GATEWAY:
		// 当为并行网关时候 找出所有线路创建并行节点
		txErr := tx.Find(&edges, "workflow_process_id = ? and source = ?", oldWfm.WorkflowProcessID, nodeInfo.ID).Error
		if txErr != nil {
			return txErr, []model.WorkflowMove{}
		}
		for _, v := range edges {
			newWfm = append(newWfm, model.WorkflowMove{
				BusinessID:        oldWfm.BusinessID,
				BusinessType:      oldWfm.BusinessType,
				PromoterID:        oldWfm.PromoterID,
				OperatorID:        0,
				WorkflowNodeID:    v.Target,
				WorkflowProcessID: oldWfm.WorkflowProcessID,
				Param:             "",
				Action:            "",
				IsActive:          true})
		}

		return nil, newWfm
	case constant.END:
		newWfm = append(newWfm, model.WorkflowMove{
			BusinessID:        oldWfm.BusinessID,
			BusinessType:      oldWfm.BusinessType,
			PromoterID:        oldWfm.PromoterID,
			OperatorID:        oldWfm.OperatorID,
			WorkflowNodeID:    targetNodeID,
			WorkflowProcessID: oldWfm.WorkflowProcessID,
			Param:             "",
			Action:            "",
			IsActive:          false})
		return nil, newWfm
	default:
		newWfm = append(newWfm, model.WorkflowMove{
			BusinessID:        oldWfm.BusinessID,
			BusinessType:      oldWfm.BusinessType,
			PromoterID:        oldWfm.PromoterID,
			OperatorID:        0,
			WorkflowNodeID:    targetNodeID,
			WorkflowProcessID: oldWfm.WorkflowProcessID,
			Param:             "",
			Action:            "",
			IsActive:          true})
		return nil, newWfm
	}
}

func GetMyStated(userID uint) (err error, wfms []model.WorkflowMove) {
	err = global.GVA_DB.Preload("Promoter").Preload("Operator").Preload("WorkflowNode").Preload("WorkflowProcess").Joins("INNER JOIN workflow_nodes as node ON workflow_moves.workflow_node_id = node.id").Find(&wfms, "promoter_id = ? and ( is_active = ? OR node.clazz = ?)", userID, true, "end").Error
	return err, wfms
}

func GetMyNeed(userID uint, AuthorityID string) (err error, wfms []model.WorkflowMove) {
	user := "%," + strconv.Itoa(int(userID)) + ",%"
	auth := "%," + AuthorityID + ",%"
	err = global.GVA_DB.Preload("Promoter").Preload("Operator").Preload("WorkflowNode").Preload("WorkflowProcess").Joins("INNER JOIN workflow_nodes as node ON workflow_moves.workflow_node_id = node.id").Where("is_active = ? AND ((node.assign_type = ? AND node.assign_value LIKE ? ) OR (node.assign_type = ? AND node.assign_value LIKE ? ) OR (node.assign_type = ? AND promoter_id = ? ))", true, "user", user, "authority", auth, "self", userID).Find(&wfms).Error
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
