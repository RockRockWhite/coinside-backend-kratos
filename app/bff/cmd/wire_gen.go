// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/client"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/router"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
)

// Injectors from wire.go:

// initRouter init gin router.
func initRouter(conf *config.Config) *gin.Engine {
	discovery := client.NewZkDiscovery(conf)
	userClient := client.NewUserClinet(conf, discovery)
	userController := controller.NewUserController(userClient)
	cardClient := client.NewCardClinet(conf, discovery)
	cardController := controller.NewCardController(userClient, cardClient)
	teamClient := client.NewTeamClinet(conf, discovery)
	teamController := controller.NewTeamController(userClient, teamClient)
	engine := router.NewApiRouter(userController, cardController, teamController)
	return engine
}
