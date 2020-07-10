package request

type DBReq struct {
	Database string `json:"database";gorm:"column:database"`
}

type TableReq struct {
	TableName string `json:"tableName"`
}

type ColumeReq struct {
	ColumeName    string `json:"columeName";gorm:"column:colume_name"`
	DataType      string `json:"dataType";gorm:"column:data_type"`
	DataTypeLong  string `json:"dataTypeLong";gorm:"column:data_type_long"`
	ColumeComment string `json:"columeComment";gorm:"column:colume_comment"`
}
