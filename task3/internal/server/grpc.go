package server

import (
	kitGtpc "go_advanced/library/kit/grpc"
	v1 "go_advanced/task3/api/v1"
	"go_advanced/task3/internal/service"
)

func NewGrpcServer(greeter *service.UserService) *kitGtpc.Server {
	s := kitGtpc.NewServer(
		kitGtpc.Network("tcp"),
		kitGtpc.Address(":10011"),
	)
	v1.RegisterUserServiceServer(s, greeter)
	return s
}
