package data

import (
	"context"
	"fmt"
	"github.com/ljxsteam/coinside-backend-kratos/app/markdown/service/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var markdownRepo MarkdownRepo

func init() {
	markdownRepo = NewMarkdownRepoNoCache(NewDB(config.NewConfig()))
}

func TestUserModelDefault_Insert(t *testing.T) {
	markdown := &Markdown{
		CardId:  6,
		Content: "# zzzz",
	}
	_, err := markdownRepo.Insert(context.Background(), markdown)
	assert.Nil(t, err)
}

func TestUserModelDefault_FindOne(t *testing.T) {
	markdown, err := markdownRepo.FindOne(context.Background(), 1)
	assert.Nil(t, err)

	fmt.Println(markdown)
}

func TestUserModelDefault_Update(t *testing.T) {
	markdown, err := markdownRepo.FindOne(context.Background(), 1)
	assert.Nil(t, err)

	markdown.Content = "# aaaaaaaaaaaaaxbwl"

	err = markdownRepo.Update(context.Background(), markdown)
	assert.Nil(t, err)

	fmt.Println(markdown)
}

func TestUserModelDefault_Delete(t *testing.T) {
	err := markdownRepo.Delete(context.Background(), 1)
	assert.Nil(t, err)
}
