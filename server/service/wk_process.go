package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)

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
		txErr = tx.Select("StartPoint", "EndPoint").Delete(&edges).Error
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
