package service

import (
	"context"
	"github.com/ljxsteam/coinside-backend-kratos/api/vote"
	api "github.com/ljxsteam/coinside-backend-kratos/api/vote"
	"github.com/ljxsteam/coinside-backend-kratos/app/vote/service/internal/data"
	"gorm.io/gorm"
)

type VoteService struct {
	vote.UnimplementedVoteServer
	repo data.VoteRepo
}

func (u VoteService) GetVoteById(ctx context.Context, request *vote.GetVoteByIdRequest) (*vote.GetVoteResponse, error) {
	data, err := u.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetVoteResponse{
			Vote: nil,
			Code: api.Code_ERROR_VOTE_NOTFOUND,
		}, nil
	default:
		return &api.GetVoteResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}
	//var commits []*api.VoteItemCommit
	//for _, m := range data.Items {
	//	commits = append(commits, &api.VoteItemCommit{
	//		Id:         m.Id,
	//		VoteItemId: m.
	//		UserId:     m.Content,
	//		CreatedAt:  m.CreatedAt.Format("2006-01-02 15:04:05"),
	//		UpdatedAt:  m.UpdatedAt.Format("2006-01-02 15:04:05"),
	//	})
	//}

	var items []*api.VoteItem
	for _, m := range data.Items {
		items = append(items, &api.VoteItem{
			Id:      m.Id,
			VoteId:  m.VoteId,
			Content: m.Content,
			//	Commits:   commits,
			CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	vote := &api.VoteInfo{
		Id:        data.Id,
		CardId:    data.CardId,
		Title:     data.Title,
		Items:     items,
		CreatedAt: data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &api.GetVoteResponse{
		Vote: vote,
		Code: api.Code_OK,
	}, nil
}

func (u VoteService) GetVoteByIdStream(server vote.Vote_GetVoteByIdStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) AddVote(ctx context.Context, vote *vote.AddVoteRequest) (*vote.AddVoteResponse, error) {
	id, err := u.repo.Insert(ctx, &data.Vote{
		CardId: vote.CardId,
		Title:  vote.Title,
	})

	if err != nil {
		return &api.AddVoteResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddVoteResponse{Id: id}, nil
}

func (u VoteService) AddVoteStream(server vote.Vote_AddVoteStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) SetVoteTitle(ctx context.Context, req *vote.SetVoteTitleRequest) (*vote.SetVoteTitleResponse, error) {
	one, err := u.repo.FindOne(ctx, req.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetVoteTitleResponse{
			Code: api.Code_ERROR_VOTE_NOTFOUND,
		}, nil
	default:
		return &api.SetVoteTitleResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	one.Title = req.Title

	if error := u.repo.Update(ctx, one); error != nil {
		return &api.SetVoteTitleResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.SetVoteTitleResponse{
		Code: api.Code_OK,
	}, nil

}

func (u VoteService) SetVoteTitleStream(request *vote.SetVoteTitleRequest, server vote.Vote_SetVoteTitleStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) DeleteVote(ctx context.Context, request *vote.DeleteVoteRequest) (*vote.DeleteVoteResponse, error) {
	if err := u.repo.Delete(ctx, request.Id); err != nil {
		return &api.DeleteVoteResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteVoteResponse{
		Code: api.Code_OK,
	}, nil
}

func (u VoteService) DeleteVoteStream(server vote.Vote_DeleteVoteStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) AddItem(ctx context.Context, item *vote.AddItemRequest) (*vote.AddItemResponse, error) {
	id, err := u.repo.InsertItem(ctx, item.VoteId, item.Content)

	if err != nil {
		return &api.AddItemResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddItemResponse{Code: api.Code_OK, Id: id}, nil
}

func (u VoteService) AddItemStream(server vote.Vote_AddItemStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) SetItemContent(ctx context.Context, req *vote.SetContentRequest) (*vote.SetContentResponse, error) {
	_, err := u.repo.FindOne(ctx, req.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetContentResponse{
			Code: api.Code_ERROR_VOTE_NOTFOUND,
		}, nil
	default:
		return &api.SetContentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	if error := u.repo.UpdateItem(ctx, req.Id, req.ItemId, req.Content); error != nil {
		return &api.SetContentResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.SetContentResponse{
		Code: api.Code_OK,
	}, nil
}

func (u VoteService) SetItemContentStream(server vote.Vote_SetItemContentStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) DeleteVoteItem(ctx context.Context, request *vote.DeleteVoteItemRequest) (*vote.DeleteVoteItemResponse, error) {
	if err := u.repo.DeleteItem(ctx, request.VoteId, request.VoteItemId); err != nil {
		return &api.DeleteVoteItemResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteVoteItemResponse{
		Code: api.Code_OK,
	}, nil
}

func (u VoteService) DeleteVoteItemStream(server vote.Vote_DeleteVoteItemStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) AddCommit(ctx context.Context, commit *vote.AddCommitRequest) (*vote.AddCommitResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) DeleteCommit(ctx context.Context, request *vote.DeleteCommitRequest) (*vote.DeleteCommitResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) mustEmbedUnimplementedVoteServer() {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) mustEmbedUnimplementedCardServer() {
	//TODO implement me
	panic("implement me")
}

func NewVoteService(repo data.VoteRepo) *VoteService {
	return &VoteService{
		repo: repo,
	}
}
