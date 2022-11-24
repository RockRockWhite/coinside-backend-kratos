package client

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ljxsteam/coinside-backend-kratos/api/todo"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
)

func NewTodoClinet(conf *config.Config, dis registry.Discovery) todo.TodoServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(fmt.Sprintf("discovery:///%s", conf.GetString("service.todo.name"))),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}

	return todo.NewTodoServiceClient(conn)
}
