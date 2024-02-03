package comment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CommentInfoRouter struct {
}

// InitCommentInfoRouter 初始化 评论信息 路由信息
func (s *CommentInfoRouter) InitCommentInfoRouter(Router *gin.RouterGroup) {
	commentDataRouter := Router.Group("commentData").Use(middleware.OperationRecord())
	commentDataRouterWithoutRecord := Router.Group("commentData")
	var commentDataApi = v1.ApiGroupApp.CommentApiGroup.CommentInfoApi
	{
		commentDataRouter.POST("createCommentInfo", commentDataApi.CreateCommentInfo)   // 新建评论信息
		commentDataRouter.DELETE("deleteCommentInfo", commentDataApi.DeleteCommentInfo) // 删除评论信息
		commentDataRouter.DELETE("deleteCommentInfoByIds", commentDataApi.DeleteCommentInfoByIds) // 批量删除评论信息
		commentDataRouter.PUT("updateCommentInfo", commentDataApi.UpdateCommentInfo)    // 更新评论信息
	}
	{
		commentDataRouterWithoutRecord.GET("findCommentInfo", commentDataApi.FindCommentInfo)        // 根据ID获取评论信息
		commentDataRouterWithoutRecord.GET("getCommentInfoList", commentDataApi.GetCommentInfoList)  // 获取评论信息列表
	}
}
