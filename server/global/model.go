package global

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type GvaModel struct {
	ID        uint      `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
	DeletedAt time.Time             `json:"-"`                                                                                  // 删除时间
	IsDEL     soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt,default:0,comment:是否删除,index" json:"-"` // 使用 1 / 0 作为 Delete Flag
}
