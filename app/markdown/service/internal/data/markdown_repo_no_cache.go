package data

import (
	"context"
	"gorm.io/gorm"
)

type MarkdownRepoNoCache struct {
	db *gorm.DB
}

func (m MarkdownRepoNoCache) Insert(ctx context.Context, data *Markdown) (uint64, error) {
	res := m.db.Create(data)
	return data.Id, res.Error
}

func (m MarkdownRepoNoCache) FindOne(ctx context.Context, id uint64) (*Markdown, error) {
	var markdown Markdown
	res := m.db.Model(&markdown).Where("id = ?", id).First(&markdown)

	return &markdown, res.Error
}

func (m MarkdownRepoNoCache) Update(ctx context.Context, newData *Markdown) error {
	res := m.db.Save(newData)
	return res.Error
}

func (m MarkdownRepoNoCache) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	res := m.db.Delete(&data)
	return res.Error
}

func NewMarkdownRepoNoCache(db *gorm.DB) MarkdownRepo {
	return MarkdownRepoNoCache{db: db}
}
