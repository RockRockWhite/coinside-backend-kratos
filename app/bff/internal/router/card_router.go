package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func CardRouter(r *gin.Engine, controller *controller.CardController) *gin.Engine {
	//selfCond := func(c *gin.Context) bool {
	//	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	//	claims := c.MustGet("claims").(*util.JwtClaims)
	//
	//	return id == claims.Id
	//}

	card := r.Group("/cards")
	{

		card.GET("/:id", middleware.JwtAuth(nil), controller.GetCardInfo)
		card.POST("", controller.CreateCard)

		card.PUT(":id/title", middleware.JwtAuth(nil), controller.SetTitle)
		card.PUT(":id/content", middleware.JwtAuth(nil), controller.SetContent)
		card.PUT(":id/members/:user_id", middleware.JwtAuth(nil), controller.SetMember)
		card.DELETE(":id/members/:user_id", middleware.JwtAuth(nil), controller.DeleteMember)
		card.PUT(":id/tags/:tag_content", middleware.JwtAuth(nil), controller.AddTag)
		card.DELETE(":id/tags/:tag_content", middleware.JwtAuth(nil), controller.DeleleTag)
		card.DELETE("/:id", middleware.JwtAuth(nil), controller.DeleteCard)
	}

	return r
}
