package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"path/filepath"
	"strings"
)

var AutoCodeSqlite = new(autoCodeSqlite)

type autoCodeSqlite struct{}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodeSqlite) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db
	sql := "PRAGMA database_list;"
	var databaseList []struct {
		File string `gorm:"column:file"`
	}
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Find(&databaseList).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Find(&databaseList).Error
	}
	for _, database := range databaseList {
		if database.File != "" {
			fileName := filepath.Base(database.File)
			fileExt := filepath.Ext(fileName)
			fileNameWithoutExt := strings.TrimSuffix(fileName, fileExt)

			entities = append(entities, response.Db{fileNameWithoutExt})
		}
	}
	// entities = append(entities, response.Db{global.GVA_CONFIG.Sqlite.Dbname})
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodeSqlite) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table
	sql := `SELECT name FROM sqlite_master WHERE type='table'`
	tabelNames := []string{}
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Find(&tabelNames).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Find(&tabelNames).Error
	}
	for _, tabelName := range tabelNames {
		entities = append(entities, response.Table{tabelName})
	}
	return entities, err
}

// GetColumn 获取指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodeSqlite) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	sql := fmt.Sprintf("PRAGMA table_info(%s);", tableName)
	var columnInfos []struct {
		Name string `gorm:"column:name"`
		Type string `gorm:"column:type"`
	}
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&columnInfos).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&columnInfos).Error
	}
	for _, columnInfo := range columnInfos {
		entities = append(entities, response.Column{
			ColumnName: columnInfo.Name,
			DataType:   columnInfo.Type,
		})
	}
	return entities, err
}
