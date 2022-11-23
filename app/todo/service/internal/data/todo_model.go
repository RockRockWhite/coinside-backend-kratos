package data

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	Id        uint64         `gorm:"primaryKey, column:id"` // todo的id
	CardId    uint64         `gorm:"column:card_id"`        // 卡片id
	Title     string         `gorm:"column:title"`          // 标题
	CreatedAt time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间

	Items []Item `gorm:"foreignKey:TodoId;references:Id"` //待办项

}

func (Todo) TableName() string {
	return "m_todo"
}

type Item struct {
	Id             uint64         `gorm:"primaryKey, column:id"`   // id
	TodoId         uint64         `gorm:"column:todo_id"`          //  待办id
	Content        string         `gorm:"column:content"`          //  待办项内容
	IsFinished     bool           `gorm:"column:is_finished"`      // 是否完成
	FinishedUserId uint64         `gorm:"column:finished_user_id"` //完成用户id
	CreatedAt      time.Time      `gorm:"column:created_at"`       // 创建时间
	UpdatedAt      time.Time      `gorm:"column:updated_at"`       // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at"`       // 删除时间
}

func (Item) TableName() string {
	return "m_todo_item"
}
