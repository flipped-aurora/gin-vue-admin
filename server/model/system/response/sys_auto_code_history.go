package response

import "time"

type AutoCodeHistory struct {
	ID           uint      `json:"ID" gorm:"column:id"`
	CreatedAt    time.Time `json:"CreatedAt" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"UpdatedAt" gorm:"column:updated_at"`
	BusinessDB   string    `json:"businessDB" gorm:"column:business_db"`
	TableName    string    `json:"tableName" gorm:"column:table_name"`
	StructName   string    `json:"structName" gorm:"column:struct_name"`
	StructCNName string    `json:"structCNName" gorm:"column:struct_cn_name"`
	Flag         int       `json:"flag" gorm:"column:flag"`
}
