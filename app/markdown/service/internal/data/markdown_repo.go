package data

import "context"

type MarkdownRepo interface {
	Insert(ctx context.Context, data *Markdown) (uint64, error)
	FindOne(ctx context.Context, id uint64) (*Markdown, error)
	//	FindByCardId(ctx context.Context, cardId uint64) (*[]Markdown, error)
	Update(ctx context.Context, newData *Markdown) error
	Delete(ctx context.Context, id uint64) error
}
