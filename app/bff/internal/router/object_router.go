package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
)

func ObjectRouter(r *gin.Engine, controller *controller.ObjectController) *gin.Engine {

	bucket := r.Group("/objects")
	{
		bucket.PUT("", middleware.JwtAuth(nil), controller.PutObject)
	}

	return r
}
