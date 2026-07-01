package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(general config.GeneralDB) *gorm.Config {
	return &gorm.Config{
		Logger: NewGormLogger(general),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   general.Prefix,
			SingularTable: general.Singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
