package request

type SysAutoCode struct {
	ID    uint   `gorm:"primarykey" json:"ID"` // 主键ID
	Label string `json:"label" gorm:"comment:展示名"`
	Desc  string `json:"desc" gorm:"comment:描述"`
}
