package data

import (
	"context"
)

type TodoRepo interface {
	Insert(ctx context.Context, data *Todo) (uint64, error)
	FindOne(ctx context.Context, id uint64) (*Todo, error)
	Update(ctx context.Context, newData *Todo) error
	Delete(ctx context.Context, id uint64) error

	InsertItem(ctx context.Context, id uint64, content string) (uint64, error)
	DeleteItem(ctx context.Context, id uint64, itemId uint64) error
	FinishItem(ctx context.Context, id uint64, itemId uint64, finishedUserId uint64) error
	RestartItem(ctx context.Context, id uint64, itemId uint64) error
	UpdateContent(ctx context.Context, id uint64, itemId uint64, content string) error
}
