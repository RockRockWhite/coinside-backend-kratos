// The build tag makes sure the stub is not built in the final build.
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/client"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/router"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/config"
)

// initRouter init gin router.
func initRouter(conf *config.Config) *gin.Engine {
	panic(wire.Build(client.ProviderSet, controller.ProviderSet, router.ProviderSet))
}
