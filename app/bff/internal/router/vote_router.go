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

	team := r.Group("/modules/votes")
	{
		team.GET("/:id", middleware.JwtAuth(nil), controller.GetVoteInfo)
		team.POST("", middleware.JwtAuth(nil), controller.CreateVote)
		team.POST("/:id/items", controller.SetVoteItem)

		team.PUT(":id/title", middleware.JwtAuth(nil), controller.SetTitle)
		team.PUT(":id/items/:item_id/content", middleware.JwtAuth(nil), controller.SetItemContent)

		team.DELETE("/:id", middleware.JwtAuth(nil), controller.DeleteVote)
		team.DELETE("/:id/items/:item_id", middleware.JwtAuth(nil), controller.DeleteVoteItem)

	}

	return r
}
