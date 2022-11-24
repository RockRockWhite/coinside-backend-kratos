package data

import (
	"context"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var todoRepo TodoRepo

func init() {
	todoRepo = NewTodoRepoNoCache(NewDB(config.NewConfig()))
}

func TestTodoRepoNoCache_Insert(t *testing.T) {
	todo := &Todo{
		CardId: 1,
		Title:  "单元测试数据",
	}

	_, err := todoRepo.Insert(context.Background(), todo)

	assert.Nil(t, err)
}

//
//func TestUserModelDefault_Insert(t *testing.T) {
//	team := &Team{
//		Name:        "xbmcjntd",
//		Description: "xbmzdcncncncncn",
//		Website:     "www.baidu.com",
//		Avatar:      "www.baudu.com/1.jpg",
//		Email:       "624707444@qq.com",
//		Members: []TeamMember{
//			{
//				UserId:  2,
//				IsAdmin: false,
//			},
//			{
//				UserId:  3,
//				IsAdmin: true,
//			},
//		},
//	}
//
//	_, err := teamRepo.Insert(context.Background(), team)
//	assert.Nil(t, err)
//}
//
//func TestUserModelDefault_FindOne(t *testing.T) {
//	team, err := teamRepo.FindOne(context.Background(), 5)
//	assert.Nil(t, err)
//
//	fmt.Println(team)
//}
//
//func TestUserModelDefault_Update(t *testing.T) {
//	team, err := teamRepo.FindOne(context.Background(), 1)
//	assert.Nil(t, err)
//
//	team.Avatar = "http://www.baidu.com/2.jpg"
//
//	err = teamRepo.Update(context.Background(), team)
//	assert.Nil(t, err)
//
//	fmt.Println(team)
//}
//
//func TestUserModelDefault_Delete(t *testing.T) {
//	err := teamRepo.Delete(context.Background(), 4)
//	assert.Nil(t, err)
//}
//
//func TestCardModelDefault_SetMember(t *testing.T) {
//	err := teamRepo.SetMember(context.Background(), 3, 2, false)
//	assert.Nil(t, err)
//}
//
//func TestCardModelDefault_DeleteMember(t *testing.T) {
//	err := teamRepo.DeleteMember(context.Background(), 1, 1)
//	assert.Nil(t, err)
//}
