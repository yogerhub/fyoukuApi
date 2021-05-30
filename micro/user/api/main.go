package main

import (
	"context"
	"encoding/json"
	user "fyoukuApi/micro/user/proto"
	"log"
	"strings"
	"time"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type User struct {
	Client user.UserServerService
}

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"10.211.55.2:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.api.fyoukuApi.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserServerService("go.micro.service.fyoukuApi.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (u *User) LoginDo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到User.LoginDo API请求")
	//接受参数
	mobile, ok := req.Post["mobile"]
	if !ok || len(mobile.Values) == 0 {
		return errors.BadRequest("go.micro.api.fyoukuApi.user", "mobile为空")
	}
	password, ok := req.Post["password"]
	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.fyoukuApi.user", "password为空")
	}

	response, err := u.Client.LoginDo(ctx, &user.RequestLoginDo{
		Mobile:   strings.Join(mobile.Values, ""),
		Password: strings.Join(password.Values, ""),
	})
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.Code,
		"msg":   response.Msg,
		"items": response.Items,
		"count": response.Count,
	})
	rsp.Body = string(b)
	return nil
}
