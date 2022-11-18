package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/util"
	"strconv"
)

func UserRouter(r *gin.Engine, controller *controller.UserController) *gin.Engine {
	user := r.Group("/users")
	{
		user.GET("/:id", middleware.JwtAuth(func(c *gin.Context) bool {
			id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
			claims := c.MustGet("claims").(*util.JwtClaims)

			return id == claims.Id
		}), controller.GetUserInfo)
		user.GET("/id", controller.GetUserId)
		user.POST("", controller.CreateUser)
	}

	token := r.Group("/tokens")
	{
		token.POST("", controller.Login)
	}

	return r
}
