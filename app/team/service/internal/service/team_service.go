package service

import (
	"context"
	api "github.com/ljxsteam/coinside-backend-kratos/api/team"
	"github.com/ljxsteam/coinside-backend-kratos/app/team/service/internal/data"
	"gorm.io/gorm"
	"strconv"
)

type TeamService struct {
	api.UnimplementedTeamServer

	repo data.TeamRepo
}

func (t TeamService) GetTeamById(ctx context.Context, request *api.GetTeamByIdRequest) (*api.GetTeamResponse, error) {
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

	var members []*api.TeamMember
	for _, m := range data.Members {
		members = append(members, &api.TeamMember{
			UserId:    m.UserId,
			IsAdmin:   m.IsAdmin,
			CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	team := &api.TeamInfo{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Website:     data.Website,
		Avatar:      data.Avatar,
		Email:       data.Email,
		Members:     members,
		CreatedAt:   data.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   data.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &api.GetTeamResponse{
		Team: team,
		Code: api.Code_OK,
	}, nil
}

func (t TeamService) GetTeamByIdStream(server api.Team_GetTeamByIdStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) GetTeamInfoList(ctx context.Context, request *api.GetTeamInfoListRequest) (*api.GetTeamInfoListResponse, error) {
	var filters []data.Filter

	// 生成过滤器参数
	for _, f := range request.Filters {
		switch f.Type {
		case api.TeamFilterType_USER_ADMIN:
			userId, err := strconv.ParseUint(f.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			filters = append(filters, data.NewUserFilter(userId, &data.FilterAdminOption{IsAdmin: true}))
		case api.TeamFilterType_USER_NO_ADMIN:
			userId, err := strconv.ParseUint(f.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			filters = append(filters, data.NewUserFilter(userId, &data.FilterAdminOption{IsAdmin: false}))
		case api.TeamFilterType_USER_ALL:
			userId, err := strconv.ParseUint(f.Value, 10, 64)
			if err != nil {
				return nil, err
			}
			filters = append(filters, data.NewUserFilter(userId, nil))
		default:
		}
	}

	all, count, err := t.repo.FindAll(ctx, request.Limit, request.Offset, filters)

	switch err {
	case nil:

	default:
		return &api.GetTeamInfoListResponse{
			Code:  api.Code_ERROR_UNKNOWN,
			Infos: nil,
		}, err
	}

	var infos []*api.TeamInfo

	// 组装members
	for _, one := range all {
		var members []*api.TeamMember
		for _, m := range one.Members {
			members = append(members, &api.TeamMember{
				UserId:    m.UserId,
				IsAdmin:   m.IsAdmin,
				CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: m.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}

		infos = append(infos, &api.TeamInfo{
			Id:          one.Id,
			Name:        one.Name,
			Description: one.Description,
			Website:     one.Website,
			Avatar:      one.Avatar,
			Email:       one.Email,
			Members:     members,
			CreatedAt:   one.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   one.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &api.GetTeamInfoListResponse{
		Code:  api.Code_OK,
		Count: count,
		Infos: infos,
	}, nil
}

func (t TeamService) GetTeamInfoListStream(server api.Team_GetTeamInfoListStreamServer) error {
	//TODO implement me
	panic("implement me")
}

func (t TeamService) AddTeam(ctx context.Context, team *api.AddTeamRequest) (*api.AddTeamResponse, error) {

	id, err := t.repo.Insert(ctx, &data.Team{
		Name:        team.Team.Name,
		Description: team.Team.Description,
		Website:     team.Team.Website,
		Avatar:      team.Team.Avatar,
		Email:       team.Team.Email,
		Members: []data.TeamMember{
			{
				UserId:  team.CreatorId,
				IsAdmin: true,
			},
		},
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

func (t TeamService) SetTeamName(ctx context.Context, req *api.SetTeamNameRequest) (*api.SetTeamNameResponse, error) {
	one, err := t.repo.FindOne(ctx, req.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetTeamNameResponse{
			Code: api.Code_ERROR_TEAM_NOTFOUND,
		}, nil
	default:
		return &api.SetTeamNameResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	one.Name = req.Name

	if error := t.repo.Update(ctx, one); error != nil {
		return &api.SetTeamNameResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.SetTeamNameResponse{
		Code: api.Code_OK,
	}, nil

}

func (t TeamService) SetTeamDescription(ctx context.Context, req *api.SetTeamDescriptionRequest) (*api.SetTeamDescriptionResponse, error) {
	one, err := t.repo.FindOne(ctx, req.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetTeamDescriptionResponse{
			Code: api.Code_ERROR_TEAM_NOTFOUND,
		}, nil
	default:
		return &api.SetTeamDescriptionResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}

	one.Description = req.Description

	if error := t.repo.Update(ctx, one); error != nil {
		return &api.SetTeamDescriptionResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.SetTeamDescriptionResponse{
		Code: api.Code_OK,
	}, nil

}
func (t TeamService) SetTeamAvatar(ctx context.Context, req *api.SetTeamAvatarRequest) (*api.SetTeamAvatarResponse, error) {
	one, err := t.repo.FindOne(ctx, req.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetTeamAvatarResponse{
			Code: api.Code_ERROR_TEAM_NOTFOUND,
		}, nil
	default:
		return &api.SetTeamAvatarResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}
	one.Avatar = req.Avatar

	if error := t.repo.Update(ctx, one); error != nil {
		return &api.SetTeamAvatarResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.SetTeamAvatarResponse{
		Code: api.Code_OK,
	}, nil

}
func (t TeamService) SetTeamEmail(ctx context.Context, req *api.SetTeamEmailRequest) (*api.SetTeamEmailResponse, error) {
	one, err := t.repo.FindOne(ctx, req.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetTeamEmailResponse{
			Code: api.Code_ERROR_TEAM_NOTFOUND,
		}, nil
	default:
		return &api.SetTeamEmailResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}
	one.Email = req.Email

	if error := t.repo.Update(ctx, one); error != nil {
		return &api.SetTeamEmailResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.SetTeamEmailResponse{
		Code: api.Code_OK,
	}, nil

}
func (t TeamService) SetTeamWebsite(ctx context.Context, req *api.SetTeamWebsiteRequest) (*api.SetTeamWebsiteResponse, error) {
	one, err := t.repo.FindOne(ctx, req.Id)

	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return &api.SetTeamWebsiteResponse{
			Code: api.Code_ERROR_TEAM_NOTFOUND,
		}, nil
	default:
		return &api.SetTeamWebsiteResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, err

	}
	one.Website = req.Website

	if error := t.repo.Update(ctx, one); error != nil {
		return &api.SetTeamWebsiteResponse{
			Code: api.Code_ERROR_UNKNOWN,
		}, error
	}

	return &api.SetTeamWebsiteResponse{
		Code: api.Code_OK,
	}, nil

}

func (t TeamService) SetTeamNameSteam(server api.Team_SetTeamNameSteamServer) error {
	//TODO implement me
	panic("implement me")
}
func (t TeamService) SetTeamDescriptionSteam(server api.Team_SetTeamDescriptionSteamServer) error {
	//TODO implement me
	panic("implement me")
}
func (t TeamService) SetTeamWebsiteSteam(server api.Team_SetTeamWebsiteSteamServer) error {
	//TODO implement me
	panic("implement me")
}
func (t TeamService) SetTeamAvatarSteam(server api.Team_SetTeamAvatarSteamServer) error {
	//TODO implement me
	panic("implement me")
}
func (t TeamService) SetTeamEmailSteam(server api.Team_SetTeamEmailSteamServer) error {
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

func (t TeamService) mustEmbedUnimplementedTeamServer() {
	//TODO implement me
	panic("implement me")
}

func NewTeamService(repo data.TeamRepo) *TeamService {
	return &TeamService{
		repo: repo,
	}
}
