package data

import (
	"gorm.io/gorm"
	"time"
)

type Markdown struct {
	Id        uint64         `gorm:"primaryKey, column:id"` // markdown的id
	CardId    uint64         `gorm:"column:card_id"`        // 卡片id
	Content   string         `gorm:"column:content"`        // 内容
	CreatedAt time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间

}

func (Markdown) TableName() string {
	return "m_markdown"
}
