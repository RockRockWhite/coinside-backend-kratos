package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewApp(name string, srv *grpc.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.Name(name),
		kratos.Server(srv),
		kratos.Registrar(r))
}

func main() {
	fmt.Println("Hello World")
}
