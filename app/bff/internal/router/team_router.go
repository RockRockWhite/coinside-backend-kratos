package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func TeamRouter(r *gin.Engine, controller *controller.TeamController) *gin.Engine {
	team := r.Group("/teams")
	{
		team.GET("/:id", middleware.JwtAuth(nil), controller.GetTeamInfo)
		team.GET("", middleware.JwtAuth(nil), controller.GetTeamInfoList)

		team.POST("", middleware.JwtAuth(nil), controller.CreateTeam)
		team.POST("/:id/members/:user_id", controller.SetTeamMember)
		team.POST("/:id/admins/:user_id", controller.SetTeamAdmin)

		team.PUT(":id", middleware.JwtAuth(nil), controller.Update)
		team.PUT(":id/name", middleware.JwtAuth(nil), controller.SetName)
		team.PUT(":id/avatar", middleware.JwtAuth(nil), controller.SetAvatar)
		team.PUT(":id/email", middleware.JwtAuth(nil), controller.SetEmail)
		team.PUT(":id/description", middleware.JwtAuth(nil), controller.SetDescription)
		team.PUT(":id/website", middleware.JwtAuth(nil), controller.SetWebsite)

		team.DELETE("/:id", middleware.JwtAuth(nil), controller.DeleteTeam)
		team.DELETE("/:id/members/:user_id", middleware.JwtAuth(nil), controller.DeleteTeamMember)

	}

	return r
}
