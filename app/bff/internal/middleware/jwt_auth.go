package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/error_code"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/util"
	"net/http"
	"strings"
)

// JwtAuth JwtToken验证中间件
func JwtAuth(condFunc func(c *gin.Context) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" || strings.Fields(token)[0] != "Bearer" {
			// 没有传token参数
			c.JSON(http.StatusOK, struct {
				error_code.Error
				Data interface{} `json:"data"`
			}{
				Error: error_code.Error{
					Code:    "ERROR_UNAUTHORIZED",
					Message: "Unauthorized.",
				},
				Data: nil,
			})

			c.Abort()
			return
		}
		token = strings.Fields(token)[1]

		claims, err := util.ParseJwtToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)

			c.Abort()
			return
		}

		// Claims写入上下文
		c.Set("claims", claims)

		access := true

		if condFunc != nil && !condFunc(c) {
			access = false
		}

		if !access {
			c.JSON(http.StatusForbidden, struct {
				error_code.Error
				Data interface{} `json:"data"`
			}{
				Error: error_code.Error{
					Code:    "ERROR_UNAUTHORIZED",
					Message: "Unauthorized.",
				},
				Data: nil,
			})

			c.Abort()
			return
		}
	}
}
