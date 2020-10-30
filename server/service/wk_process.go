package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gorm.io/gorm"
)

// @title    CreateWorkflowProcess
// @description   create a WorkflowProcess
// @param     workflowProcess               model.WorkflowProcess
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateWorkflowProcess(workflowProcess model.WorkflowProcess) (err error) {
	err = global.GVA_DB.Create(&workflowProcess).Error
	return err
}

// @title    DeleteWorkflowProcess
// @description   delete a WorkflowProcess
// @auth                     （2020/04/05  20:22）
// @param     workflowProcess               model.WorkflowProcess
// @return                    error

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

// @title    DeleteWorkflowProcessByIds
// @description   delete WorkflowProcesss
// @auth                     （2020/04/05  20:22）
// @param     workflowProcess               model.WorkflowProcess
// @return                    error

func DeleteWorkflowProcessByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.WorkflowProcess{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateWorkflowProcess
// @description   update a WorkflowProcess
// @param     workflowProcess          *model.WorkflowProcess
// @auth                     （2020/04/05  20:22）
// @return                    error

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

// @title    GetWorkflowProcess
// @description   get the info of a WorkflowProcess
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    WorkflowProcess        WorkflowProcess

func GetWorkflowProcess(id string) (err error, workflowProcess model.WorkflowProcess) {
	err = global.GVA_DB.Preload("Nodes").Preload("Edges").Where("id = ?", id).First(&workflowProcess).Error
	return
}

// @title    GetWorkflowProcessInfoList
// @description   get WorkflowProcess list by pagination, 分页获取WorkflowProcess
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

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
