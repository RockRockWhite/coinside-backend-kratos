package data

import "context"

type UserRepo interface {
	Insert(ctx context.Context, data *User) (uint64, error)
	FindOne(ctx context.Context, id uint64) (*User, error)
	FindOneByNickname(ctx context.Context, nickname string) (*User, error)
	FindOneByEmail(ctx context.Context, email string) (*User, error)
	FindOneByMobile(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, newData *User) error
	Delete(ctx context.Context, id uint64) error
}
