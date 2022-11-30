package data

import "context"

type AttachmentRepo interface {
	Insert(ctx context.Context, data *Attachment) (uint64, error)
	FindOne(ctx context.Context, id uint64) (*Attachment, error)
	//	FindByCardId(ctx context.Context, cardId uint64) (*[]Attachment, error)
	Update(ctx context.Context, newData *Attachment) error
	Delete(ctx context.Context, id uint64) error
}
