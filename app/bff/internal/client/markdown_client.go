package client

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	grpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ljxsteam/coinside-backend-kratos/api/markdown"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
)

func NewMarkdownClinet(conf *config.Config, dis registry.Discovery) markdown.MarkdownClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(fmt.Sprintf("discovery:///%s", conf.GetString("service.markdown.name"))),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}

	return markdown.NewMarkdownClient(conn)
}
