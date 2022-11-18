package main

import (
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/client"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/router"
	"github.com/ljxsteam/coinside-backend-kratos/app/user/service/config"
	"log"
)

func Server(conf *config.Config, userController *controller.UserController) {

	r := router.NewApiRouter(userController)

	err := r.Run(conf.GetString("server.addr"))

	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	c := controller.NewUserController(client.NewUserClinet(config.NewConfig(), client.NewZkDiscovery(config.NewConfig())))
	Server(config.NewConfig(), c)

	//
	//c, _ := client.NewUserClinet(config.NewConfig(), client.NewZkDiscovery(config.NewConfig()))
	//
	//res, err := c.GetUserInfo(context.Background(), &api.GetUserInfoRequest{
	//	Id: 1,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(res)
}
