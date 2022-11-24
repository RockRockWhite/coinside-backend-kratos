package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func CardRouter(r *gin.Engine, controller *controller.CardController) *gin.Engine {
	card := r.Group("/cards")
	{
		card.GET("/:id", middleware.JwtAuth(nil), controller.IsCardMember(false), controller.GetCardInfo)
		card.GET("", middleware.JwtAuth(nil), controller.GetCardInfoList)

		card.POST("", middleware.JwtAuth(nil), controller.CreateCard)

		card.PUT(":id/title", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.SetTitle)
		card.PUT(":id/content", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.SetContent)
		card.PUT(":id/deadline", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.SetDeadline)
		card.PUT(":id/status", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.SetStatus)
		card.PUT(":id/members/:user_id", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.SetMember)
		card.PUT(":id/tags/:tag_content", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.AddTag)

		card.DELETE(":id/members/:user_id", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.DeleteMember)
		card.DELETE(":id/tags/:tag_content", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.DeleleTag)
		card.DELETE("/:id", middleware.JwtAuth(nil), controller.IsCardMember(true), controller.DeleteCard)
	}

	return r
}
