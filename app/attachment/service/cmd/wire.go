// The build tag makes sure the stub is not built in the final build.
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"github.com/ljxsteam/coinside-backend-kratos/app/attachment/service/config"
	"github.com/ljxsteam/coinside-backend-kratos/app/attachment/service/internal/data"
	"github.com/ljxsteam/coinside-backend-kratos/app/attachment/service/internal/server"
	"github.com/ljxsteam/coinside-backend-kratos/app/attachment/service/internal/service"
)

// initApp init kratos application.
func initApp(*config.Config) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
