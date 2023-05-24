package oauth

import "strings"

const (
	// 用户相关

	// ScopeUserInfo 获取用户公开信息
	ScopeUserInfo = "user_info"
	// ScopeFansList 粉丝列表
	ScopeFansList = "fans.list"
	// ScopeFollowingList 关注列表
	ScopeFollowingList = "following.list"
	// ScopeFansData 查询创作者粉丝数据
	ScopeFansData = "fans.data"

	// 视频相关

	// ScopeVideoCreate 上传视频到文件服务器 - 创建抖音视频 -上传图片到文件服务器 - 发布图片
	ScopeVideoCreate = "video.create"
	// ScopeVideoList 列出已发布的视频
	ScopeVideoList = "video.list"
	// ScopeVideoData 查询指定视频数据
	ScopeVideoData = "video.data"
	// ScopeAwemeshare 抖音分享id机制
	ScopeAwemeshare = "aweme.share"
	// ScopeVideoDelete 删除抖音视频
	ScopeVideoDelete = "video.delete"
	// ScopeHotsearch 获取实时热点词 --获取热点词聚合的视频
	ScopeHotsearch = "hotsearch"

	// 互动

	// ScopeVideoComment 评论列表 ---评论回复列表 ---回复视频评论 ---置顶视频评论(企业号)
	ScopeVideoComment = "video.comment"
	// ScopeIm 给抖音用户发送消息  --- 上传素材
	ScopeIm = "im"
)

// GetUserScope 获取用户相关Scope.
func GetUserScope() string {
	scopes := []string{ScopeUserInfo, ScopeFansList, ScopeFollowingList, ScopeFansData}
	return strings.Join(scopes, ",")
}

// GetVideoScope 获取视频相关Scope.
func GetVideoScope() string {
	scopes := []string{ScopeVideoCreate, ScopeVideoList, ScopeVideoData, ScopeAwemeshare, ScopeVideoDelete, ScopeHotsearch}
	return strings.Join(scopes, ",")
}

// GetInteractScope 获取互动相关Scope.
func GetInteractScope() string {
	scopes := []string{ScopeVideoComment, ScopeIm}
	return strings.Join(scopes, ",")
}

// GetAllScope 获取所有Scope.
func GetAllScope() string {
	scopes := []string{GetInteractScope(), GetVideoScope(), GetUserScope()}
	return strings.Join(scopes, ",")
}
