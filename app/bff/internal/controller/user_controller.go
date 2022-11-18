package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	api "github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"net/http"
	"strconv"
)

type UserController struct {
	client api.UserClient
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	user, err := u.client.GetUserInfo(context.Background(), &api.GetUserInfoRequest{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) CreateUser(c *gin.Context) {
	var req api.CreateUserRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.client.CreateUser(context.Background(), &req)

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
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

func (u *UserController) PatchUser(c *gin.Context) {

}

func (u *UserController) PatchPassword(c *gin.Context) {

}

func (u *UserController) DeleteUser(c *gin.Context) {

}

func NewUserController(client api.UserClient) *UserController {
	return &UserController{client: client}
}
