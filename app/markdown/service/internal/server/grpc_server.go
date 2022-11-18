package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	api "github.com/ljxsteam/coinside-backend-kratos/api/markdown"
	"github.com/ljxsteam/coinside-backend-kratos/app/markdown/service/config"
	"github.com/ljxsteam/coinside-backend-kratos/app/markdown/service/internal/service"
)

func NewGrpcServer(config *config.Config, server *service.MarkdownService) *grpc.Server {
	srv := grpc.NewServer(grpc.Address(config.GetString("server.addr")),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(log.DefaultLogger),
		))
	api.RegisterMarkdownServer(srv, server)

	return srv
}
