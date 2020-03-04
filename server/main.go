package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	proto "github.com/PieterVoorwinden/micro-http/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	go func() {
		t := time.NewTicker(3 * time.Second)
		for range t.C {
			fmt.Println(runtime.NumGoroutine())
		}
	}()
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(consul.NewRegistry()),
	)
	service.Init()

	if err := proto.RegisterGreeterHandler(service.Server(), new(Greeter)); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
