package main

import (
	"github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-zookeeper/zk"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/config"
	"time"
)

func newApp(name string, srv *grpc.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(name),
		kratos.Server(srv),
		kratos.Registrar(r))
}

func main() {
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second*10)
	if err != nil {
		panic(err)
	}

	r := zookeeper.New(conn)

	app, _, err := initApp("user", r, config.NewConfig())
	if err != nil {
		return
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
