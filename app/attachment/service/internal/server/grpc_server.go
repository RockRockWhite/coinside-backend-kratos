package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	api "github.com/ljxsteam/coinside-backend-kratos/api/attachment"
	"github.com/ljxsteam/coinside-backend-kratos/app/attachment/service/config"
	"github.com/ljxsteam/coinside-backend-kratos/app/attachment/service/internal/service"
)

func NewGrpcServer(config *config.Config, server *service.AttachmentService) *grpc.Server {
	srv := grpc.NewServer(grpc.Address(config.GetString("server.addr")),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(log.DefaultLogger),
		))
	api.RegisterAttachmentServer(srv, server)

	return srv
}
