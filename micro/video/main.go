package main

import (
	"fmt"
	"fyoukuApi/controllers"
	"fyoukuApi/micro/video/proto"
	_ "fyoukuApi/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {

	beego.LoadAppConfig("ini","../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb)

	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"10.211.55.2:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.fyoukuApi.video"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8022"),
		//添加注册中心
		micro.Registry(consul),
	)

	service.Init()

	proto.RegisterVideoServiceHandler(service.Server(), new(controllers.VideoRpcController))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
