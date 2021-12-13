package system

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var AutoCodePgsql = new(autoCodePgsql)

type autoCodePgsql struct{}

// GetDB 获取数据库的所有数据库名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePgsql) GetDB() (data []response.Db, err error) {
	var entities []response.Db
	sql := `SELECT datname as database FROM pg_database WHERE datistemplate = false`
	err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	return entities, err
}

// GetTables 获取数据库的所有表名
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePgsql) GetTables(dbName string) (data []response.Table, err error) {
	var entities []response.Table
	sql := `select table_name as table_name from information_schema.tables where table_catalog = ? and table_schema = ?`
	db, _err := gorm.Open(postgres.Open(global.GVA_CONFIG.Pgsql.LinkDsn(dbName)), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if _err != nil {
		return nil, errors.Wrapf(err, "[pgsql] 连接 数据库(%s)的表失败!", dbName)
	}
	err = db.Raw(sql, dbName, "public").Scan(&entities).Error
	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *autoCodePgsql) GetColumn(tableName string, dbName string) (data []response.Column, err error) {
	// todo 数据获取不全, 待完善sql
	sql := `
		SELECT columns.COLUMN_NAME                                                                                      as column_name,
		   columns.DATA_TYPE                                                                                        as data_type,
		   CASE
			   columns.DATA_TYPE
			   WHEN 'text' THEN
				   concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
			   WHEN 'varchar' THEN
				   concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
			   WHEN 'smallint' THEN
				   concat_ws(',', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
			   WHEN 'decimal' THEN
				   concat_ws(',', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
			   WHEN 'integer' THEN
				   concat_ws('', '', columns.NUMERIC_PRECISION)
			   WHEN 'bigint' THEN
				   concat_ws('', '', columns.NUMERIC_PRECISION)
			   ELSE ''
			   END                                                                                                  AS data_type_long,
		   (select description.description
			from pg_description description
			where description.objoid = (select attribute.attrelid
										from pg_attribute attribute
										where attribute.attrelid =
											  (select oid from pg_class class where class.relname = '@table_name') and attname =columns.COLUMN_NAME )
			  and description.objsubid = (select attribute.attnum
										  from pg_attribute attribute
										  where attribute.attrelid =
												(select oid from pg_class class where class.relname = '@table_name') and attname =columns.COLUMN_NAME )) as column_comment
		FROM INFORMATION_SCHEMA.COLUMNS columns
		WHERE table_catalog = '@table_catalog'
		  and table_schema = 'public'
		  and table_name = '@table_name';
	`
	var entities []response.Column
	db, _err := gorm.Open(postgres.Open(global.GVA_CONFIG.Pgsql.LinkDsn(dbName)), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if _err != nil {
		return nil, errors.Wrapf(err, "[pgsql] 连接 数据库(%s)的表(%s)失败!", dbName, tableName)
	}
	sql = strings.ReplaceAll(sql, "@table_catalog", dbName)
	sql = strings.ReplaceAll(sql, "@table_name", tableName)
	err = db.Raw(sql).Scan(&entities).Error
	return entities, err
}
