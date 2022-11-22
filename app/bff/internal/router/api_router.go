package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func NewApiRouter(userController *controller.UserController, cardController *controller.CardController, teamController *controller.TeamController, markdownController *controller.MarkdownController, objectController *controller.ObjectController) *gin.Engine {
	// 初始化Controllers
	router := gin.Default()

	// 配置中间件
	router.Use(middleware.Cors)

	// 配置路由
	router = UserRouter(router, userController)
	router = CardRouter(router, cardController)
	router = TeamRouter(router, teamController)
	router = MarkdownRouter(router, markdownController)
	router = ObjectRouter(router, objectController)

	return router
}
