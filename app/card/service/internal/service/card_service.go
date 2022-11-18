package service

import (
	"context"
	api "github.com/ljxsteam/coinside-backend-kratos/api/card"
	"github.com/ljxsteam/coinside-backend-kratos/app/card/service/internal/data"
	"gorm.io/gorm"
)

type CardService struct {
	api.UnimplementedCardServer

	repo data.CardRepo
}

func (c CardService) CreateCard(ctx context.Context, request *api.CreateCardRequest) (*api.CreateCardResponse, error) {
	id, err := c.repo.Insert(ctx, &data.Card{
		TeamId:  request.TeamId,
		Title:   request.Title,
		Content: request.Content,
		Status:  0,
	})

	if err != nil {
		return &api.CreateCardResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.CreateCardResponse{Id: id}, nil
}

func (c CardService) CreateCardStream(server api.Card_CreateCardStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardInfo(ctx context.Context, request *api.GetCardInfoRequest) (*api.GetCardInfoResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetCardInfoResponse{
			Code: api.Code_ERROR_CARD_NOTFOUND,
			Card: nil,
		}, nil

	default:
		return &api.GetCardInfoResponse{
			Code: api.Code_ERROR_UNKNOWN,
			Card: nil,
		}, err
	}

	var members []*api.CardMember
	for _, m := range one.Members {
		members = append(members, &api.CardMember{
			UserId:    m.UserId,
			IsAdmin:   m.IsAdmin,
			CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	var tags []*api.CardTag
	for _, t := range one.Tags {
		tags = append(tags, &api.CardTag{
			Content:   t.Content,
			CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: t.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	info := &api.CardInfo{
		Id:        one.Id,
		TeamId:    one.TeamId,
		Title:     one.Title,
		Content:   one.Content,
		Status:    one.Status,
		Members:   members,
		Tags:      tags,
		CreatedAt: one.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: one.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.GetCardInfoResponse{
		Card: info,
		Code: api.Code_OK,
	}, nil
}

func (c CardService) GetCardInfoStream(server api.Card_GetCardInfoStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardInfoList(ctx context.Context, request *api.GetCardInfoListRequest) (*api.GetCardInfoListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardInfoListStream(server api.Card_GetCardInfoListStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) UpdateCardTitle(ctx context.Context, request *api.UpdateCardTitleRequest) (*api.UpdateCardTitleResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.UpdateCardTitleResponse{
			Code: api.Code_ERROR_CARD_NOTFOUND,
		}, nil
	default:
		return &api.UpdateCardTitleResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Title = request.Title

	if err = c.repo.Update(ctx, one); err != nil {
		return &api.UpdateCardTitleResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.UpdateCardTitleResponse{
		Code: api.Code_OK,
	}, nil
}

func (c CardService) UpdateCardTitleStream(server api.Card_UpdateCardTitleStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) UpdateCardContent(ctx context.Context, request *api.UpdateCardContentRequest) (*api.UpdateCardContentResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.UpdateCardContentResponse{
			Code: api.Code_ERROR_CARD_NOTFOUND,
		}, nil
	default:
		return &api.UpdateCardContentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Content = request.Content

	if err = c.repo.Update(ctx, one); err != nil {
		return &api.UpdateCardContentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.UpdateCardContentResponse{
		Code: api.Code_OK,
	}, nil
}

func (c CardService) UpdateCardContentStream(server api.Card_UpdateCardContentStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) DeleteCard(ctx context.Context, request *api.DeleteCardRequest) (*api.DeleteCardResponse, error) {
	if err := c.repo.Delete(ctx, request.Id); err != nil {
		return &api.DeleteCardResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteCardResponse{
		Code: api.Code_OK,
	}, nil
}

func (c CardService) DeleteCardStream(server api.Card_DeleteCardStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) AddCardTag(ctx context.Context, request *api.AddCardTagRequest) (*api.AddCardTagResponse, error) {
	err := c.repo.InsertTag(ctx, request.Id, request.Content)

	if err != nil {
		return &api.AddCardTagResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddCardTagResponse{Code: api.Code_OK}, nil
}

func (c CardService) AddCardTagStream(server api.Card_AddCardTagStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) DeleteCardTag(ctx context.Context, request *api.DeleteCardTagRequest) (*api.DeleteCardTagResponse, error) {
	if err := c.repo.DeleteTag(ctx, request.Id, request.Content); err != nil {
		return &api.DeleteCardTagResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteCardTagResponse{
		Code: api.Code_OK,
	}, nil
}

func (c CardService) DeleteCardTagStream(server api.Card_DeleteCardTagStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) SetCardMember(ctx context.Context, request *api.SetCardMemberRequest) (*api.SetCardMemberResponse, error) {
	err := c.repo.SetMember(ctx, request.Id, request.UserId, request.IsAdmin)

	if err != nil {
		return &api.SetCardMemberResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.SetCardMemberResponse{Code: api.Code_OK}, nil
}

func (c CardService) SetCardMemberStream(server api.Card_SetCardMemberStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) DeleteCardMember(ctx context.Context, request *api.DeleteCardMemberRequest) (*api.DeleteCardMemberResponse, error) {
	if err := c.repo.DeleteMember(ctx, request.Id, request.UserId); err != nil {
		return &api.DeleteCardMemberResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteCardMemberResponse{
		Code: api.Code_OK,
	}, nil
}

func (c CardService) DeleteCardMemberStream(server api.Card_DeleteCardMemberStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) mustEmbedUnimplementedCardServer() {
	//TODO implement me
	panic("implement me")
}

func NewCardService(repo data.CardRepo) *CardService {
	return &CardService{
		repo: repo,
	}
}
