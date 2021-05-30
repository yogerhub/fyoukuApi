package main

import (
	"context"
	"encoding/json"
	video "fyoukuApi/micro/video/proto"
	"log"
	"strings"
	"time"

	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type Video struct {
	Client video.VideoService
}

func main() {
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"10.211.55.2:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.api.fyoukuApi.video"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&Video{Client: video.NewVideoService("go.micro.service.fyoukuApi.video", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (v *Video) ChannelAdvert(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到Video.ChannelAdvert API请求")
	//接受参数
	channelId, ok := req.Get["channelId"]
	if !ok || len(channelId.Values) == 0 {
		return errors.BadRequest("go.micro.api.fyoukuApi.video", "channelId为空")
	}

	response, err := v.Client.ChannelAdvert(ctx, &video.RequestChannelAdvert{
		ChannelId: strings.Join(channelId.Values, ""),
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
func (v *Video) ChannelHotList(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到Video.ChannelHotList API请求")
	//接受参数
	channelId, ok := req.Get["channelId"]
	if !ok || len(channelId.Values) == 0 {
		return errors.BadRequest("go.micro.api.fyoukuApi.video", "channelId为空")
	}

	response, err := v.Client.ChannelHotList(ctx, &video.RequestChannelHotList{
		ChannelId: strings.Join(channelId.Values, ""),
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
