package service

import (
	"context"
	api "github.com/ljxsteam/coinside-backend-kratos/api/card"
	"github.com/ljxsteam/coinside-backend-kratos/app/card/service/internal/data"
	"gorm.io/gorm"
)

type CardService struct {
	api.UnimplementedCardServiceServer

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

func (c CardService) CreateCardStream(server api.CardService_CreateCardStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardInfo(ctx context.Context, request *api.GetCardInfoRequest) (*api.GetCardInfoResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetCardInfoResponse{
			//: nil,
			Code: api.Code_ERROR_CARD_NOTFOUND,
		}, nil

	default:
		return &api.GetCardInfoResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	info := &api.CardInfo{
		Id:      one.Id,
		TeamId:  one.TeamId,
		Title:   one.Title,
		Content: one.Content,
		//todo:这个bug
		Status:    api.CardStatus(one.Status),
		Members:   one.Members,
		Tags:      one.Tags,
		CreatedAt: one.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: one.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return &api.GetCardInfoResponse{
		Card: info,
		Code: api.Code_OK,
	}, nil
}

func (c CardService) GetCardInfoStream(server api.CardService_GetCardInfoStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardList(ctx context.Context, request *api.GetCardListRequest) (*api.GetCardListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardService) GetCardListStream(server api.CardService_GetCardListStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) UpdateCard(ctx context.Context, request *api.UpdateCardRequest) (*api.UpdateCardResponse, error) {
	one, err := c.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.UpdateCardResponse{
			Code: api.Code_ERROR_CARD_NOTFOUND,
		}, nil
	default:
		return &api.UpdateCardResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	NewCard := data.Card{
		Id:      one.Id,
		TeamId:  one.TeamId,
		Title:   one.Title,
		Content: one.Content,
		Status:  one.Status,
	}

	if request.Title != "" && request.Title != NewCard.Title {
		NewCard.Title = request.Title
	}
	if request.Content != "" && request.Content != NewCard.Content {
		NewCard.Content = request.Content
	}

	if request.Status != 0 && request.Status != api.CardStatus(NewCard.Status) {
		NewCard.Status = data.CardStatus(request.Status)
	}

	if error := c.repo.Update(ctx, &NewCard); error != nil {
		return &api.UpdateCardResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.UpdateCardResponse{
		Code: api.Code_OK,
	}, nil

}

func (c CardService) UpdateCardStream(server api.CardService_UpdateCardStreamServer) error {
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

func (c CardService) DeleteCardStream(server api.CardService_DeleteCardStreamServer) error {
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

func (c CardService) AddCardTagStream(server api.CardService_AddCardTagStreamServer) error {
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

func (c CardService) DeleteCardTagStream(server api.CardService_DeleteCardTagStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) AddCardMember(ctx context.Context, request *api.AddCardMemberRequest) (*api.AddCardMemberResponse, error) {

	err := c.repo.SetMember(ctx, request.Id, request.UserId, request.IsAdmin)

	if err != nil {
		return &api.AddCardMemberResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddCardMemberResponse{Code: api.Code_OK}, nil
}

func (c CardService) AddCardMemberStream(server api.CardService_AddCardMemberStreamServer) error {
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

func (c CardService) DeleteCardMemberStream(server api.CardService_DeleteCardMemberStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (c CardService) mustEmbedUnimplementedCardServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewCardService(repo data.CardRepo) *CardService {
	return &CardService{
		repo: repo,
	}
}
