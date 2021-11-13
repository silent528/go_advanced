package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	kitHttp "go_advanced/library/kit/http"
	v1 "go_advanced/task3/api/v1"
	"go_advanced/task3/internal/service"
	"google.golang.org/grpc"
	"log"
)

func NewHttpServer(greeter *service.UserService) *kitHttp.Server {
	conn, err := grpc.Dial("localhost:10011", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	mux := runtime.NewServeMux()
	err = v1.RegisterUserServiceHandler(context.Background(), mux, conn)
	if err != nil {
		panic("grpc register err:" + err.Error())
	}
	return kitHttp.NewServer(
		mux,
		kitHttp.Network("tcp"),
		kitHttp.Address(":10010"),
	)
}
