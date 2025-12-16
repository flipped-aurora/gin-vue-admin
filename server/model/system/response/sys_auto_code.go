package response

type Db struct {
	Database string `json:"database" gorm:"column:database"`
}

type Table struct {
	TableName string `json:"tableName" gorm:"column:table_name"`
}

type Column struct {
	DataType      string `json:"dataType" gorm:"column:data_type"`
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
	PrimaryKey    bool   `json:"primaryKey" gorm:"column:primary_key"`
	EnumType      string `json:"enumType" gorm:"column:enum_type"`      // 枚举类型（如：enum）
	EnumValues    string `json:"enumValues" gorm:"column:enum_values"`  // 枚举值（逗号分隔）
	IndexName     string `json:"indexName" gorm:"column:index_name"`    // 索引名称（多个索引用逗号分隔）
	IndexType     string `json:"indexType" gorm:"column:index_type"`     // 索引类型（PRIMARY, UNIQUE, INDEX等，多个用逗号分隔）
}
