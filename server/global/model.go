package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey" json:"ID"`           // 主键ID
	CreatedAt time.Time      `json:"createdAt"`                      // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`                      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"` // 删除时间
}
