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
	SELECT
    lower(a.COLUMN_NAME) as "column_name",
    (CASE WHEN a.DATA_TYPE = 'NUMBER' AND a.DATA_SCALE=0 THEN 'int' else lower(a.DATA_TYPE) end)  as "data_type",
    (CASE WHEN a.DATA_TYPE = 'NUMBER' THEN a.DATA_PRECISION else a.DATA_LENGTH end) as "data_type_long",
    b.COMMENTS as "column_comment",
    (CASE WHEN pk.COLUMN_NAME IS NOT NULL THEN 1 ELSE 0 END) as "primary_key",
    a.COLUMN_ID
FROM
    all_tab_columns a
JOIN
    all_col_comments b ON a.OWNER = b.OWNER AND a.TABLE_NAME = b.TABLE_NAME AND a.COLUMN_NAME = b.COLUMN_NAME
LEFT JOIN
    (
        SELECT
            acc.OWNER,
            acc.TABLE_NAME,
            acc.COLUMN_NAME
        FROM
            all_cons_columns acc
        JOIN
            all_constraints ac ON acc.OWNER = ac.OWNER AND acc.CONSTRAINT_NAME = ac.CONSTRAINT_NAME
        WHERE
            ac.CONSTRAINT_TYPE = 'P'
    ) pk ON a.OWNER = pk.OWNER AND a.TABLE_NAME = pk.TABLE_NAME AND a.COLUMN_NAME = pk.COLUMN_NAME
WHERE
    lower(a.table_name) = ?
    AND lower(a.OWNER) = ?
ORDER BY
    a.COLUMN_ID
`

	err = global.GVA_DBList[businessDB].Raw(sql, tableName, dbName).Scan(&entities).Error
	return entities, err
}
