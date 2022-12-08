package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

var AutoCodeMssql = new(autoCodeMssql)

type autoCodeMssql struct{}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db
	sql := "select name AS 'database' from sysdatabases;"
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table

	sql := `use ? ;select name as 'table_name' from sysobjects where xtype='U'`
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql, dbName).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql, dbName).Scan(&entities).Error
	}

	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	sql := `use ?;select sc.name column_name
								,st.name data_type
	 from syscolumns sc,systypes st where sc.xtype=st.xtype and sc.id in(select id from sysobjects where xtype='U' and name=?);`
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql,dbName,tableName).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql, tableName, dbName).Scan(&entities).Error
	}

	return entities, err
}
