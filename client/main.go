package main

import (
	"context"
	"log"
	"time"

	proto "github.com/PieterVoorwinden/micro-http/proto"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService()
	service.Init()
	for {
		time.Sleep(1 * time.Second)
		client := proto.NewGreeterService("Greeter", service.Client())
		rsp, err := client.Hello(context.Background(), &proto.Request{
			Name: "Pieter",
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(rsp.Msg)
	}
}
