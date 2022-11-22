package data

import (
	"context"
)

type TodoRepo interface {
	Insert(ctx context.Context, data *Todo) (uint64, error)
	FindOne(ctx context.Context, id uint64) (*Todo, error)
	//FindOneByname(ctx context.Context, nickname string) (*User, error)
	Update(ctx context.Context, newData *Todo) error
	Delete(ctx context.Context, id uint64) error

	SetItem(ctx context.Context, data *Item) error
	DeleteItem(ctx context.Context, id uint64) error
}
