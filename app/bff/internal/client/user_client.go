package client

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ljxsteam/coinside-backend-kratos/api/user"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
)

func NewUserClinet(conf *config.Config, dis registry.Discovery) user.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(fmt.Sprintf("discovery:///%s", conf.GetString("service.user.name"))),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}

	return user.NewUserClient(conn)
}
