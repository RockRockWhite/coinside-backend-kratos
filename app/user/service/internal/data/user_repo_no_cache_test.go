package data

import (
	"context"
	"fmt"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userRepo UserRepo

func init() {
	userRepo = NewUserRepoNoCache(NewDB(config.NewConfig()))
}

func TestUserModelDefault_Insert(t *testing.T) {
	user := &User{
		Nickname:      "lllljjjjxxxxx",
		PasswdHash:    "",
		PasswdSalt:    "",
		Fullname:      "",
		Avatar:        "",
		Email:         "11111111",
		EmailVerified: false,
		Mobile:        "22222222",
		Config:        "",
	}

	_, err := userRepo.Insert(context.Background(), user)
	assert.Nil(t, err)
}

func TestUserModelDefault_FindOne(t *testing.T) {
	user, err := userRepo.FindOne(context.Background(), 1)
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestUserModelDefault_FindOneByMobile(t *testing.T) {
	user, err := userRepo.FindOneByMobile(context.Background(), "10086")
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestUserModelDefault_FindOneByEmail(t *testing.T) {
	user, err := userRepo.FindOne(context.Background(), 1)
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestUserModelDefault_FindOneByNickname(t *testing.T) {
	user, err := userRepo.FindOneByNickname(context.Background(), "admin")
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestUserModelDefault_Update(t *testing.T) {
	user, err := userRepo.FindOne(context.Background(), 1)
	assert.Nil(t, err)

	user.Avatar = "http://www.baidu.com"

	err = userRepo.Update(context.Background(), user)
	assert.Nil(t, err)

	fmt.Println(user)
}

func TestUserModelDefault_Delete(t *testing.T) {
	err := userRepo.Delete(context.Background(), 1)
	assert.Nil(t, err)
}
