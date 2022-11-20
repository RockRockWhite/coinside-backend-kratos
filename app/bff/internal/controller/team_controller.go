package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/team"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"net/http"
	"strconv"
)

type TeamController struct {
	client team.TeamClient
}

func (t *TeamController) GetTeamInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := t.client.GetTeamById(context.Background(), &team.GetTeamByIdRequest{Id: id})

	resDto := dto.ResponseDto{
		Code:    dto.TeamErrorCode[res.Code].Code,
		Message: dto.TeamErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != team.Code_OK {
		resDto.Data = err
	} else {
		resDto.Data = res.Team
	}

	c.JSON(http.StatusOK, resDto)
}

func (t *TeamController) CreateTeam(c *gin.Context) {
	var req team.TeamInfo

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := t.client.AddTeam(context.Background(), &req)

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

	res, err := t.client.SetTeamName(context.Background(), &team.SetTeamNameRequest{
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

	res, err := t.client.SetTeamDescription(context.Background(), &team.SetTeamDescriptionRequest{
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

	res, err := t.client.SetTeamWebsite(context.Background(), &team.SetTeamWebsiteRequest{
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

	res, err := t.client.SetTeamAvatar(context.Background(), &team.SetTeamAvatarRequest{
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

	res, err := t.client.SetTeamEmail(context.Background(), &team.SetTeamEmailRequest{
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

	res, err := t.client.DeleteTeam(context.Background(), &team.DeleteTeamRequest{
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

	res, err := t.client.AddMember(context.Background(), &team.AddMemberRequest{
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

	res, err := t.client.DeleteMember(context.Background(), &team.DeleteMemberRequest{
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

	res, err := t.client.AddAdmin(context.Background(), &team.AddAdminRequest{
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

func NewTeamController(client team.TeamClient) *TeamController {
	return &TeamController{client: client}
}
