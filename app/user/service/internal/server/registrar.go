package server

import (
	"github.com/go-kratos/kratos/contrib/registry/zookeeper/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-zookeeper/zk"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/config"
	"time"
)

func NewRegistrar(conf *config.Config) registry.Registrar {
	conn, _, err := zk.Connect([]string{conf.GetString("zookeeper.address")}, time.Duration(int64(conf.GetInt("zookeeper.timeout")))*time.Second)
	if err != nil {
		panic(err)
	}
	return zookeeper.New(conn)
}
