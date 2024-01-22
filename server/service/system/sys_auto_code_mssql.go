package system

import (
	"fmt"
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

	sql := fmt.Sprintf(`select name as 'table_name' from %s.DBO.sysobjects where xtype='U'`, dbName)
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}

// GetColumn 获取指定数据库和指定数据表的所有字段名,类型值等
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	sql := fmt.Sprintf(`
SELECT 
    sc.name AS column_name,
    st.name AS data_type, 
    sc.length AS data_type_long,
    CASE 
        WHEN pk.index_id IS NOT NULL THEN 1
        ELSE 0
    END AS primary_key
FROM 
    %s.DBO.syscolumns sc
JOIN 
    systypes st ON sc.xtype=st.xtype
LEFT JOIN 
    %s.DBO.sysobjects so ON so.name='%s' AND so.xtype='U'
LEFT JOIN 
    %s.DBO.sysindexes si ON si.id = so.id AND si.indid < 2
LEFT JOIN 
    %s.DBO.sysindexkeys sik ON sik.id = si.id AND sik.indid = si.indid AND sik.colid = sc.colid
LEFT JOIN 
    %s.DBO.syskeyconstraints pk ON pk.constid = sik.constid
WHERE 
    st.usertype=0 AND sc.id = so.id
`, dbName, dbName, tableName, dbName, dbName, dbName)

	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}
