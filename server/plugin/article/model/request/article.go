package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ArticleSearch struct {
	Title      string     `json:"title" form:"title"`
	Lang       string     `json:"lang" form:"lang"`
	Status     *int       `json:"status" form:"status"`
	CategoryID uint       `json:"categoryID" form:"categoryID"`
	StartTime  *time.Time `json:"startTime" form:"startTime"`
	EndTime    *time.Time `json:"endTime" form:"endTime"`
	request.PageInfo
}

type CategorySearch struct {
	Name string `json:"name" form:"name"`
	Lang string `json:"lang" form:"lang"`
	request.PageInfo
}
