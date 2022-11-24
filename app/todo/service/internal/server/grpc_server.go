package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	api "github.com/ljxsteam/coinside-backend-kratos/api/todo"
	"github.com/ljxsteam/coinside-backend-kratos/app/todo/service/internal/service"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
)

func NewGrpcServer(config *config.Config, server *service.TodoService) *grpc.Server {
	srv := grpc.NewServer(grpc.Address(config.GetString("server.addr")),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(log.DefaultLogger),
		))
	api.RegisterTodoServiceServer(srv, server)

	return srv
}
