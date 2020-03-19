package main

import (
	"context"
	"fmt"
	"log"
	"time"

	proto "github.com/PieterVoorwinden/micro-http/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(consul.NewRegistry()),
	)
	service.Init()

	go func() {
		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(service.Client().Options())
	srv := proto.NewGreeterService("go.micro.srv.greeter", service.Client())
	for {
		_, err := srv.Hello(context.Background(), &proto.Request{
			Name: "Pieter",
		})
		if err != nil {
			log.Println(err)
		}
	}
}
