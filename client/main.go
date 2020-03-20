package main

import (
	"context"
	"fmt"
	"log"
	"time"

	proto "github.com/PieterVoorwinden/micro-http/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	httpServer "github.com/micro/go-plugins/server/http/v2"
)

func main() {
	service := micro.NewService(
		micro.Server(httpServer.NewServer()),
		micro.Name("go.micro.api.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(consul.NewRegistry()),
	)
	service.Init()

	message := make([]byte, (1024*1024*4)-5)

	time.Sleep(5 * time.Second)
	fmt.Println(service.Client().Options())
	srv := proto.NewGreeterService("go.micro.srv.greeter", service.Client())
	for {
		_, err := srv.Hello(context.Background(), &proto.Request{
			Name: string(message),
		})
		if err != nil {
			log.Println(err)
		}
	}
}
