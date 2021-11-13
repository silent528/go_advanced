package main

import (
	"go_advanced/library/kit"
	"go_advanced/library/kit/grpc"
	"go_advanced/library/kit/http"
)

// 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，
//以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。
func main() {
	app, cleanup, err := initApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Start(); err != nil {
		panic(err)
	}
}

func newApp(hs *http.Server, gs *grpc.Server) *kit.App {
	return kit.New(
		kit.Server(hs, gs),
	)
}
