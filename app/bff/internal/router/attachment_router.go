package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func AttachmentRouter(r *gin.Engine, controller *controller.AttachmentController) *gin.Engine {
	//selfCond := func(c *gin.Context) bool {
	//	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	//	claims := c.MustGet("claims").(*util.JwtClaims)
	//
	//	return id == claims.Id
	//}

	attachment := r.Group("/modules")
	{
		attachment.GET("/attachments/:id", middleware.JwtAuth(nil), controller.GetAttachmentInfo)

		attachment.POST("/attachments", middleware.JwtAuth(nil), controller.CreateAttachment)

		attachment.PUT("/attachments/:id", middleware.JwtAuth(nil), controller.SetLink)

		attachment.DELETE("/attachments/:id", middleware.JwtAuth(nil), controller.DeleteAttachment)

		attachment.PUT("/attachments/:id", middleware.JwtAuth(nil), controller.SetDownloadCount)
	}

	return r
}
