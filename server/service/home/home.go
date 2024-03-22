package home

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type HomeService struct {
}

type Search struct {
	global.GVA_MODEL
	CateId string `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:栏目分类id;size:10;"`
	Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:191;"`
	Thumb  string `json:"thumb" form:"thumb" gorm:"column:thumb;comment:缩略图;size:191;"`
	Desc   string `json:"desc" form:"desc" gorm:"column:desc;comment:介绍;size:1024;"`
}

func (h *HomeService) SearchInfoList(pageInfo request.PageInfo) (searchlist []Search, err error) {
	var classlist, courselist []Search
	filed := "`id`,`cate_id`,`title`,`thumb`,`desc`"
	err = global.GVA_DB.Table("class").Select(filed).Where("title like ?", fmt.Sprint("%", pageInfo.Keyword, "%")).Find(&classlist).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Table("course").Select(filed).Where("title like ?", fmt.Sprint("%", pageInfo.Keyword, "%")).Find(&courselist).Error
	if err != nil {
		return
	}
	searchlist = append(append(searchlist, classlist...), courselist...)
	return
}
