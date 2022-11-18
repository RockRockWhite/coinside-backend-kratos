package service

import (
	"context"
	api "github.com/ljxsteam/coinside-backend-kratos/api/team"
	"github.com/ljxsteam/coinside-backend-kratos/app/team/service/internal/data"
	"gorm.io/gorm"
)

type TeamService struct {
	api.UnimplementedTeamServer

	repo data.TeamRepo
}

func (t TeamService) GetTeamByID(ctx context.Context, request *api.GetTeamByIdRequest) (*api.GetTeamResponse, error) {
	data, err := t.repo.FindOne(ctx, request.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.GetTeamResponse{
			Team: nil,
			Code: api.Code_ERROR_TEAM_NOTFOUND,
		}, nil
	default:
		return &api.GetTeamResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	team := &api.TeamInfo{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Website:     data.Website,
		Avatar:      data.Avatar,
		Email:       data.Email,
		CreatedAt:   data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &api.GetTeamResponse{
		Team: team,
		Code: api.Code_OK,
	}, nil
}

func (t TeamService) GetTeamByIDStream(server api.Team_GetTeamByIDStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) GetTeamsByName(ctx context.Context, request *api.GetTeamsByNameRequest) (*api.GetTeamsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) GetTeamsByNameStream(server api.Team_GetTeamsByNameStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) AddTeam(ctx context.Context, team *api.TeamInfo) (*api.AddTeamResponse, error) {

	id, err := t.repo.Insert(ctx, &data.Team{
		Name:        team.Name,
		Description: team.Description,
		Website:     team.Website,
		Avatar:      team.Avatar,
		Email:       team.Email,
	})

	if err != nil {
		return &api.AddTeamResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddTeamResponse{Id: id}, nil
}

func (t TeamService) AddTeamStream(server api.Team_AddTeamStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) UpdateTeam(ctx context.Context, team *api.TeamInfo) (*api.UpdateTeamResponse, error) {
	one, err := t.repo.FindOne(ctx, team.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.UpdateTeamResponse{
			Code: api.Code_ERROR_TEAM_NOTFOUND,
		}, nil
	default:
		return &api.UpdateTeamResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	NewTeam := data.Team{
		Id:          one.Id,
		Name:        one.Name,
		Description: one.Description,
		Website:     one.Website,
		Avatar:      one.Avatar,
		Email:       one.Email,
	}

	if team.Name != "" && team.Name != NewTeam.Name {
		NewTeam.Name = team.Name
	}
	if team.Description != "" && team.Description != NewTeam.Description {
		NewTeam.Description = team.Description
	}
	if team.Website != "" && team.Website != NewTeam.Website {
		NewTeam.Website = team.Website
	}
	if team.Avatar != "" && team.Avatar != NewTeam.Avatar {
		NewTeam.Avatar = team.Avatar
	}
	if team.Email != "" && team.Email != NewTeam.Email {
		NewTeam.Email = team.Email
	}

	if error := t.repo.Update(ctx, &NewTeam); error != nil {
		return &api.UpdateTeamResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.UpdateTeamResponse{
		Code: api.Code_OK,
	}, nil

}

func (t TeamService) UpdateTeamSteam(team *api.TeamInfo, server api.Team_UpdateTeamSteamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) DeleteTeam(ctx context.Context, request *api.DeleteTeamRequest) (*api.DeleteTeamResponse, error) {
	if err := t.repo.Delete(ctx, request.Id); err != nil {
		return &api.DeleteTeamResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteTeamResponse{
		Code: api.Code_OK,
	}, nil
}

func (t TeamService) DeleteTeamStream(server api.Team_DeleteTeamStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) AddMember(ctx context.Context, request *api.AddMemberRequest) (*api.AddMemberResponse, error) {
	if err := t.repo.SetMember(ctx, request.TeamId, request.UserId, request.IsAdmin); err != nil {
		return &api.AddMemberResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddMemberResponse{
		Code: api.Code_OK,
	}, nil

}

func (t TeamService) AddMemberStream(server api.Team_AddMemberStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) DeleteMember(ctx context.Context, request *api.DeleteMemberRequest) (*api.DeleteMemberResponse, error) {
	if err := t.repo.DeleteMember(ctx, request.TeamId, request.UserId); err != nil {
		return &api.DeleteMemberResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}
	return &api.DeleteMemberResponse{
		Code: api.Code_OK,
	}, nil
}

func (t TeamService) DeleteMemberStream(server api.Team_DeleteMemberStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) AddAdmin(ctx context.Context, request *api.AddAdminRequest) (*api.AddAdminResponse, error) {
	if err := t.repo.SetMember(ctx, request.TeamId, request.UserId, true); err != nil {
		return &api.AddAdminResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err
	}

	return &api.AddAdminResponse{
		Code: api.Code_OK,
	}, nil

}

func (t TeamService) AddAdminStream(server api.Team_AddAdminStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) mustEmbedUnimplementedTeamServiceServer() {
	//TODO implement me
	panic("implement me")
}

func NewTeamService(repo data.TeamRepo) *TeamService {
	return &TeamService{
		repo: repo,
	}
}
