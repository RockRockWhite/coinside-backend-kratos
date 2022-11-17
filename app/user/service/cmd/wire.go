// The build tag makes sure the stub is not built in the final build.
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/config"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/internal/data"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/internal/server"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/internal/service"
)

// initApp init kratos application.
func initApp(string, registry.Registrar, *config.Config) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
