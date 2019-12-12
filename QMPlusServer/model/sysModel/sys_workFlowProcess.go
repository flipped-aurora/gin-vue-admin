package sysModel

import "github.com/jinzhu/gorm"

// 工作流流转表
type SysWorkFlowProcess struct {
	gorm.Model
	ApplicationID  uint   // 当前工作流所属申请的ID
	CurrentNode    string // 当前进度节点
	HistoricalNode string //上一个进度节点
	CurrentUser    string // 当前进度操作人
	HistoricalUser string // 上一个进度的操作人
	State          bool   // 状态 是否是正在进行的状态
}
