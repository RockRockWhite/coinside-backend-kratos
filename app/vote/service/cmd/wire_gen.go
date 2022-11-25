// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/ljxsteam/coinside-backend-kratos/app/vote/service/internal/data"
	"github.com/ljxsteam/coinside-backend-kratos/app/vote/service/internal/server"
	"github.com/ljxsteam/coinside-backend-kratos/app/vote/service/internal/service"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(configConfig *config.Config) (*kratos.App, func(), error) {
	db := data.NewDB(configConfig)
	voteRepo := data.NewVoteRepoNoCache(db)
	voteService := service.NewVoteService(voteRepo)
	grpcServer := server.NewGrpcServer(configConfig, voteService)
	registrar := server.NewRegistrar(configConfig)
	app := newApp(configConfig, grpcServer, registrar)
	return app, func() {
	}, nil
}