package data

import (
	"context"
	"fmt"
	"github.com/ljxsteam/coinside-backend-kratos/app/attachment/service/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var attachmentRepo AttachmentRepo

func init() {
	attachmentRepo = NewAttachmentRepoNoCache(NewDB(config.NewConfig()))
}

func TestUserModelDefault_Insert(t *testing.T) {
	attachment := &Attachment{
		CardId: 6,
		Link:   "https://www.baidu.com/",
	}
	_, err := attachmentRepo.Insert(context.Background(), attachment)
	assert.Nil(t, err)
}

func TestUserModelDefault_FindOne(t *testing.T) {
	attachment, err := attachmentRepo.FindOne(context.Background(), 1)
	assert.Nil(t, err)

	fmt.Println(attachment)
}

func TestUserModelDefault_Update(t *testing.T) {
	attachment, err := attachmentRepo.FindOne(context.Background(), 1)
	assert.Nil(t, err)

	attachment.Link = "https://www.baidu.com/"

	err = attachmentRepo.Update(context.Background(), attachment)
	assert.Nil(t, err)

	fmt.Println(attachment)
}

func TestUserModelDefault_Delete(t *testing.T) {
	err := attachmentRepo.Delete(context.Background(), 1)
	assert.Nil(t, err)
}
