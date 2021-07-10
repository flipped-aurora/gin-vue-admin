package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
	"strings"

	"go.uber.org/zap"
)

// CreateAutoCodeHistory RouterPath : RouterPath@RouterString;RouterPath2@RouterString2
func CreateAutoCodeHistory(autoCodeMeta string, injectionMeta string, tableName string) error {
	return global.GVA_DB.Create(&model.SysAutoCodeHistory{
		AutoCodeMeta:  autoCodeMeta,
		InjectionMeta: injectionMeta,
		TableName:     tableName,
	}).Error
}

func RollBack(id uint) error {
	md := model.SysAutoCodeHistory{}
	if err := global.GVA_DB.First(&md, id).Error; err != nil {
		return err
	}
	// 切分数据
	err, dbNames := GetTables(global.GVA_CONFIG.Mysql.Dbname)
	if err != nil {
		return err
	}
	// 删除表
	for _, name := range dbNames {
		if strings.Contains(strings.ToUpper(strings.Replace(name.TableName, "_", "", -1)), strings.ToUpper(md.TableName)) {
			// 删除表
			if err = DropTable(name.TableName); err != nil {
				global.GVA_LOG.Error("ClearTag DropTable:", zap.Error(err))

			}
		}
	}
	// 删除文件

	for _, path := range strings.Split(md.AutoCodeMeta, ";") {
		_ = utils.DeLFile(path)
	}
	// 清除注入
	for _, v := range strings.Split(md.InjectionMeta, ";") {
		// RouterPath@functionName@RouterString
		meta := strings.Split(v, "@")
		if len(meta) != 3 {
			return errors.New("split InjectionMeta Err")
		}
		_ = utils.AutoClearCode(meta[0], meta[2])
	}
	md.Flag = 1
	return global.GVA_DB.Save(&md).Error
}
