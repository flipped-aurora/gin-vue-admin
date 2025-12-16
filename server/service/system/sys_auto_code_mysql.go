package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

var AutoCodeMysql = new(autoCodeMysql)

type autoCodeMysql struct{}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMysql) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db
	sql := "SELECT SCHEMA_NAME AS `database` FROM INFORMATION_SCHEMA.SCHEMATA;"
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
func (s *autoCodeMysql) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table
	sql := `select table_name as table_name from information_schema.tables where table_schema = ?`
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
func (s *autoCodeMysql) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	sql := `
	SELECT 
    c.COLUMN_NAME column_name,
    c.DATA_TYPE data_type,
    CASE 
        WHEN c.DATA_TYPE IN ('varchar', 'char', 'text', 'tinytext', 'mediumtext', 'longtext') THEN 
            COALESCE(CAST(c.CHARACTER_MAXIMUM_LENGTH AS CHAR), '')
        WHEN c.DATA_TYPE IN ('decimal', 'numeric', 'float', 'double') THEN 
            CONCAT_WS(',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE)
        WHEN c.DATA_TYPE IN ('tinyint', 'smallint', 'mediumint', 'int', 'integer', 'bigint') THEN 
            COALESCE(CAST(c.NUMERIC_PRECISION AS CHAR), '')
        WHEN c.DATA_TYPE IN ('datetime', 'timestamp', 'time') THEN 
            COALESCE(CAST(c.DATETIME_PRECISION AS CHAR), '')
        ELSE '' 
    END AS data_type_long,
    c.COLUMN_COMMENT column_comment,
    CASE WHEN kcu.COLUMN_NAME IS NOT NULL THEN 1 ELSE 0 END AS primary_key,
    CASE WHEN c.DATA_TYPE = 'enum' THEN 'enum' ELSE '' END AS enum_type,
    CASE 
        WHEN c.DATA_TYPE = 'enum' THEN 
            REPLACE(REPLACE(SUBSTRING(c.COLUMN_TYPE, 6), '(', ''), ')', '')
        ELSE '' 
    END AS enum_values,
    COALESCE(GROUP_CONCAT(DISTINCT s.INDEX_NAME ORDER BY s.INDEX_NAME SEPARATOR ','), '') AS index_name,
    COALESCE(GROUP_CONCAT(DISTINCT 
        CASE 
            WHEN s.INDEX_NAME = 'PRIMARY' THEN 'PRIMARY'
            WHEN s.NON_UNIQUE = 0 THEN 'UNIQUE'
            ELSE 'INDEX'
        END
        ORDER BY s.INDEX_NAME SEPARATOR ','), '') AS index_type,
    c.ORDINAL_POSITION
FROM 
    INFORMATION_SCHEMA.COLUMNS c
LEFT JOIN 
    INFORMATION_SCHEMA.KEY_COLUMN_USAGE kcu 
ON 
    c.TABLE_SCHEMA = kcu.TABLE_SCHEMA 
    AND c.TABLE_NAME = kcu.TABLE_NAME 
    AND c.COLUMN_NAME = kcu.COLUMN_NAME 
    AND kcu.CONSTRAINT_NAME = 'PRIMARY'
LEFT JOIN
    INFORMATION_SCHEMA.STATISTICS s
ON
    c.TABLE_SCHEMA = s.TABLE_SCHEMA
    AND c.TABLE_NAME = s.TABLE_NAME
    AND c.COLUMN_NAME = s.COLUMN_NAME
WHERE 
    c.TABLE_NAME = ? 
    AND c.TABLE_SCHEMA = ?
GROUP BY
    c.COLUMN_NAME, c.DATA_TYPE, c.CHARACTER_MAXIMUM_LENGTH, c.NUMERIC_PRECISION, 
    c.NUMERIC_SCALE, c.DATETIME_PRECISION, c.COLUMN_COMMENT, kcu.COLUMN_NAME,
    c.COLUMN_TYPE, c.ORDINAL_POSITION
ORDER BY 
    c.ORDINAL_POSITION;`
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql, tableName, dbName).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql, tableName, dbName).Scan(&entities).Error
	}

	return entities, err
}
