package client

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	grpc "github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/ljxsteam/coinside-backend-kratos/api/vote"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
)

func NewVoteClinet(conf *config.Config, dis registry.Discovery) vote.VoteClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(fmt.Sprintf("discovery:///%s", conf.GetString("service.vote.name"))),
		grpc.WithDiscovery(dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}

	return vote.NewVoteClient(conn)
}
