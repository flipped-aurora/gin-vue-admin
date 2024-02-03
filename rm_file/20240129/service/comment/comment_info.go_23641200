package comment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/comment"
    commentReq "github.com/flipped-aurora/gin-vue-admin/server/model/comment/request"
)

type CommentInfoService struct {
}

// CreateCommentInfo 创建评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService) CreateCommentInfo(commentData *comment.CommentInfo) (err error) {
	err = global.GVA_DB.Create(commentData).Error
	return err
}

// DeleteCommentInfo 删除评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService)DeleteCommentInfo(id string) (err error) {
	err = global.GVA_DB.Delete(&comment.CommentInfo{},"id = ?",id).Error
	return err
}

// DeleteCommentInfoByIds 批量删除评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService)DeleteCommentInfoByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]comment.CommentInfo{},"id in ?",ids).Error
	return err
}

// UpdateCommentInfo 更新评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService)UpdateCommentInfo(commentData comment.CommentInfo) (err error) {
	err = global.GVA_DB.Save(&commentData).Error
	return err
}

// GetCommentInfo 根据id获取评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService)GetCommentInfo(id string) (commentData comment.CommentInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&commentData).Error
	return
}

// GetCommentInfoInfoList 分页获取评论信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentDataService *CommentInfoService)GetCommentInfoInfoList(info commentReq.CommentInfoSearch) (list []comment.CommentInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&comment.CommentInfo{})
    var commentDatas []comment.CommentInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CommentUserId != nil {
        db = db.Where("comment_user_id = ?",info.CommentUserId)
    }
    if info.CommentDisId != nil {
        db = db.Where("comment_dis_id = ?",info.CommentDisId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&commentDatas).Error
	return  commentDatas, total, err
}
