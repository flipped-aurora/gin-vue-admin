package comment

import (
	"app/model/request"
	"app/model/user"
	"gorm.io/gorm"
	"time"
)

// 评论信息 结构体  CommentInfo
type CommentInfo struct {
	ID                uint             `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt         time.Time        // 创建时间
	UpdatedAt         time.Time        // 更新时间
	DeletedAt         gorm.DeletedAt   `gorm:"index" json:"-"`
	CommentUserId     uint             `json:"commentUserId" form:"commentUserId" gorm:"column:comment_user_id;comment:评论人id;"` //评论人id
	User              user.UserInfo    `json:"userInfo" form:"userInfo" gorm:"foreignKey:CommentUserId"`
	CommentDisId      uint             `json:"commentDisId" form:"commentDisId" gorm:"column:comment_dis_id;comment:关联帖子id;"` //关联帖子id
	CommentSons       []CommentSonInfo `json:"commentSons" form:"commentSons" gorm:"foreignKey:CommentCommentId"`
	CommentText       string           `json:"commentText" form:"commentText" gorm:"column:comment_text;comment:评论信息;size:300;"`              //评论信息
	CommentLikeNumber int              `json:"commentLikeNumber" form:"commentLikeNumber" gorm:"column:comment_like_number;comment:评论的评论id;"` //评论点赞
}

// TableName 评论信息 CommentInfo自定义表名 commentInfo
func (CommentInfo) TableName() string {
	return "commentInfo"
}

type CommentInfoSearch struct {
	StartCreatedAt   *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt     *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	CommentUserId    *int       `json:"commentUserId" form:"commentUserId" `
	CommentCommentId *int       `json:"commentCommentId" form:"commentCommentId" `
	CommentDisId     *int       `json:"commentDisId" form:"commentDisId" `
	request.PageInfo
}
type CommentSonInfo struct {
	ID               uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt        time.Time      // 创建时间
	UpdatedAt        time.Time      // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	CommentCommentId uint           `json:"commentCommentId" form:"commentCommentId" gorm:"column:comment_comment_id;comment:评论的评论id;"` //评论的评论id
	CommentUserId    uint           `json:"commentUserId" form:"commentUserId" gorm:"column:comment_user_id;comment:评论人id;"`            //评论人id
	User             user.UserInfo  `json:"userInfo" form:"userInfo" gorm:"foreignKey:CommentUserId"`
	CommentText      string         `json:"commentText" form:"commentText" gorm:"column:comment_text;comment:评论信息;size:300;"` //评论信息
}

func (CommentSonInfo) TableName() string {
	return "commentSonInfo"
}
