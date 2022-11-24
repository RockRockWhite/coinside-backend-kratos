package data

import (
	"context"
	"github.com/ljxsteam/coinside-backend-kratos/api/card"
	"gorm.io/gorm"
)

type CardRepo interface {
	Insert(ctx context.Context, data *Card) (uint64, error)
	FindOne(ctx context.Context, id uint64) (*Card, error)
	FindAll(ctx context.Context, limit uint64, offset uint64, filters []Filter) ([]Card, error)
	Update(ctx context.Context, newData *Card) error
	Delete(ctx context.Context, id uint64) error

	InsertTag(ctx context.Context, id uint64, content string) error
	DeleteTag(ctx context.Context, id uint64, content string) error

	SetMember(ctx context.Context, id uint64, userId uint64, admin bool) error
	DeleteMember(ctx context.Context, id uint64, userId uint64) error
}

type Filter interface {
	Filter(db *gorm.DB) *gorm.DB
}

type TeamFilter struct {
	teamId uint64
}

func (m TeamFilter) Filter(db *gorm.DB) *gorm.DB {
	return db.Where("team_id = ?", m.teamId)
}

func NewTeamFilter(teamId uint64) Filter {
	return &TeamFilter{teamId: teamId}
}

type StatusFilter struct {
	status card.CardStatus
}

func (s StatusFilter) Filter(db *gorm.DB) *gorm.DB {
	return db.Where("status = ?", s.status)
}

func NewStatusFilter(status card.CardStatus) Filter {
	return &StatusFilter{status: status}
}

type MemberFilter struct {
	userId uint64
}

func (m MemberFilter) Filter(db *gorm.DB) *gorm.DB {
	return db.Joins("JOIN c_member ON c_member.card_id = c_card.id AND c_member.user_id = ?", m.userId)
}

func NewMemberFilter(userId uint64) Filter {
	return &MemberFilter{userId: userId}
}

type TagFilter struct {
	content string
}

func (t TagFilter) Filter(db *gorm.DB) *gorm.DB {
	return db.
		Joins("JOIN c_card_tag ON c_card_tag.card_id = c_card.id").
		Joins("JOIN c_tag ON c_card_tag.tag_id = c_tag.id AND c_tag.content = ?", t.content)
}

func NewTagFilter(content string) Filter {
	return &TagFilter{content: content}
}
