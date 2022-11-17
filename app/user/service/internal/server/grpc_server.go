package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	api "github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/config"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/internal/service"
)

func NewGrpcServer(config *config.Config, server *service.UserService) *grpc.Server {
	srv := grpc.NewServer(grpc.Address(config.GetString("server.addr")),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(log.DefaultLogger),
		))
	api.RegisterUserServer(srv, server)

	return srv
}
