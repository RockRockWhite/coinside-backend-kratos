package service

import (
	"context"
	api "github.com/ljxsteam/coinside-backend-kratos/api/markdown"
	"github.com/ljxsteam/coinside-backend-kratos/app/markdown/service/internal/data"
	"gorm.io/gorm"
)

type MarkdownService struct {
	api.UnimplementedMarkdownServer

	repo data.MarkdownRepo
}

func (m MarkdownService) GetMarkdownById(ctx context.Context, request *api.GetMarkdownByIdRequest) (*api.GetMarkdownResponse, error) {
	data, err := m.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetMarkdownResponse{
			Markdown: nil,
			Code:     api.Code_ERROR_MARKDOWN_NOTFOUND,
		}, nil
	default:
		return &api.GetMarkdownResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	markdown := &api.MarkdownInfo{
		Id:        data.Id,
		CardId:    data.CardId,
		Content:   data.Content,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &api.GetMarkdownResponse{
		Markdown: markdown,
		Code:     api.Code_OK,
	}, nil
}

func (m MarkdownService) GetMarkdownByIdStream(server api.Markdown_GetMarkdownByIdStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (m MarkdownService) GetMarkdownsByCardId(ctx context.Context, request *api.GetMarkdownsByCardIdRequest) (*api.GetMarkdownsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MarkdownService) GetMarkdownsByCardIdStream(server api.Markdown_GetMarkdownsByCardIdStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (m MarkdownService) AddMarkdown(ctx context.Context, markdown *api.MarkdownInfo) (*api.AddMarkdownResponse, error) {
	id, err := m.repo.Insert(ctx, &data.Markdown{
		CardId:  markdown.CardId,
		Content: markdown.Content,
	})

	if err != nil {
		return &api.AddMarkdownResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddMarkdownResponse{Id: id}, nil
}

func (m MarkdownService) AddMarkdownStream(server api.Markdown_AddMarkdownStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (m MarkdownService) UpdateMarkdown(ctx context.Context, markdown *api.MarkdownInfo) (*api.UpdateMarkdownResponse, error) {
	one, err := m.repo.FindOne(ctx, markdown.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.UpdateMarkdownResponse{
			Code: api.Code_ERROR_MARKDOWN_NOTFOUND,
		}, nil
	default:
		return &api.UpdateMarkdownResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	NewMarkdown := data.Markdown{
		Id:        one.Id,
		CardId:    one.CardId,
		Content:   one.Content,
		CreatedAt: one.CreatedAt,
	}

	if markdown.CardId != 0 && markdown.CardId != NewMarkdown.CardId {
		NewMarkdown.CardId = markdown.CardId
	}

	if markdown.Content != "" && markdown.Content != NewMarkdown.Content {
		NewMarkdown.Content = markdown.Content
	}

	if error := m.repo.Update(ctx, &NewMarkdown); error != nil {
		return &api.UpdateMarkdownResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.UpdateMarkdownResponse{
		Code: api.Code_OK,
	}, nil

}

func (m MarkdownService) UpdateMarkdownStream(info *api.MarkdownInfo, server api.Markdown_UpdateMarkdownStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (m MarkdownService) DeleteMarkdown(ctx context.Context, request *api.DeleteMarkdownRequest) (*api.DeleteMarkdownResponse, error) {
	if err := m.repo.Delete(ctx, request.Id); err != nil {
		return &api.DeleteMarkdownResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteMarkdownResponse{
		Code: api.Code_OK,
	}, nil
}

func (m MarkdownService) DeleteMarkdownStream(server api.Markdown_DeleteMarkdownStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (m MarkdownService) mustEmbedUnimplementedMarkdownServer() {
	//TODO implement me
	panic("implement me")
}

func NewMarkdownService(repo data.MarkdownRepo) *MarkdownService {
	return &MarkdownService{
		repo: repo,
	}
}
