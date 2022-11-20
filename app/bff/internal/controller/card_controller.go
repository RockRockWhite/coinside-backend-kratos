package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/card"
	"github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"net/http"
	"strconv"
)

type CardController struct {
	userClient user.UserClient
	cardClient card.CardClient
}

func (u *CardController) GetCardInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := u.cardClient.GetCardInfo(context.Background(), &card.GetCardInfoRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	switch res.Code {
	case card.Code_OK:
		// 获取冗余用户信息
		type MemberInfo struct {
			*card.CardMember
			Nickname string `json:"nickname"`
			Fullname string `json:"fullname"`
			Email    string `json:"email"`
			Avatar   string `json:"avatar"`
		}

		var members []MemberInfo
		// 获取成员信息
		stream, err := u.userClient.GetUserInfoStream(context.Background())
		defer stream.CloseSend()
		if err != nil {
			c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
			return
		}

		for _, m := range res.Info.Members {
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
				CardMember: m,
				Nickname:   userInfo.Info.Nickname,
				Fullname:   userInfo.Info.Fullname,
				Email:      userInfo.Info.Email,
				Avatar:     userInfo.Info.Avatar,
			})
		}

		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.CardErrorCode[res.Code].Code,
			Message: dto.CardErrorCode[res.Code].Message,
			Data: struct {
				*card.CardInfo
				Members []MemberInfo `json:"members"`
			}{
				CardInfo: res.Info,
				Members:  members,
			},
		})

	default:
		c.JSON(http.StatusOK, &dto.ResponseDto{
			Code:    dto.CardErrorCode[res.Code].Code,
			Message: dto.CardErrorCode[res.Code].Message,
			Data:    nil,
		})
	}
}

func (u *CardController) CreateCard(c *gin.Context) {
	var req card.CreateCardRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.CreateCard(context.Background(), &req)

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
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

func (u *CardController) SetTitle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Title string `json:"title"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.UpdateCardTitle(context.Background(), &card.UpdateCardTitleRequest{
		Id:    id,
		Title: reqDto.Title,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) SetContent(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Content string `json:"content"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.UpdateCardContent(context.Background(), &card.UpdateCardContentRequest{
		Id:      id,
		Content: reqDto.Content,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) SetDeadline(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Deadline string `json:"deadline"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.SetCardDeadline(context.Background(), &card.SetCardDeadlineRequest{
		Id:       id,
		Deadline: reqDto.Deadline,
	})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	})
}

func (u *CardController) SetStatus(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Status card.CardStatus `json:"status"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.cardClient.SetCardStatus(context.Background(), &card.SetCardStatusRequest{
		Id:     id,
		Status: reqDto.Status,
	})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	c.JSON(http.StatusOK, &dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	})
}

func (u *CardController) AddTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	tagContent := c.Param("tag_content")

	res, err := u.cardClient.AddCardTag(context.Background(), &card.AddCardTagRequest{
		Id:      id,
		Content: tagContent,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) DeleleTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	tagContent := c.Param("tag_content")

	res, err := u.cardClient.DeleteCardTag(context.Background(), &card.DeleteCardTagRequest{
		Id:      id,
		Content: tagContent,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) SetMember(c *gin.Context) {
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
	if res, err := u.userClient.GetUserInfo(context.Background(), &user.GetUserInfoRequest{Id: userId}); err != nil {
		// error
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
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

	// 设置团队成员
	res, err := u.cardClient.SetCardMember(context.Background(), &card.SetCardMemberRequest{
		Id:      id,
		UserId:  userId,
		IsAdmin: reqDto.IsAdmin,
	})

	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	c.JSON(http.StatusOK, dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	})
}

func (u *CardController) DeleteMember(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	res, err := u.cardClient.DeleteCardMember(context.Background(), &card.DeleteCardMemberRequest{
		Id:     id,
		UserId: userId,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) DeleteCard(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := u.cardClient.DeleteCard(context.Background(), &card.DeleteCardRequest{
		Id: id,
	})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func NewCardController(userClient user.UserClient, cardClient card.CardClient) *CardController {
	return &CardController{userClient: userClient, cardClient: cardClient}
}
