package main

import (
	"context"
	"log"
	"time"

	proto "github.com/PieterVoorwinden/micro-http/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	srv := server.NewServer(
		server.Name("Greeter"),
		server.Wait(nil),
	)
	service := micro.NewService(
		micro.Name("Greeter"),
		micro.Server(srv),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init()

	if err := proto.RegisterGreeterHandler(service.Server(), &Greeter{}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
