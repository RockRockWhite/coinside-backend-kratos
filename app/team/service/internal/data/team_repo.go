package data

import (
	"context"
)

type TeamRepo interface {
	Insert(ctx context.Context, data *Team) (uint64, error)
	FindOne(ctx context.Context, id uint64) (*Team, error)
	//FindOneByname(ctx context.Context, nickname string) (*User, error)
	Update(ctx context.Context, newData *Team) error
	Delete(ctx context.Context, id uint64) error

	SetMember(ctx context.Context, id uint64, userId uint64, admin bool) error
	DeleteMember(ctx context.Context, id uint64, userId uint64) error
}
