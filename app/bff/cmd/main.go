package main

import (
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/util"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
	"log"
)

func main() {
	conf := config.NewConfig()
	util.InitJwtUtil(conf)
	r := initRouter(conf)

	err := r.Run(conf.GetString("server.addr"))

	if err != nil {
		log.Fatal(err.Error())
	}
}
