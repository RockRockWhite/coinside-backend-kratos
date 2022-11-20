package data

import (
	"context"
	"github.com/ljxsteam/coinside-backend-kratos/app/card/service/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var cardRepo CardRepo

func init() {
	cardRepo = NewCardRepoNoCache(NewDB(config.NewConfig()))
}

func TestCardModelDefault_Insert(t *testing.T) {

	card := &Card{
		Title:   "单元测试标题",
		Content: "测试内容",
		Status:  0,
		Members: []Member{
			{
				UserId:  0,
				IsAdmin: false,
			},
			{
				UserId:  1,
				IsAdmin: false,
			},
		},
		Tags: []Tag{
			{
				Content: "测试标签1",
			},
			{
				Content: "测试标签2",
			},
		},
	}
	_, err := cardRepo.Insert(context.Background(), card)
	assert.Nil(t, err)
}

func TestCardModelDefault_FindOne(t *testing.T) {
	card, err := cardRepo.FindOne(context.Background(), 13)
	assert.Nil(t, err)
	assert.NotNil(t, card)
}

func TestCardModelDefault_Update(t *testing.T) {
	card, err := cardRepo.FindOne(context.Background(), 13)
	assert.Nil(t, err)
	assert.NotNil(t, card)

	card.Title = "修改后的标题1"

	card.Members = []Member{
		{
			UserId:  444,
			IsAdmin: true,
		},
	}
	err = cardRepo.Update(context.Background(), card)
	assert.Nil(t, err)
}

func TestCardModelDefault_Delete(t *testing.T) {
	err := cardRepo.Delete(context.Background(), 11)
	assert.Nil(t, err)
}

func TestCardModelDefault_InsertTag(t *testing.T) {
	err := cardRepo.InsertTag(context.Background(), 13, "测试标签3")
	assert.Nil(t, err)
}

func TestCardModelDefault_DeleteTag(t *testing.T) {
	err := cardRepo.DeleteTag(context.Background(), 13, "测试标签3")
	assert.Nil(t, err)
}

func TestCardModelDefault_SetMember(t *testing.T) {
	err := cardRepo.SetMember(context.Background(), 13, 114514, false)
	assert.Nil(t, err)
}

func TestCardModelDefault_DeleteMember(t *testing.T) {
	err := cardRepo.DeleteMember(context.Background(), 13, 114514)
	assert.Nil(t, err)
}
