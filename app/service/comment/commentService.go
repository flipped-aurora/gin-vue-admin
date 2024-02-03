package comment

import (
	"app/model/comment"
	"app/mysqlDB"
)

type CommentInfoService struct {
}

// CreateCommentInfo 创建评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) CreateCommentInfo(commentData *comment.CommentInfo) (err error) {
	err = mysqlDB.DB.Create(commentData).Error
	return err
}

// CreateCommentSonInfo 创建评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) CreateCommentSonInfo(commentData *comment.CommentSonInfo) (err error) {
	err = mysqlDB.DB.Create(commentData).Error
	return err
}

// DeleteCommentInfo 删除评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) DeleteCommentInfo(id string) (err error) {
	err = mysqlDB.DB.Delete(&comment.CommentInfo{}, "id = ?", id).Error
	return err
}

// DeleteCommentSonInfo 删除评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) DeleteCommentSonInfo(id string) (err error) {
	err = mysqlDB.DB.Delete(&comment.CommentSonInfo{}, "id = ?", id).Error
	return err
}

// DeleteCommentInfoByIds 批量删除评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) DeleteCommentInfoByIds(ids []string) (err error) {
	err = mysqlDB.DB.Delete(&[]comment.CommentInfo{}, "id in ?", ids).Error
	return err
}

// UpdateCommentInfo 更新评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) UpdateCommentInfo(commentData comment.CommentInfo) (err error) {
	err = mysqlDB.DB.Updates(&commentData).Error
	return err
}

// UpdateCommentSonInfo 更新评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) UpdateCommentSonInfo(commentData comment.CommentSonInfo) (err error) {
	err = mysqlDB.DB.Updates(&commentData).Error
	return err
}

// GetCommentInfo 根据id获取评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) GetCommentInfo(id string) (commentData comment.CommentInfo, err error) {
	err = mysqlDB.DB.Where("id = ?", id).Preload("User").Preload("CommentSons").First(&commentData).Error
	return
}

// GetCommentInfoInfoList 分页获取评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) GetCommentInfoInfoList(info comment.CommentInfoSearch) (list []comment.CommentInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := mysqlDB.DB.Model(&comment.CommentInfo{})
	var commentDatas []comment.CommentInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CommentUserId != nil {
		db = db.Where("comment_user_id = ?", info.CommentUserId)
	}
	if info.CommentDisId != nil {
		db = db.Where("comment_dis_id = ?", info.CommentDisId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&commentDatas).Error
	return commentDatas, total, err
}
