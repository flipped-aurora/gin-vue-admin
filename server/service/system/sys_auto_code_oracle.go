package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

var AutoCodeOracle = new(autoCodeOracle)

type autoCodeOracle struct{}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeOracle) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db
	sql := `SELECT lower(username) AS "database" FROM all_users`
	err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeOracle) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table
	sql := `select lower(table_name) as "table_name" from all_tables where lower(owner) = ?`

	err = global.GVA_DBList[businessDB].Raw(sql, dbName).Scan(&entities).Error
	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeOracle) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	sql := `
		select lower(a.COLUMN_NAME) as "column_name",
			   (CASE WHEN a.DATA_TYPE = 'NUMBER' AND a.DATA_SCALE=0 THEN 'int' else lower(a.DATA_TYPE) end)  as "data_type", 
			   (CASE WHEN a.DATA_TYPE = 'NUMBER' THEN a.DATA_PRECISION else a.DATA_LENGTH end) as "data_type_long",
			   b.COMMENTS as "column_comment"
		from all_tab_columns a , all_col_comments b
		 where a.OWNER = b.OWNER
		 and a.TABLE_NAME = b.TABLE_NAME
		 and a.COLUMN_NAME = b.COLUMN_NAME
		and lower(a.table_name) = ?
		 and lower(a.OWNER) = ?		 
`

	err = global.GVA_DBList[businessDB].Raw(sql, tableName, dbName).Scan(&entities).Error
	return entities, err
}
