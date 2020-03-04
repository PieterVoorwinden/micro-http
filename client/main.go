package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	proto "github.com/PieterVoorwinden/micro-http/proto"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
)

type Handler struct {
	c proto.GreeterService
}

func (h *Handler) Greeter(w http.ResponseWriter, r *http.Request) {
	fmt.Println(runtime.NumGoroutine())
	log.Println("Got request")
	rsp, err := h.c.Hello(r.Context(), &proto.Request{
		Name: "Pieter",
	})
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(200)
	w.Write([]byte(rsp.Msg))
	fmt.Println(runtime.NumGoroutine())
}

func main() {
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
	)
	service.Init()

	h := &Handler{
		c: proto.NewGreeterService("go.micro.srv.greeter", client.DefaultClient),
	}
	router := mux.NewRouter()
	router.HandleFunc("/greeter", h.Greeter)
	service.Handle("/", router)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
