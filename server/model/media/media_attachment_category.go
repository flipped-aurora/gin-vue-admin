package media

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type AttachmentCategory struct {
	global.GVA_MODEL
	Name     string                `json:"name" form:"name" gorm:"default:null;type:varchar(255);column:name;comment:分类名称;"`
	Pid      uint                  `json:"pid" form:"pid" gorm:"default:0;type:int;column:pid;comment:父节点ID;"`
	Children []*AttachmentCategory `json:"children" gorm:"-"`
}

func (AttachmentCategory) TableName() string {
	return "media_attachment_category"
}
