package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz_apphub"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(biz_apphub.BizAppHub{}, biz_apphub.BizCmdToolApi{}, biz_apphub.BizToolCmdApi{}, biz_apphub.BizToolCmdSrvApi{})
	if err != nil {
		return err
	}
	return nil
}
