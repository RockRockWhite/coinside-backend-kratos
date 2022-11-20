package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/card"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"net/http"
	"strconv"
)

type CardController struct {
	client card.CardClient
}

func (u *CardController) GetCardInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := u.client.GetCardInfo(context.Background(), &card.GetCardInfoRequest{Id: id})

	resDto := dto.ResponseDto{
		Code:    dto.CardErrorCode[res.Code].Code,
		Message: dto.CardErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != card.Code_OK {
		resDto.Data = err
	} else {
		resDto.Data = res.Info
	}

	c.JSON(http.StatusOK, resDto)
}

func (u *CardController) CreateCard(c *gin.Context) {
	var req card.CreateCardRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := u.client.CreateCard(context.Background(), &req)

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

	res, err := u.client.UpdateCardTitle(context.Background(), &card.UpdateCardTitleRequest{
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

	res, err := u.client.UpdateCardContent(context.Background(), &card.UpdateCardContentRequest{
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

func (u *CardController) AddTag(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	tagContent := c.Param("tag_content")

	res, err := u.client.AddCardTag(context.Background(), &card.AddCardTagRequest{
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

	res, err := u.client.DeleteCardTag(context.Background(), &card.DeleteCardTagRequest{
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

	res, err := u.client.SetCardMember(context.Background(), &card.SetCardMemberRequest{
		Id:      id,
		UserId:  userId,
		IsAdmin: reqDto.IsAdmin,
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

func (u *CardController) DeleteMember(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, _ := strconv.ParseUint(c.Param("user_id"), 10, 64)

	res, err := u.client.DeleteCardMember(context.Background(), &card.DeleteCardMemberRequest{
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

	res, err := u.client.DeleteCard(context.Background(), &card.DeleteCardRequest{
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

func NewCardController(client card.CardClient) *CardController {
	return &CardController{client: client}
}
