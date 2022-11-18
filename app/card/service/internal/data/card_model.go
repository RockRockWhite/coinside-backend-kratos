package data

import (
	"github.com/ljxsteam/coinside-backend-kratos/api/card"
	"gorm.io/gorm"
	"time"
)

type Card struct {
	Id        uint64          `gorm:"primaryKey, column:id"` // id
	TeamId    uint64          `gorm:"column:team_id"`        // 团队id
	Title     string          `gorm:"column:title"`          // 卡片标题
	Content   string          `gorm:"column:content"`        // 卡片详细内容，以标记语言存储
	Status    card.CardStatus `gorm:"column:status"`         // 卡片状态, 0：进行中，1：已完成
	CreatedAt time.Time       `gorm:"column:created_at"`     // 创建时间
	UpdatedAt time.Time       `gorm:"column:updated_at"`     // 更新时间
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at"`     // 删除时间

	Members []Member `gorm:"foreignKey:CardId;references:Id"` // 卡片成员
	Tags    []Tag    `gorm:"many2many:c_card_tag"`            // 卡片标签
}

func (Card) TableName() string {
	return "c_card"
}

type Member struct {
	Id        uint64         `gorm:"primaryKey, column:id"` // id
	CardId    uint64         `gorm:"column:card_id"`        // 所属卡片id
	UserId    uint64         `gorm:"column:user_id"`        // 用户id
	IsAdmin   bool           `gorm:"column:is_admin"`       // 管理员：0：非管理员 1：管理员
	CreatedAt time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间,软删除支持字段
}

func (Member) TableName() string {
	return "c_member"
}

type Tag struct {
	Id        uint64         `gorm:"primaryKey, column:id"` // id
	Content   string         `gorm:"column:content"`        // 标签内容
	CreatedAt time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间
}

func (Tag) TableName() string {
	return "c_tag"
}
