// 自动生成模板CliSet
package xiao

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

// 结算设置 结构体  CliSet
type CliSet struct {
	global.GVA_MODEL
	Static   *decimal.Decimal `json:"static" form:"static" gorm:"type:decimal(14,4);column:static;comment:静态收益率;"`                 //静态收益
	Straight *decimal.Decimal `json:"straight" form:"straight" gorm:"type:decimal(14,4);column:straight;comment:直推收益;"`            //直推收益
	Inderict *decimal.Decimal `json:"inderict" form:"inderict" gorm:"type:decimal(14,4);column:inderict;comment:间接推荐;"`            //间接推荐
	Fee      *decimal.Decimal `json:"fee" form:"fee" gorm:"type:decimal(14,4);column:fee;comment:提币手续费;"`                          //提币手续费
	Assets   datatypes.JSON   `json:"assets" form:"assets" gorm:"column:assets;comment:产品类型;type:text;"swaggertype:"array,object"` //产品类型
	Status   string           `json:"status" form:"status" gorm:"default:正常;column:status;comment:状态;"`                            //状态
	Desc     string           `json:"desc" form:"desc" gorm:"column:desc;comment:文本备注;"`                                           //文本备注
	Descnum  *decimal.Decimal `json:"descnum" form:"descnum" gorm:"type:decimal(14,4);default:0;column:descnum;comment:金额备注;"`     //金额备注
}

// TableName 结算设置 CliSet自定义表名 cli_set
func (CliSet) TableName() string {
	return "cli_set"
}
