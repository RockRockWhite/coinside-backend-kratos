package middleware

import (
	"github.com/gin-gonic/gin"
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
			// todo: 统一错误
			c.JSON(http.StatusUnauthorized, nil)

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

		//access := false
		//if roleFlag&Role_All != 0 { // 所有人通行
		//	access = true
		//} else if roleFlag&Role_Admin != 0 && claims.IsAdmin { // 管理员通行
		//	access = true
		//} else if roleFlag&Role_Cond != 0 && condFunc != nil && condFunc(c) { // 符合条件用户通行
		//	access = true
		//}
		//
		//if !access {
		//	c.JSON(http.StatusForbidden, dtos.ErrorDto{
		//		Message:          "The token cannot access this resource.",
		//		DocumentationUrl: viper.GetString("Document.Url"),
		//	})
		//
		//	c.Abort()
		//	return
		//}
	}
}
