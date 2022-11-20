package service

import (
	"context"
	"github.com/ljxsteam/coinside-backend-kratos/api/card"
	"github.com/ljxsteam/coinside-backend-kratos/app/card/service/internal/data"
	"gorm.io/gorm"
	"time"
)

type CardService struct {
	card.UnimplementedCardServer

	repo data.CardRepo
}

func (c CardService) CreateCard(ctx context.Context, request *card.CreateCardRequest) (*card.CreateCardResponse, error) {
	id, err := c.repo.Insert(ctx, &data.Card{
		TeamId:  request.TeamId,
		Title:   request.Title,
		Content: request.Content,
		Status:  0,
	})

	if err != nil {
		return &card.CreateCardResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	return &card.CreateCardResponse{Id: id}, nil
}

func (c CardService) CreateCardStream(server card.Card_CreateCardStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardInfo(ctx context.Context, request *card.GetCardInfoRequest) (*card.GetCardInfoResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &card.GetCardInfoResponse{
			Code: card.Code_ERROR_CARD_NOTFOUND,
			Info: nil,
		}, nil

	default:
		return &card.GetCardInfoResponse{
			Code: card.Code_ERROR_UNKNOWN,
			Info: nil,
		}, err
	}

	var members []*card.CardMember
	for _, m := range one.Members {
		members = append(members, &card.CardMember{
			UserId:    m.UserId,
			IsAdmin:   m.IsAdmin,
			CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	var tags []*card.CardTag
	for _, t := range one.Tags {
		tags = append(tags, &card.CardTag{
			Content:   t.Content,
			CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: t.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	info := &card.CardInfo{
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

	return &card.GetCardInfoResponse{
		Info: info,
		Code: card.Code_OK,
	}, nil
}

func (c CardService) GetCardInfoStream(server card.Card_GetCardInfoStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardInfoList(ctx context.Context, request *card.GetCardInfoListRequest) (*card.GetCardInfoListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardInfoListStream(server card.Card_GetCardInfoListStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) UpdateCardTitle(ctx context.Context, request *card.UpdateCardTitleRequest) (*card.UpdateCardTitleResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &card.UpdateCardTitleResponse{
			Code: card.Code_ERROR_CARD_NOTFOUND,
		}, nil
	default:
		return &card.UpdateCardTitleResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Title = request.Title

	if err = c.repo.Update(ctx, one); err != nil {
		return &card.UpdateCardTitleResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	return &card.UpdateCardTitleResponse{
		Code: card.Code_OK,
	}, nil
}

func (c CardService) UpdateCardTitleStream(server card.Card_UpdateCardTitleStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) UpdateCardContent(ctx context.Context, request *card.UpdateCardContentRequest) (*card.UpdateCardContentResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &card.UpdateCardContentResponse{
			Code: card.Code_ERROR_CARD_NOTFOUND,
		}, nil
	default:
		return &card.UpdateCardContentResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Content = request.Content

	if err = c.repo.Update(ctx, one); err != nil {
		return &card.UpdateCardContentResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	return &card.UpdateCardContentResponse{
		Code: card.Code_OK,
	}, nil
}

func (c CardService) UpdateCardContentStream(server card.Card_UpdateCardContentStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) DeleteCard(ctx context.Context, request *card.DeleteCardRequest) (*card.DeleteCardResponse, error) {
	if err := c.repo.Delete(ctx, request.Id); err != nil {
		return &card.DeleteCardResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}
	return &card.DeleteCardResponse{
		Code: card.Code_OK,
	}, nil
}

func (c CardService) DeleteCardStream(server card.Card_DeleteCardStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) SetCardDeadline(ctx context.Context, request *card.SetCardDeadlineRequest) (*card.SetCardDeadlineResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &card.SetCardDeadlineResponse{
			Code: card.Code_ERROR_CARD_NOTFOUND,
		}, nil
	default:
		return &card.SetCardDeadlineResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Deadline, err = time.Parse("2006-01-02 15:04:05", request.Deadline)
	if err != nil {
		return &card.SetCardDeadlineResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	if err = c.repo.Update(ctx, one); err != nil {
		return &card.SetCardDeadlineResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	return &card.SetCardDeadlineResponse{
		Code: card.Code_OK,
	}, nil
}

func (c CardService) SetCardDeadlineStream(server card.Card_SetCardDeadlineStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) SetCardStatus(ctx context.Context, request *card.SetCardStatusRequest) (*card.SetCardStatusResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &card.SetCardStatusResponse{
			Code: card.Code_ERROR_CARD_NOTFOUND,
		}, nil
	default:
		return &card.SetCardStatusResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	one.Status = request.Status

	if err = c.repo.Update(ctx, one); err != nil {
		return &card.SetCardStatusResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	return &card.SetCardStatusResponse{
		Code: card.Code_OK,
	}, nil
}

func (c CardService) SetCardStatusStream(server card.Card_SetCardStatusStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) AddCardTag(ctx context.Context, request *card.AddCardTagRequest) (*card.AddCardTagResponse, error) {
	err := c.repo.InsertTag(ctx, request.Id, request.Content)

	if err != nil {
		return &card.AddCardTagResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	return &card.AddCardTagResponse{Code: card.Code_OK}, nil
}

func (c CardService) AddCardTagStream(server card.Card_AddCardTagStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) DeleteCardTag(ctx context.Context, request *card.DeleteCardTagRequest) (*card.DeleteCardTagResponse, error) {
	if err := c.repo.DeleteTag(ctx, request.Id, request.Content); err != nil {
		return &card.DeleteCardTagResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}
	return &card.DeleteCardTagResponse{
		Code: card.Code_OK,
	}, nil
}

func (c CardService) DeleteCardTagStream(server card.Card_DeleteCardTagStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) SetCardMember(ctx context.Context, request *card.SetCardMemberRequest) (*card.SetCardMemberResponse, error) {
	err := c.repo.SetMember(ctx, request.Id, request.UserId, request.IsAdmin)

	if err != nil {
		return &card.SetCardMemberResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}

	return &card.SetCardMemberResponse{Code: card.Code_OK}, nil
}

func (c CardService) SetCardMemberStream(server card.Card_SetCardMemberStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) DeleteCardMember(ctx context.Context, request *card.DeleteCardMemberRequest) (*card.DeleteCardMemberResponse, error) {
	if err := c.repo.DeleteMember(ctx, request.Id, request.UserId); err != nil {
		return &card.DeleteCardMemberResponse{
			Code: card.Code_ERROR_UNKNOWN,
		}, err
	}
	return &card.DeleteCardMemberResponse{
		Code: card.Code_OK,
	}, nil
}

func (c CardService) DeleteCardMemberStream(server card.Card_DeleteCardMemberStreamServer) error {
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
