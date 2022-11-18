package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
)

func UserRouter(r *gin.Engine, controller *controller.UserController) *gin.Engine {
	user := r.Group("/users")
	{
		user.GET("/:id", controller.GetUserInfo)
		user.GET("/id", controller.GetUserId)
		user.POST("", controller.CreateUser)
	}

	token := r.Group("/tokens")
	{
		token.POST("", controller.Login)
	}

	return r
}
