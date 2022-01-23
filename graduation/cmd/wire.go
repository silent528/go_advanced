// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"go_advanced/graduation/internal/biz"
	"go_advanced/graduation/internal/data"
	"go_advanced/graduation/internal/server"
	"go_advanced/graduation/internal/service"
	"go_advanced/library/kit"
)

func initApp() (*kit.App, func(), error) {
	// 依赖注入
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
