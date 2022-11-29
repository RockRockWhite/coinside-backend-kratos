package data

import (
	"gorm.io/gorm"
	"time"
)

type Vote struct {
	Id        uint64         `gorm:"primaryKey, column:id"` // id
	CardId    uint64         `gorm:"column:card_id"`        // 卡片id
	Title     string         `gorm:"column:title"`          // 投票标题
	CreatedAt time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间

	Items []Item `gorm:"foreignKey:VoteId;references:Id"` // 投票项

}

func (Vote) TableName() string {
	return "m_vote"
}

type Item struct {
	Id        uint64         `gorm:"primaryKey, column:id"`               // id
	VoteId    uint64         `gorm:"column:vote_id"`                      // 投票id
	Content   string         `gorm:"column:content"`                      // 内容
	CreatedAt time.Time      `gorm:"column:created_at"`                   // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at"`                   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`                   // 删除时间,软删除支持字段
	Commits   []Commit       `gorm:"foreignKey:VoteItemId;references:Id"` // 投票提交
}

func (Item) TableName() string {
	return "m_vote_item"
}

type Commit struct {
	Id         uint64         `gorm:"primaryKey, column:id"` // id
	VoteItemId uint64         `gorm:"column:vote_item_id"`   // 投票项id
	UserId     uint64         `gorm:"column:user_id"`        //用户id
	CreatedAt  time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt  time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间
}

func (Commit) TableName() string {
	return "m_vote_commit"
}
