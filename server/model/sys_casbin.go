package model

type CasbinModel struct {
	Ptype       string `json:"ptype" gorm:"column:p_type"`
	AuthorityId string `json:"rolename" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}
