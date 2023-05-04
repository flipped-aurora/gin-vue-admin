package clothing

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type AppJwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

// TableName appJwtBlacklist 表名
func (AppJwtBlacklist) TableName() string {
	return "app_jwt_blacklist"
}
