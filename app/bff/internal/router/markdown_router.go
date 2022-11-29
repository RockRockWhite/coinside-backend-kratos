package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func MarkdownRouter(r *gin.Engine, controller *controller.MarkdownController) *gin.Engine {
	//selfCond := func(c *gin.Context) bool {
	//	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	//	claims := c.MustGet("claims").(*util.JwtClaims)
	//
	//	return id == claims.Id
	//}

	markdown := r.Group("/modules")
	{
		markdown.GET("/markdowns/:id", middleware.JwtAuth(nil), controller.GetMarkdownInfo)

		markdown.POST("/markdowns", middleware.JwtAuth(nil), controller.CreateMarkdown)

		markdown.PUT("/markdowns/:id", middleware.JwtAuth(nil), controller.SetContent)

		markdown.DELETE("/markdowns/:id", middleware.JwtAuth(nil), controller.DeleteMarkdown)
	}

	return r
}
