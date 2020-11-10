package request

type DBReq struct {
	Database string `json:"database" gorm:"column:database"`
}

type TableReq struct {
	TableName string `json:"tableName"`
}

type ColumnReq struct {
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataType      string `json:"dataType" gorm:"column:data_type"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
}
