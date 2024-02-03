package competition

import (
	"app/model/competition"
	"app/mysqlDB"
)

type ComInfoService struct {
}

// CreateComInfo 创建比赛信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (comDataService *ComInfoService) CreateComInfo(comData *competition.ComInfo) (err error) {
	err = mysqlDB.DB.Create(comData).Error
	return err
}

// DeleteComInfo 删除比赛信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (comDataService *ComInfoService) DeleteComInfo(id string) (err error) {
	err = mysqlDB.DB.Delete(&competition.ComInfo{}, "id = ?", id).Error
	return err
}

// DeleteComInfoByIds 批量删除比赛信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (comDataService *ComInfoService) DeleteComInfoByIds(ids []string) (err error) {
	err = mysqlDB.DB.Delete(&[]competition.ComInfo{}, "id in ?", ids).Error
	return err
}

// UpdateComInfo 更新比赛信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (comDataService *ComInfoService) UpdateComInfo(comData competition.ComInfo) (err error) {
	err = mysqlDB.DB.Updates(&comData).Error
	return err
}

// GetComInfo 根据id获取比赛信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (comDataService *ComInfoService) GetComInfo(id string) (comData competition.ComInfo, err error) {
	err = mysqlDB.DB.Where("id = ?", id).First(&comData).Error
	return
}

// GetComInfoInfoList 分页获取比赛信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (comDataService *ComInfoService) GetComInfoInfoList(info competition.ComInfoSearch) (list []competition.ComInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := mysqlDB.DB.Model(&competition.ComInfo{})
	var comDatas []competition.ComInfo
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ComTitle != "" {
		db = db.Where("com_title LIKE ?", "%"+info.ComTitle+"%")
	}
	if info.ComModel != nil {
		db = db.Where("com_model = ?", info.ComModel)
	}
	if info.StartComStart != nil && info.EndComStart != nil {
		db = db.Where("com_start BETWEEN ? AND ? ", info.StartComStart, info.EndComStart)
	}
	if info.StartComEnd != nil && info.EndComEnd != nil {
		db = db.Where("com_end BETWEEN ? AND ? ", info.StartComEnd, info.EndComEnd)
	}
	if info.ComType != nil {
		db = db.Where("com_type = ?", info.ComType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&comDatas).Error
	return comDatas, total, err
}
func (this ComInfoService) GetInfo(id string) (reData interface{}, err error) {
	return this.GetComInfo(id)
}
