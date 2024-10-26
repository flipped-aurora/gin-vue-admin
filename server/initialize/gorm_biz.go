package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/xiao"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(xiao.CliLoad{}, xiao.CliUser{}, xiao.CliTree{}, xiao.CliOrder{}, xiao.CliMainorder{}, xiao.CliMainprofit{}, xiao.CliProfit{}, xiao.CliWithdraw{}, xiao.CliMainwith{}, xiao.CliSetvip{}, xiao.CliSet{})
	if err != nil {
		return err
	}
	return nil
}
