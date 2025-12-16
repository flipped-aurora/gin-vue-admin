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
	sql := `
SELECT
    psc.COLUMN_NAME AS COLUMN_NAME,
    psc.udt_name AS data_type,
    CASE
        WHEN psc.udt_name IN ('text', 'varchar', 'char', 'character', 'character varying') THEN
            COALESCE(CAST(psc.CHARACTER_MAXIMUM_LENGTH AS TEXT), '')
        WHEN psc.udt_name IN ('smallint', 'integer', 'int', 'int4', 'bigint', 'int8', 'numeric', 'decimal') THEN
            CASE 
                WHEN psc.NUMERIC_SCALE IS NOT NULL AND psc.NUMERIC_SCALE > 0 THEN
                    CONCAT_WS(',', psc.NUMERIC_PRECISION, psc.NUMERIC_SCALE)
                ELSE
                    COALESCE(CAST(psc.NUMERIC_PRECISION AS TEXT), '')
            END
        WHEN psc.udt_name IN ('timestamp', 'timestamp without time zone', 'timestamp with time zone', 'time', 'time without time zone', 'time with time zone') THEN
            COALESCE(CAST(psc.datetime_precision AS TEXT), '')
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
                attnum::integer
            FROM
                pg_attribute
            WHERE
                attrelid = conrelid
              AND attname = psc.column_name
        )]
    ) > 0 AS primary_key,
    CASE 
        WHEN EXISTS (
            SELECT 1 
            FROM pg_type t 
            JOIN pg_enum e ON t.oid = e.enumtypid 
            WHERE t.typname = psc.udt_name
        ) THEN 'enum'
        ELSE ''
    END AS enum_type,
    CASE 
        WHEN EXISTS (
            SELECT 1 
            FROM pg_type t 
            JOIN pg_enum e ON t.oid = e.enumtypid 
            WHERE t.typname = psc.udt_name
        ) THEN (
            SELECT string_agg(quote_literal(e.enumlabel), ',')
            FROM pg_type t 
            JOIN pg_enum e ON t.oid = e.enumtypid 
            WHERE t.typname = psc.udt_name
            ORDER BY e.enumsortorder
        )
        ELSE ''
    END AS enum_values,
    COALESCE((
        SELECT string_agg(DISTINCT i.relname, ',' ORDER BY i.relname)
        FROM pg_index idx
        JOIN pg_class i ON i.oid = idx.indexrelid
        JOIN pg_class t ON t.oid = idx.indrelid
        JOIN pg_attribute a ON a.attrelid = t.oid AND a.attnum = ANY(idx.indkey)
        WHERE t.relname = psc.table_name
          AND a.attname = psc.column_name
    ), '') AS index_name,
    COALESCE((
        SELECT string_agg(DISTINCT 
            CASE 
                WHEN i.relname LIKE '%_pkey' THEN 'PRIMARY'
                WHEN idx.indisunique THEN 'UNIQUE'
                ELSE 'INDEX'
            END, ',' ORDER BY 
            CASE 
                WHEN i.relname LIKE '%_pkey' THEN 'PRIMARY'
                WHEN idx.indisunique THEN 'UNIQUE'
                ELSE 'INDEX'
            END
        )
        FROM pg_index idx
        JOIN pg_class i ON i.oid = idx.indexrelid
        JOIN pg_class t ON t.oid = idx.indrelid
        JOIN pg_attribute a ON a.attrelid = t.oid AND a.attnum = ANY(idx.indkey)
        WHERE t.relname = psc.table_name
          AND a.attname = psc.column_name
    ), '') AS index_type,
    psc.ordinal_position
FROM
    INFORMATION_SCHEMA.COLUMNS psc
WHERE
  table_catalog = ?
  AND table_schema = 'public' 
  AND TABLE_NAME = ?
ORDER BY
    psc.ordinal_position;
`
	var entities []response.Column
	db := global.GVA_DB
	if businessDB != "" {
		db = global.GVA_DBList[businessDB]
	}

	err = db.Raw(sql, dbName, tableName).Scan(&entities).Error
	return entities, err
}
