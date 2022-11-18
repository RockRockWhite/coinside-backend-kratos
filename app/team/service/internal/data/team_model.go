package data

import (
	"gorm.io/gorm"
	"time"
)

type Team struct {
	Id          uint64         `gorm:"primaryKey, column:id"` // 团队id
	Name        string         `gorm:"column:name"`           // 团队名
	Description string         `gorm:"column:description"`    // 用户描述
	Website     string         `gorm:"column:website"`        // 团队网址
	Avatar      string         `gorm:"column:avatar"`         // 团队头像
	Email       string         `gorm:"column:email"`          // 团队邮箱
	CreatedAt   time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间

	Members []TeamMember `gorm:"foreignKey:TeamId;references:Id"` //团队成员

}

func (Team) TableName() string {
	return "t_team"
}

type TeamMember struct {
	Id        uint64         `gorm:"primaryKey, column:id"` // 团队成员id
	TeamId    uint64         `gorm:"column:team_id"`        // 团队id
	UserId    uint64         `gorm:"column:user_id"`        // 用户id
	IsAdmin   bool           `gorm:"column:is_admin"`       // 是否为管理员
	CreatedAt time.Time      `gorm:"column:created_at"`     // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at"`     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`     // 删除时间
}

func (TeamMember) TableName() string {
	return "t_member"
}
