package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strings"

	"go.uber.org/zap"
)

// CreateAutoCodeHistory RouterPath : RouterPath@RouterString;RouterPath2@RouterString2
func CreateAutoCodeHistory(meta, autoCodePath string, injectionMeta string, tableName string, apiIds string) error {
	return global.GVA_DB.Create(&model.SysAutoCodeHistory{
		RequestMeta:   meta,
		AutoCodePath:  autoCodePath,
		InjectionMeta: injectionMeta,
		TableName:     tableName,
		ApiIDs:        apiIds,
	}).Error
}

// RollBack 回滚
func RollBack(id uint) error {
	md := model.SysAutoCodeHistory{}
	if err := global.GVA_DB.First(&md, id).Error; err != nil {
		return err
	}
	// 清除API表
	err := DeleteApiByIds(strings.Split(md.ApiIDs, ";"))
	if err != nil {
		global.GVA_LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
	}
	// 获取全部表名
	err, dbNames := GetTables(global.GVA_CONFIG.Mysql.Dbname)
	if err != nil {
		global.GVA_LOG.Error("ClearTag GetTables:", zap.Error(err))
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
	for _, path := range strings.Split(md.AutoCodePath, ";") {
		_ = utils.DeLFile(path)
	}
	// 清除注入
	for _, v := range strings.Split(md.InjectionMeta, ";") {
		// RouterPath@functionName@RouterString
		meta := strings.Split(v, "@")
		if len(meta) == 3 {
			_ = utils.AutoClearCode(meta[0], meta[2])
		}
	}
	md.Flag = 1
	return global.GVA_DB.Save(&md).Error
}

func GetMeta(id uint) (string, error) {
	var meta string
	return meta, global.GVA_DB.Model(model.SysAutoCodeHistory{}).Select("request_meta").First(&meta, id).Error
}

func GetSysHistoryPage(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var fileLists []model.SysAutoCodeHistory
	err = db.Find(&fileLists).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Select("id,created_at,updated_at,table_name,auto_code_path,injection_meta,api_ids,flag").Find(&fileLists).Error
	return err, fileLists, total
}
