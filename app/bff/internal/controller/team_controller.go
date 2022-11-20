package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/team"
	"github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
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

	//resDto := dto.ResponseDto{
	//	Code:    dto.TeamErrorCode[res.Code].Code,
	//	Message: dto.TeamErrorCode[res.Code].Message,
	//	Data:    nil,
	//}
	//
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
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
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
			return
		}

		for _, m := range res.Team.Members {
			if err := stream.Send(&user.GetUserInfoRequest{Id: m.UserId}); err != nil {
				c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
				return
			}

			userInfo, err := stream.Recv()
			if err != nil {
				c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
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
	//
	//if res.Code != team.Code_OK {
	//	resDto.Data = err
	//} else {
	//	resDto.Data = res.Team
	//}
	//
	//c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) CreateTeam(c *gin.Context) {
	var req team.TeamInfo

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

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
