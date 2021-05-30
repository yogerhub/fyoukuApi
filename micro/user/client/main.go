package main

import (
	"context"
	"fmt"
	"fyoukuApi/micro/user/proto"
	_ "fyoukuApi/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"10.211.55.2:8500",
		}
	})

	// New Service
	service := micro.NewService(
		//添加注册中心
		micro.Registry(consul),
	)
	service.Init()

	user := proto.NewUserServerService("go.micro.service.fyoukuApi.user", service.Client())

	rsp, err := user.LoginDo(context.TODO(), &proto.RequestLoginDo{
		Mobile:   "18340812561",
		Password: "123456",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp)
}
