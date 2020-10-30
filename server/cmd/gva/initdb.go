/*
Copyright © 2020 SliverHorn

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gva

import (
	"gin-vue-admin/cmd/datas"
	"gin-vue-admin/core"
	"gin-vue-admin/initialize"

	"github.com/gookit/color"

	_ "gin-vue-admin/core"
	"gin-vue-admin/global"

	"github.com/spf13/cobra"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "gin-vue-admin初始化数据",
	Long: `gin-vue-admin初始化数据适配数据库情况: 
1. mysql完美适配,
2. postgresql不能保证完美适配,
3. sqlite未适配,
4. sqlserver未适配`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		global.GVA_VP = core.Viper(path)
		global.GVA_LOG = core.Zap()           // 初始化zap日志库
		db := initialize.GormMysql()
		switch global.GVA_CONFIG.System.DbType {
		case "mysql":
			datas.InitMysqlTables(db)
			datas.InitMysqlData(db)
		case "postgresql":
			color.Info.Println("postgresql功能开发中")
		case "sqlite":
			color.Info.Println("sqlite功能开发中")
		case "sqlserver":
			color.Info.Println("sqlserver功能开发中")
		default:
			datas.InitMysqlTables(db)
			datas.InitMysqlData(db)
		}
		frame, _ := cmd.Flags().GetString("frame")
		if frame == "gf" {
			color.Info.Println("gf功能开发中")
			return
		} else {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", "./config.yaml", "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("frame", "f", "gin", "可选参数为gin,gf")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}
