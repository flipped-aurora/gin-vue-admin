// 自动生成模板SysDictionaryDetail
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 如果含有time.Time 请自行import time包
type AutoCodeExample struct {
	global.GVA_MODEL
	AutoCodeExampleField string `json:"autoCodeExampleField" form:"autoCodeExampleField" gorm:"column:auto_code_example_field;comment:仅作示例条目无实际作用"` // 展示值
}
