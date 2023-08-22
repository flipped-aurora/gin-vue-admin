package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey"`        // 主键ID
	CreatedAt time.Time      `gorm:"column:created_at"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`    // 删除时间
}
