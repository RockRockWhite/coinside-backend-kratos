package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
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
		fmt.Println(card)
		//user.GET("/:id", middleware.JwtAuth(selfCond), controller.GetUserInfo)
		//user.GET("/id", controller.GetUserId)
		//user.POST("", controller.CreateUser)
		//
		//user.PUT(":id/fullname", middleware.JwtAuth(selfCond), controller.SetFullname)
		//user.PUT(":id/avatar", middleware.JwtAuth(selfCond), controller.SetAvatar)
		//user.PUT(":id/config", middleware.JwtAuth(selfCond), controller.SetConfig)
		//user.PUT(":id/mobile", middleware.JwtAuth(selfCond), controller.SetMobile)
		//user.PUT(":id/email", middleware.JwtAuth(selfCond), controller.SetEmail)
		//
		//user.DELETE("/:id", middleware.JwtAuth(selfCond), controller.DeleteUser)
	}

	return r
}
