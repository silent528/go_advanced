package server

import (
	"go_advanced/graduation/internal/middleware"
	"go_advanced/graduation/internal/service"
	kitGrpc "go_advanced/library/kit/grpc"
	v1 "go_advanced/task3/api/v1"
)

func NewGrpcServer(greeter *service.UserService) *kitGrpc.Server {
	s := kitGrpc.NewServer(
		kitGrpc.Network("tcp"),
		kitGrpc.Address(":10011"),
		kitGrpc.Middleware(middleware.LoggingMiddleware),
	)
	v1.RegisterUserServiceServer(s, greeter)
	return s
}
