package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

var AutoCodePgsql = new(autoCodePgsql)

type autoCodePgsql struct{}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePgsql) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db
	sql := `SELECT datname as database FROM pg_database WHERE datistemplate = false`
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
func (a *autoCodePgsql) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table
	sql := `select table_name as table_name from information_schema.tables where table_catalog = ? and table_schema = ?`

	db := global.GVA_DB
	if businessDB != "" {
		db = global.GVA_DBList[businessDB]
	}

	err = db.Raw(sql, dbName, "public").Scan(&entities).Error
	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePgsql) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	// todo 数据获取不全, 待完善sql
	sql := `
SELECT 
  psc.COLUMN_NAME AS COLUMN_NAME,
  psc.udt_name AS data_type,
  CASE
    psc.udt_name 
    WHEN 'text' THEN
      concat_ws ( '', '', psc.CHARACTER_MAXIMUM_LENGTH ) 
    WHEN 'varchar' THEN
      concat_ws ( '', '', psc.CHARACTER_MAXIMUM_LENGTH ) 
    WHEN 'smallint' THEN
      concat_ws ( ',', psc.NUMERIC_PRECISION, psc.NUMERIC_SCALE ) 
    WHEN 'decimal' THEN
      concat_ws ( ',', psc.NUMERIC_PRECISION, psc.NUMERIC_SCALE ) 
    WHEN 'integer' THEN
      concat_ws ( '', '', psc.NUMERIC_PRECISION ) 
    WHEN 'int4' THEN
      concat_ws ( '', '', psc.NUMERIC_PRECISION ) 
    WHEN 'int8' THEN
      concat_ws ( '', '', psc.NUMERIC_PRECISION ) 
    WHEN 'bigint' THEN
      concat_ws ( '', '', psc.NUMERIC_PRECISION ) 
    WHEN 'timestamp' THEN
      concat_ws ( '', '', psc.datetime_precision ) 
    ELSE '' 
  END AS data_type_long,
  (
    SELECT
      pd.description 
    FROM
      pg_description pd 
    WHERE
      (pd.objoid,pd.objsubid) in (
        SELECT pa.attrelid,pa.attnum
        FROM
          pg_attribute pa 
        WHERE pa.attrelid = ( SELECT oid FROM pg_class pc WHERE 
        pc.relname = psc.table_name
        ) 
        and attname = psc.column_name
      ) 
  ) AS column_comment,
  (
    SELECT
      COUNT(*) 
    FROM
      pg_constraint 
    WHERE
      contype = 'p' 
      AND conrelid = (
        SELECT
          oid 
        FROM
          pg_class 
        WHERE
          relname = psc.table_name
      ) 
      AND conkey::int[] @> ARRAY[(
        SELECT
          attnum 
        FROM
          pg_attribute 
        WHERE
          attrelid = conrelid 
          AND attname = psc.column_name
      )]
  ) > 0 AS primary_key
FROM
  INFORMATION_SCHEMA.COLUMNS psc 
WHERE
  table_catalog = ?
  AND table_schema = 'public' 
  AND TABLE_NAME = ?;
`
	var entities []response.Column
	//sql = strings.ReplaceAll(sql, "@table_catalog", dbName)
	//sql = strings.ReplaceAll(sql, "@table_name", tableName)
	db := global.GVA_DB
	if businessDB != "" {
		db = global.GVA_DBList[businessDB]
	}

	err = db.Raw(sql, dbName, tableName).Scan(&entities).Error
	return entities, err
}
