package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func NewApiRouter(userController *controller.UserController, cardController *controller.CardController, teamController *controller.TeamController) *gin.Engine {
	// 初始化Controllers
	router := gin.Default()

	// 配置中间件
	router.Use(middleware.Cors)

	// 配置路由
	router = UserRouter(router, userController)
	router = CardRouter(router, cardController)
	router = TeamRouter(router, teamController)
	//user := router.Group("/users")
	//{
	//	user.GET("/:username", controllers.GetUser)
	//	user.GET("", controllers.GetUsers)
	//	user.GET("/count", controllers.CountUser)
	//	user.POST("", controllers.AddUser)
	//	user.PATCH(
	//		"/:username",
	//		middlewares.JwtAuth(middlewares.Role_Admin|middlewares.Role_Cond, func(c *gin.Context) bool {
	//			username := c.Param("username")
	//			claims := c.MustGet("claims").(*utils.JwtClaims)
	//			return username == claims.Username
	//		}), controllers.PatchUser)
	//	user.PATCH(
	//		"/:username/password",
	//		middlewares.JwtAuth(middlewares.Role_Cond, func(c *gin.Context) bool {
	//			username := c.Param("username")
	//			claims := c.MustGet("claims").(*utils.JwtClaims)
	//			return username == claims.Username
	//		}), controllers.PatchPassword)
	//	user.DELETE("/:username",
	//		middlewares.JwtAuth(middlewares.Role_Admin|middlewares.Role_Cond, func(c *gin.Context) bool {
	//			username := c.Param("username")
	//			claims := c.MustGet("claims").(*utils.JwtClaims)
	//			return username == claims.Username
	//		}), controllers.DeleteUser)
	//}
	//
	//token := router.Group("/tokens")
	//{
	//	token.POST("", controllers.CreateToken)
	//}

	return router
}
