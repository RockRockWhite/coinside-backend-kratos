package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	api "github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/util"
	"net/http"
	"strconv"
)

type UserController struct {
	client api.UserClient
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := u.client.GetUserInfo(context.Background(), &api.GetUserInfoRequest{Id: id})

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
		resDto.Data = err
	} else {
		resDto.Data = res.Info
	}

	c.JSON(http.StatusOK, resDto)
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

func (u *UserController) GetUserId(c *gin.Context) {
	nickname := c.Query("nickname")
	if nickname != "" {
		res, err := u.client.GetUserInfoByNickname(context.Background(), &api.GetUserInfoByNicknameRequest{Nickname: nickname})

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
				Id: res.Info.Id,
			}
		}

		c.JSON(http.StatusOK, resDto)
		return
	}

	email := c.Query("email")
	if email != "" {
		res, err := u.client.GetUserInfoByEmail(context.Background(), &api.GetUserInfoByEmailRequest{Email: email})

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
				Id: res.Info.Id,
			}
		}

		c.JSON(http.StatusOK, resDto)
		return
	}

	mobile := c.Query("mobile")
	if mobile != "" {
		res, err := u.client.GetUserInfoByMobile(context.Background(), &api.GetUserInfoByMobileRequest{Mobile: mobile})

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
				Id: res.Info.Id,
			}
		}

		c.JSON(http.StatusOK, resDto)
		return
	}

	// 三个参数都未正确传入
	c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
}

func (u *UserController) SetFullname(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Fullname string `json:"fullname"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.client.SetFullname(context.Background(), &api.SetFullnameRequest{
		Id:       id,
		Fullname: reqDto.Fullname,
	})

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *UserController) SetAvatar(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Avatar string `json:"avatar"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.client.SetAvatar(context.Background(), &api.SetAvatarRequest{
		Id:     id,
		Avatar: reqDto.Avatar,
	})

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *UserController) SetConfig(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Config string `json:"config"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.client.SetConfig(context.Background(), &api.SetConfigRequest{
		Id:     id,
		Config: reqDto.Config,
	})

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *UserController) SetEmail(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Email      string `json:"email"`
		VerifyCode string `json:"verify_code"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.client.SetEmail(context.Background(), &api.SetEmailRequest{
		Id:         id,
		Email:      reqDto.Email,
		VerifyCode: reqDto.VerifyCode,
	})

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *UserController) SetMobile(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Mobile     string `json:"mobile"`
		VerifyCode string `json:"verify_code"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.client.SetMobile(context.Background(), &api.SetMobileRequest{
		Id:         id,
		Mobile:     reqDto.Mobile,
		VerifyCode: reqDto.VerifyCode,
	})

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *UserController) SetPassword(c *gin.Context) {
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := u.client.DeleteUser(context.Background(), &api.DeleteUserRequest{
		Id: id,
	})

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *UserController) Login(c *gin.Context) {
	// 获得登陆信息
	reqDto := struct {
		Id       uint64 `json:"id"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	// rpc请求登陆
	res, err := u.client.Login(context.Background(), &api.LoginRequest{
		Id:       reqDto.Id,
		Password: reqDto.Password,
	})

	resDto := dto.ResponseDto{
		Code:    dto.UserErrorCode[res.Code].Code,
		Message: dto.UserErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != api.Code_OK {
		resDto.Data = err
	} else {
		// 登陆成功，派发token
		token, err := util.GenerateJwtToken(
			&util.JwtClaims{Id: reqDto.Id})
		if err == nil {
			resDto.Data = struct {
				Token string `json:"token"`
			}{
				Token: token,
			}
		} else {
			resDto.Code = dto.UserErrorCode[api.Code_ERROR_UNKNOWN].Code
			resDto.Message = dto.UserErrorCode[api.Code_ERROR_UNKNOWN].Message
			resDto.Data = err
		}
	}

	c.JSON(http.StatusOK, resDto)
}

func NewUserController(client api.UserClient) *UserController {
	return &UserController{client: client}
}