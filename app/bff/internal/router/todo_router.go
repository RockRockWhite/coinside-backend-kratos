package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func TodoRouter(r *gin.Engine, controller *controller.TodoController) *gin.Engine {
	//selfCond := func(c *gin.Context) bool {
	//	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	//	claims := c.MustGet("claims").(*util.JwtClaims)
	//
	//	return id == claims.Id
	//}

	team := r.Group("/modules/todos")
	{
		team.GET("/:id", middleware.JwtAuth(nil), controller.GetTodoInfo)
		team.POST("", middleware.JwtAuth(nil), controller.CreateTodo)
		team.POST("/:id/items", controller.SetTodoItem)

		team.PUT(":id/title", middleware.JwtAuth(nil), controller.SetTitle)
		team.PUT(":id/items/:item_id/content", middleware.JwtAuth(nil), controller.SetItemContent)
		team.PUT(":id/items/:item_id/finished", middleware.JwtAuth(nil), controller.SetItemFinished)

		team.DELETE("/:id", middleware.JwtAuth(nil), controller.DeleteTodo)
		team.DELETE("/:id/items/:item_id", middleware.JwtAuth(nil), controller.DeleteTodoItem)

	}

	return r
}
