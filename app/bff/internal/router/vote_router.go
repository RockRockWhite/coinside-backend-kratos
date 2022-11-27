package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func VoteRouter(r *gin.Engine, controller *controller.VoteController) *gin.Engine {
	//selfCond := func(c *gin.Context) bool {
	//	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	//	claims := c.MustGet("claims").(*util.JwtClaims)
	//
	//	return id == claims.Id
	//}

	vote := r.Group("/modules/votes")
	{
		vote.GET("/:id", middleware.JwtAuth(nil), controller.GetVoteInfo)
		vote.POST("", middleware.JwtAuth(nil), controller.CreateVote)
		vote.POST("/:id/items", controller.SetVoteItem)

		vote.PUT(":id/title", middleware.JwtAuth(nil), controller.SetTitle)
		vote.PUT(":id/items/:item_id/content", middleware.JwtAuth(nil), controller.SetItemContent)

		vote.DELETE("/:id", middleware.JwtAuth(nil), controller.DeleteVote)
		vote.DELETE("/:id/items/:item_id", middleware.JwtAuth(nil), controller.DeleteVoteItem)

	}

	return r
}
