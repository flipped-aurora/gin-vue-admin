package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BizCloudFunctionSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    CnName  string `json:"cn_name" form:"cn_name" `
    CodeName  string `json:"code_name" form:"code_name" `
    Classify  string `json:"classify" form:"classify" `
    ExecMode  string `json:"exec_mode" form:"exec_mode" `
    Title  string `json:"title" form:"title" `
    Content  string `json:"content" form:"content" `
    ContentType  string `json:"content_type" form:"content_type" `
    IsPublic  string `json:"is_public" form:"is_public" `
    Tags  string `json:"tags" form:"tags" `
    StartViews  *int  `json:"startViews" form:"startViews"`
    EndViews  *int  `json:"endViews" form:"endViews"`
    StartExecCount  *int  `json:"startExecCount" form:"startExecCount"`
    EndExecCount  *int  `json:"endExecCount" form:"endExecCount"`
    StartColl  *int  `json:"startColl" form:"startColl"`
    EndColl  *int  `json:"endColl" form:"endColl"`
    StartLike  *int  `json:"startLike" form:"startLike"`
    EndLike  *int  `json:"endLike" form:"endLike"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
