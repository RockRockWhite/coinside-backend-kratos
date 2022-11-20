package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func TeamRouter(r *gin.Engine, controller *controller.TeamController) *gin.Engine {
	//selfCond := func(c *gin.Context) bool {
	//	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	//	claims := c.MustGet("claims").(*util.JwtClaims)
	//
	//	return id == claims.Id
	//}

	team := r.Group("/teams")
	{
		team.GET("/:id", middleware.JwtAuth(nil), controller.GetTeamInfo)
		team.POST("", controller.CreateTeam)
		team.POST("/:team_id/mumbers", controller.SetTeamMember)
		team.POST("/:id", controller.SetTeamAdmin)

		team.PUT(":id/name", middleware.JwtAuth(nil), controller.SetName)
		team.PUT(":id/avatar", middleware.JwtAuth(nil), controller.SetAvatar)
		team.PUT(":id/email", middleware.JwtAuth(nil), controller.SetEmail)
		team.PUT(":id/description", middleware.JwtAuth(nil), controller.SetDescription)
		team.PUT(":id/website", middleware.JwtAuth(nil), controller.SetWebsite)

		team.DELETE("/:id", middleware.JwtAuth(nil), controller.DeleteTeam)
		team.DELETE("/:team_id/members/:user_id", middleware.JwtAuth(nil), controller.DeleteTeamMember)

	}

	return r
}
