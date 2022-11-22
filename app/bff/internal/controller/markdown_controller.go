package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/markdown"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"net/http"
	"strconv"
)

type MarkdownController struct {
	client markdown.MarkdownClient
}

func (m *MarkdownController) GetMarkdownInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := m.client.GetMarkdownById(context.Background(), &markdown.GetMarkdownByIdRequest{Id: id})

	resDto := dto.ResponseDto{
		Code:    dto.MarkdownErrorCode[res.Code].Code,
		Message: dto.MarkdownErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != markdown.Code_OK {
		resDto.Data = err
	} else {
		resDto.Data = res.Markdown
	}

	c.JSON(http.StatusOK, resDto)
}

func (m *MarkdownController) CreateMarkdown(c *gin.Context) {
	var req markdown.MarkdownInfo

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := m.client.AddMarkdown(context.Background(), &req)

	resDto := dto.ResponseDto{
		Code:    dto.MarkdownErrorCode[res.Code].Code,
		Message: dto.MarkdownErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != markdown.Code_OK {
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

func (m *MarkdownController) SetContent(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Content string `json:"content"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := m.client.UpdateMarkdown(context.Background(), &markdown.MarkdownInfo{
		Id:      id,
		Content: reqDto.Content,
	})

	resDto := dto.ResponseDto{
		Code:    dto.MarkdownErrorCode[res.Code].Code,
		Message: dto.MarkdownErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != markdown.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (m *MarkdownController) DeleteMarkdown(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := m.client.DeleteMarkdown(context.Background(), &markdown.DeleteMarkdownRequest{
		Id: id,
	})

	resDto := dto.ResponseDto{
		Code:    dto.MarkdownErrorCode[res.Code].Code,
		Message: dto.MarkdownErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != markdown.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func NewMarkdownController(client markdown.MarkdownClient) *MarkdownController {
	return &MarkdownController{client: client}
}
