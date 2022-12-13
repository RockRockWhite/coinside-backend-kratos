package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
)

func NewApiRouter(conf *config.Config, userController *controller.UserController, cardController *controller.CardController, attachmentController *controller.AttachmentController, teamController *controller.TeamController, voteController *controller.VoteController, todoController *controller.TodoController, markdownController *controller.MarkdownController, objectController *controller.ObjectController) *gin.Engine {
	// 初始化Controllers
	router := gin.Default()

	// 配置跨域中间件
	router.Use(middleware.Cors)

	// 配置路由
	router = UserRouter(router, conf, userController)
	router = CardRouter(router, cardController)
	router = TeamRouter(router, teamController)
	router = MarkdownRouter(router, markdownController)
	router = ObjectRouter(router, objectController)
	router = TodoRouter(router, todoController)
	router = VoteRouter(router, voteController)
	router = AttachmentRouter(router, attachmentController)

	return router
}
