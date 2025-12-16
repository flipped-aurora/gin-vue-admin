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
		Pk   int    `gorm:"column:pk"`
	}
	
	db := global.GVA_DB
	if businessDB != "" {
		db = global.GVA_DBList[businessDB]
	}
	
	err = db.Raw(sql).Scan(&columnInfos).Error
	if err != nil {
		return entities, err
	}
	
	// 获取索引信息
	indexMap := make(map[string][]struct {
		IndexName string
		IndexType string
	})
	
	indexListSQL := fmt.Sprintf("PRAGMA index_list(%s);", tableName)
	var indexList []struct {
		Name   string `gorm:"column:name"`
		Unique int    `gorm:"column:unique"`
	}
	err = db.Raw(indexListSQL).Scan(&indexList).Error
	if err == nil {
		for _, idx := range indexList {
			// 获取索引的列信息
			indexInfoSQL := fmt.Sprintf("PRAGMA index_info(%s);", idx.Name)
			var indexColumns []struct {
				SeqNo int    `gorm:"column:seqno"`
				Cid   int    `gorm:"column:cid"`
				Name  string `gorm:"column:name"`
			}
			err = db.Raw(indexInfoSQL).Scan(&indexColumns).Error
			if err == nil {
				for _, col := range indexColumns {
					indexType := "INDEX"
					if idx.Unique == 1 {
						indexType = "UNIQUE"
					}
					indexMap[col.Name] = append(indexMap[col.Name], struct {
						IndexName string
						IndexType string
					}{
						IndexName: idx.Name,
						IndexType: indexType,
					})
				}
			}
		}
	}
	
	for _, columnInfo := range columnInfos {
		// 解析类型和长度（SQLite类型定义如：VARCHAR(255), INTEGER等）
		dataType := columnInfo.Type
		dataTypeLong := ""
		enumType := ""
		enumValues := ""
		
		// 提取类型长度（如果有括号）
		if strings.Contains(dataType, "(") {
			parts := strings.SplitN(dataType, "(", 2)
			dataType = strings.TrimSpace(parts[0])
			if len(parts) > 1 {
				lengthPart := strings.TrimSuffix(parts[1], ")")
				dataTypeLong = strings.TrimSpace(lengthPart)
			}
		}
		
		// 获取该字段的索引信息
		indexNames := []string{}
		indexTypes := []string{}
		if indexes, ok := indexMap[columnInfo.Name]; ok {
			for _, idx := range indexes {
				indexNames = append(indexNames, idx.IndexName)
				indexTypes = append(indexTypes, idx.IndexType)
			}
		}
		
		// 如果是主键，添加 PRIMARY 类型
		if columnInfo.Pk == 1 {
			indexTypes = append(indexTypes, "PRIMARY")
		}
		
		entities = append(entities, response.Column{
			ColumnName:   columnInfo.Name,
			DataType:     dataType,
			DataTypeLong: dataTypeLong,
			PrimaryKey:   columnInfo.Pk == 1,
			EnumType:     enumType,
			EnumValues:   enumValues,
			IndexName:    strings.Join(indexNames, ","),
			IndexType:    strings.Join(indexTypes, ","),
		})
	}
	return entities, err
}
