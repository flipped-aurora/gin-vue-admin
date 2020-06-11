package model

import "github.com/jinzhu/gorm"

// 工作流流转表
type SysWorkFlowProcess struct {
	gorm.Model
	ApplicationID  uint   `json:"applicationID" gorm:"comment:'当前工作流所属申请的ID'"` // 当前工作流所属申请的ID
	CurrentNode    string `json:"currentNode" gorm:"comment:'当前进度节点'"`         // 当前进度节点
	HistoricalNode string `json:"historicalNode" gorm:"comment:'上一个进度节点'"`     // 上一个进度节点
	CurrentUser    string `json:"currentUser" gorm:"comment:'当前进度操作人'"`        // 当前进度操作人
	HistoricalUser string `json:"historicalUser" gorm:"comment:'上一个进度的操作人'"`   // 上一个进度的操作人
	State          bool   `json:"state" gorm:"comment:'状态 是否是正在进行的状态'"`        // 状态 是否是正在进行的状态
}
