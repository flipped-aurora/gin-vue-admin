package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
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
	err = global.GVA_DB.Delete(workflowProcess).Error
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
	err = global.GVA_DB.Save(workflowProcess).Error
	return err
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
