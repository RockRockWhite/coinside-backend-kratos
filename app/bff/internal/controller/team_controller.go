package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/team"
	"github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/util"
	"net/http"
	"strconv"
)

type TeamController struct {
	userClient user.UserClient
	teamClient team.TeamClient
}

func (t *TeamController) GetTeamInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := t.teamClient.GetTeamById(context.Background(), &team.GetTeamByIdRequest{Id: id})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	}

	switch res.Code {
	case team.Code_OK:
		// 获取冗余用户信息
		type MemberInfo struct {
			*team.TeamMember
			Nickname string `json:"nickname"`
			Fullname string `json:"fullname"`
			Email    string `json:"email"`
			Avatar   string `json:"avatar"`
		}

		var members []MemberInfo
		// 获取成员信息
		stream, err := t.userClient.GetUserInfoStream(context.Background())
		defer stream.CloseSend()
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}

		for _, m := range res.Team.Members {
			if err := stream.Send(&user.GetUserInfoRequest{Id: m.UserId}); err != nil {
				c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
				return
			}

			userInfo, err := stream.Recv()
			if err != nil {
				c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
				return
			}

			members = append(members, MemberInfo{
				TeamMember: m,
				Nickname:   userInfo.Info.Nickname,
				Fullname:   userInfo.Info.Fullname,
				Email:      userInfo.Info.Email,
				Avatar:     userInfo.Info.Avatar,
			})
		}

		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.TeamErrorCode[res.Code].Code,
			Message: dto.TeamErrorCode[res.Code].Message,
			Data: struct {
				*team.TeamInfo
				Members []MemberInfo `json:"members"`
			}{
				TeamInfo: res.Team,
				Members:  members,
			},
		})

	default:
		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.TeamErrorCode[res.Code].Code,
			Message: dto.TeamErrorCode[res.Code].Message,
			Data:    nil,
		})
	}
}

func (t *TeamController) GetIsAdmin(c *gin.Context) {
	userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)
	teamId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := t.teamClient.GetIsAdmin(context.Background(), &team.GetIsAdminRequest{UserId: userId, TeamId: teamId})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	}

	switch res.Code {
	case team.Code_OK:

		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.TeamErrorCode[res.Code].Code,
			Message: dto.TeamErrorCode[res.Code].Message,
			Data: struct {
				IsAdmin bool `json:"is_admin"`
			}{
				IsAdmin: res.IsAdmin,
			},
		})

	default:
		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.TeamErrorCode[res.Code].Code,
			Message: dto.TeamErrorCode[res.Code].Message,
			Data:    nil,
		})
	}
}

func (t *TeamController) GetTeamInfoList(c *gin.Context) {
	userId := c.Query("user_id")
	claims := c.MustGet("claims").(*util.JwtClaims)
	if userId != fmt.Sprint(claims.Id) {
		c.JSON(http.StatusOK, dto.ErrorForbidden)
		return
	}
	isAdminQuery := c.Query("is_admin")
	limit, _ := strconv.ParseUint(c.Query("limit"), 10, 64)
	if limit == 0 {
		limit = 20
	}
	offset, _ := strconv.ParseUint(c.Query("offset"), 10, 64)

	// 生成过滤器参数
	var filters []*team.TeamFilter

	if isAdminQuery != "" {
		isAdmin, err := strconv.ParseBool(isAdminQuery)
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}

		switch isAdmin {
		case true:
			filters = append(filters, &team.TeamFilter{
				Type:  team.TeamFilterType_USER_ADMIN,
				Value: userId,
			})
		case false:
			filters = append(filters, &team.TeamFilter{
				Type:  team.TeamFilterType_USER_NO_ADMIN,
				Value: userId,
			})
		}
	} else {
		filters = append(filters, &team.TeamFilter{
			Type:  team.TeamFilterType_USER_ALL,
			Value: userId,
		})
	}

	res, err := t.teamClient.GetTeamInfoList(context.Background(), &team.GetTeamInfoListRequest{
		Limit:   limit,
		Offset:  offset,
		Filters: filters,
	})
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	}

	// 获取冗余用户信息
	type MemberInfo struct {
		*team.TeamMember
		Nickname string `json:"nickname"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Avatar   string `json:"avatar"`
	}
	var data []struct {
		*team.TeamInfo
		Members []MemberInfo `json:"members"`
	}

	infos := res.Infos
	for _, info := range infos {
		var members []MemberInfo
		// 获取成员信息
		stream, err := t.userClient.GetUserInfoStream(context.Background())
		defer stream.CloseSend()
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
			return
		}

		for _, m := range info.Members {
			if err := stream.Send(&user.GetUserInfoRequest{Id: m.UserId}); err != nil {
				c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
				return
			}

			userInfo, err := stream.Recv()
			if err != nil {
				c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
				return
			}

			members = append(members, MemberInfo{
				TeamMember: m,
				Nickname:   userInfo.Info.Nickname,
				Fullname:   userInfo.Info.Fullname,
				Email:      userInfo.Info.Email,
				Avatar:     userInfo.Info.Avatar,
			})
		}

		data = append(data, struct {
			*team.TeamInfo
			Members []MemberInfo `json:"members"`
		}{
			TeamInfo: info,
			Members:  members,
		})
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Code:    dto.TeamErrorCode[team.Code_OK].Code,
		Message: dto.TeamErrorCode[team.Code_OK].Message,
		Data: struct {
			Count uint64 `json:"count,omitempty"`
			Infos interface{}
		}{
			Count: res.Count,
			Infos: data,
		},
	})
}

func (t *TeamController) CreateTeam(c *gin.Context) {
	var req team.AddTeamRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}
	req.CreatorId = c.MustGet("claims").(*util.JwtClaims).Id

	res, err := t.teamClient.AddTeam(context.Background(), &req)

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
		resDto.Data = struct {
			Id uint64 `json:"id"`
		}{
			Id: res.Id,
		}
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) SetName(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Name string `json:"name"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.teamClient.SetTeamName(context.Background(), &team.SetTeamNameRequest{
		Id:   id,
		Name: reqDto.Name,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Website     string `json:"website"`
		Avatar      string `json:"avatar"`
		Email       string `json:"email"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.teamClient.UpdateTeam(context.Background(), &team.UpdateTeamRequest{
		Id:          id,
		Name:        reqDto.Name,
		Description: reqDto.Description,
		Website:     reqDto.Website,
		Avatar:      reqDto.Avatar,
		Email:       reqDto.Email,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) SetDescription(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Description string `json:"description"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.teamClient.SetTeamDescription(context.Background(), &team.SetTeamDescriptionRequest{
		Id:          id,
		Description: reqDto.Description,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) SetWebsite(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Website string `json:"website"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.teamClient.SetTeamWebsite(context.Background(), &team.SetTeamWebsiteRequest{
		Id:      id,
		Website: reqDto.Website,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) SetAvatar(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Avatar string `json:"avatar"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.teamClient.SetTeamAvatar(context.Background(), &team.SetTeamAvatarRequest{
		Id:     id,
		Avatar: reqDto.Avatar,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) SetEmail(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Email string `json:"email"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.teamClient.SetTeamEmail(context.Background(), &team.SetTeamEmailRequest{
		Id:    id,
		Email: reqDto.Email,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) DeleteTeam(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := t.teamClient.DeleteTeam(context.Background(), &team.DeleteTeamRequest{
		Id: id,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) SetTeamMember(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	reqDto := struct {
		IsAdmin bool `json:"is_admin"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	// 判断用户是否存在
	if res, err := t.userClient.GetUserInfo(context.Background(), &user.GetUserInfoRequest{Id: userId}); err != nil {
		// error
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err.Error()))
		return
	} else {
		// no error
		switch res.Code {
		case user.Code_OK:
		case user.Code_ERROR_USER_NOTFOUND:
			c.JSON(http.StatusOK, &dto.ResponseDto{
				Code:    dto.UserErrorCode[res.Code].Code,
				Message: dto.UserErrorCode[res.Code].Message,
				Data:    nil,
			})
			return
		default:
			c.JSON(http.StatusOK, &dto.ResponseDto{
				Code:    dto.UserErrorCode[user.Code_ERROR_UNKNOWN].Code,
				Message: dto.UserErrorCode[user.Code_ERROR_UNKNOWN].Message,
				Data:    err,
			})
			return
		}
	}

	res, err := t.teamClient.AddMember(context.Background(), &team.AddMemberRequest{
		TeamId:  id,
		UserId:  userId,
		IsAdmin: reqDto.IsAdmin,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) DeleteTeamMember(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	res, err := t.teamClient.DeleteMember(context.Background(), &team.DeleteMemberRequest{
		TeamId: id,
		UserId: userId,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) SetTeamAdmin(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	reqDto := struct {
		IsAdmin bool `json:"is_admin"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.teamClient.AddAdmin(context.Background(), &team.AddAdminRequest{
		TeamId: id,
		UserId: userId,
	})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func NewTeamController(userClient user.UserClient, client team.TeamClient) *TeamController {
	return &TeamController{userClient: userClient, teamClient: client}
}
