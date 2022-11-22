package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/dto"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/util"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"strings"
)

type ObjectController struct {
	cosClient *cos.Client
}

func (o *ObjectController) PutObject(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	obj, err := file.Open()
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	// 获得文件名 文件后缀
	prefix, postfix := util.GetPrefixAndPosfix(file.Filename)

	// 生成随机后缀
	id, _ := uuid.NewUUID()
	prefix += strings.Replace(id.String(), "-", "", -1)

	// 上传文件
	name := fmt.Sprintf("%s.%s", prefix, postfix)
	_, err = o.cosClient.Object.Put(context.Background(), name, obj, nil)
	if err != nil {
		c.JSON(http.StatusOK, dto.NewErrorInternalDto(err))
		return
	}

	c.JSON(http.StatusOK, dto.NewOkDto(struct {
		Url string `json:"url"`
	}{
		Url: "https://" + o.cosClient.BaseURL.BucketURL.Host + "/" + name,
	}))
}

func NewObjectController(cosClient *cos.Client) *ObjectController {
	return &ObjectController{cosClient: cosClient}
}
