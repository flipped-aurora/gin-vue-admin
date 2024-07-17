package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func bizModel() error {
	err := global.GVA_DB.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
