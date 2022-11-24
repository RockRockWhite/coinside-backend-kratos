package service

import (
	"context"
	"github.com/ljxsteam/coinside-backend-kratos/api/vote"
	"github.com/ljxsteam/coinside-backend-kratos/app/vote/service/internal/data"
)

type VoteService struct {
	vote.UnimplementedVoteServer
	repo data.VoteRepo
}

func (u VoteService) GetvoteById(ctx context.Context, request *vote.GetVoteByIdRequest) (*vote.GetVoteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) GetVoteByIdStream(server vote.Vote_GetVoteByIdStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) AddVote(ctx context.Context, info *vote.VoteInfo) (*vote.AddVoteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) AddVoteStream(server vote.Vote_AddVoteStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) SetVoteTitle(ctx context.Context, request *vote.SetVoteTitleRequest) (*vote.SetVoteTitleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) SetVoteTitleStream(request *vote.SetVoteTitleRequest, server vote.Vote_SetVoteTitleStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) DeleteVote(ctx context.Context, request *vote.DeleteVoteRequest) (*vote.DeleteVoteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) DeleteVoteStream(server vote.Vote_DeleteVoteStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) AddItem(ctx context.Context, item *vote.VoteItem) (*vote.AddItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) AddItemStream(server vote.Vote_AddItemStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) SetItemContent(ctx context.Context, request *vote.SetContentRequest) (*vote.SetContentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) SetItemContentStream(server vote.Vote_SetItemContentStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) DeleteVoteItem(ctx context.Context, request *vote.DeleteVoteItemRequest) (*vote.DeleteVoteItemResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) DeleteVoteItemStream(server vote.Vote_DeleteVoteItemStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (u VoteService) AddCommit(ctx context.Context, commit *vote.VoteItemCommit) (*vote.AddCommitResponse, error) {
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
