package data

import (
	"context"
	"gorm.io/gorm"
)

type AttachmentRepoNoCache struct {
	db *gorm.DB
}

func (a AttachmentRepoNoCache) Insert(ctx context.Context, data *Attachment) (uint64, error) {
	res := a.db.Create(data)
	return data.Id, res.Error
}

func (a AttachmentRepoNoCache) FindOne(ctx context.Context, id uint64) (*Attachment, error) {
	var attachment Attachment
	res := a.db.Model(&attachment).Where("id = ?", id).First(&attachment)

	return &attachment, res.Error
}

func (a AttachmentRepoNoCache) Update(ctx context.Context, newData *Attachment) error {
	res := a.db.Save(newData)
	return res.Error
}

func (a AttachmentRepoNoCache) Delete(ctx context.Context, id uint64) error {
	data, err := a.FindOne(ctx, id)
	if err != nil {
		return err
	}

	res := a.db.Delete(&data)
	return res.Error
}

func NewAttachmentRepoNoCache(db *gorm.DB) AttachmentRepo {
	return AttachmentRepoNoCache{db: db}
}
