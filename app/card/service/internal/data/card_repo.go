package data

import (
	"context"
)

type CardRepo interface {
	Insert(ctx context.Context, data *Card) (uint64, error)
	FindOne(ctx context.Context, id uint64) (*Card, error)
	FindAll(ctx context.Context, filter string) ([]*Card, error)
	Update(ctx context.Context, newData *Card) error
	Delete(ctx context.Context, id uint64) error

	InsertTag(ctx context.Context, id uint64, content string) error
	DeleteTag(ctx context.Context, id uint64, content string) error

	SetMember(ctx context.Context, id uint64, userId uint64, admin bool) error
	DeleteMember(ctx context.Context, id uint64, userId uint64) error
}
