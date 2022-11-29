package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/api/attachment"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"net/http"
	"strconv"
)

type AttachmentController struct {
	client attachment.AttachmentClient
}

func (a *AttachmentController) GetAttachmentInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := a.client.GetAttachmentById(context.Background(), &attachment.GetAttachmentByIdRequest{Id: id})

	resDto := dto.ResponseDto{
		Code:    dto.AttachmentErrorCode[res.Code].Code,
		Message: dto.AttachmentErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != attachment.Code_OK {
		resDto.Data = err
	} else {
		resDto.Data = res.Attachment
	}

	c.JSON(http.StatusOK, resDto)
}

func (a *AttachmentController) CreateAttachment(c *gin.Context) {
	var req attachment.AddAttachmentRequest

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := a.client.AddAttachment(context.Background(), &req)

	resDto := dto.ResponseDto{
		Code:    dto.AttachmentErrorCode[res.Code].Code,
		Message: dto.AttachmentErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != attachment.Code_OK {
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

func (a *AttachmentController) SetLink(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		Link          string `json:"link"`
		DownloadCount uint64 `json:"download_count"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := a.client.UpdateAttachment(context.Background(), &attachment.UpdateAttachmentRequest{
		Id:            id,
		Link:          reqDto.Link,
		DownloadCount: reqDto.DownloadCount,
	})

	resDto := dto.ResponseDto{
		Code:    dto.AttachmentErrorCode[res.Code].Code,
		Message: dto.AttachmentErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != attachment.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (a *AttachmentController) SetDownloadCount(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	reqDto := struct {
		DownloadCount uint64 `json:"download_count"`
		Link          string `json:"link"`
	}{}
	if err := c.ShouldBind(&reqDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorBadRequest)
		return
	}

	res, err := a.client.UpdateAttachment(context.Background(), &attachment.UpdateAttachmentRequest{
		Id:            id,
		DownloadCount: reqDto.DownloadCount,
		Link:          reqDto.Link,
	})

	resDto := dto.ResponseDto{
		Code:    dto.AttachmentErrorCode[res.Code].Code,
		Message: dto.AttachmentErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != attachment.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func (a *AttachmentController) DeleteAttachment(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	res, err := a.client.DeleteAttachment(context.Background(), &attachment.DeleteAttachmentRequest{
		Id: id,
	})

	resDto := dto.ResponseDto{
		Code:    dto.AttachmentErrorCode[res.Code].Code,
		Message: dto.AttachmentErrorCode[res.Code].Message,
		Data:    nil,
	}

	if res.Code != attachment.Code_OK {
		resDto.Data = err
	} else {
	}

	c.JSON(http.StatusOK, resDto)
}

func NewAttachmentController(client attachment.AttachmentClient) *AttachmentController {
	return &AttachmentController{client: client}
}
