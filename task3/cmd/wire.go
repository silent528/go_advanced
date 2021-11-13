// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"go_advanced/library/kit"
	"go_advanced/task3/internal/biz"
	"go_advanced/task3/internal/data"
	"go_advanced/task3/internal/server"
	"go_advanced/task3/internal/service"
)

func initApp() (*kit.App, func(), error) {
	// 依赖注入
	panic(wire.Build(server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet, newApp))
}
