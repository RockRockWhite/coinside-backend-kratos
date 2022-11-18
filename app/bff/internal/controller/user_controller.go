package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	api "github.com/ljxsteam/coinside-backend-kratos/api/user"
	"net/http"
	"strconv"
)

type UserController struct {
	client api.UserClient
}

func (u *UserController) GetUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	user, err := u.client.GetUserInfo(context.Background(), &api.GetUserInfoRequest{Id: id})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) AddUser(c *gin.Context) {

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
