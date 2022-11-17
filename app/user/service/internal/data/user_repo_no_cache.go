package data

import (
	"context"
	"gorm.io/gorm"
)

type UserRepoNoCache struct {
	db *gorm.DB
}

func (u UserRepoNoCache) Insert(ctx context.Context, data *User) (uint64, error) {
	res := u.db.Create(data)
	return data.Id, res.Error
}

func (u UserRepoNoCache) FindOne(ctx context.Context, id uint64) (*User, error) {
	var user User
	res := u.db.Where("id = ?", id).First(&user)

	return &user, res.Error
}

func (u UserRepoNoCache) FindOneByNickname(ctx context.Context, nickname string) (*User, error) {
	var user User
	res := u.db.Where("nickname = ?", nickname).First(&user)

	return &user, res.Error
}

func (u UserRepoNoCache) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	res := u.db.Where("email = ?", email).First(&user)

	return &user, res.Error
}

func (u UserRepoNoCache) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	var user User
	res := u.db.Where("mobile = ?", mobile).First(&user)

	return &user, res.Error
}

func (u UserRepoNoCache) Update(ctx context.Context, newData *User) error {
	res := u.db.Save(newData)
	return res.Error
}

func (u UserRepoNoCache) Delete(ctx context.Context, id uint64) error {
	res := u.db.Delete(&User{}, id)
	return res.Error
}

func NewUserRepoNoCache(db *gorm.DB) UserRepo {
	return &UserRepoNoCache{db: db}
}
