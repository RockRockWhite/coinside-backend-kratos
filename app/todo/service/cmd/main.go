package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ljxsteam/coinside-backend-kratos/app/todo/service/config"
)

func newApp(conf *config.Config, srv *grpc.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(conf.GetString("server.name")),
		kratos.Server(srv),
		kratos.Registrar(r))
}

func main() {
	app, _, err := initApp(config.NewConfig())
	if err != nil {
		return
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
