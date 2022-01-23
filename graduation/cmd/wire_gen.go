// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"go_advanced/graduation/internal/biz"
	"go_advanced/graduation/internal/data"
	"go_advanced/graduation/internal/server"
	"go_advanced/graduation/internal/service"
	"go_advanced/library/kit"
)

// Injectors from wire.go:

func initApp() (*kit.App, func(), error) {
	client := data.NewEntClient()
	dataData, cleanup, err := data.NewData(client)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData)
	userUsecase := biz.NewUserUsecase(userRepo)
	userService := service.NewUserService(userUsecase)
	httpServer := server.NewHttpServer(userService)
	grpcServer := server.NewGrpcServer(userService)
	app := newApp(httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
