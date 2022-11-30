package data

import (
	"gorm.io/gorm"
	"time"
)

type Attachment struct {
	Id            uint64         `gorm:"primaryKey, column:id"` // attachment的id
	CardId        uint64         `gorm:"column:card_id"`        // 卡片id
	Link          string         `gorm:"column:link"`           // Attachment链接
	DownloadCount uint64         `gorm:"column:download_count"` //下载量
	CreatedAt     time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt     time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间

}

func (Attachment) TableName() string {
	return "m_attachment"
}
