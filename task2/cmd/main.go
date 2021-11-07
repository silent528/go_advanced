package main

import (
	"fmt"
	"net/http"
)

// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello word111")
	})

	s1 := NewServer(":10010", mux1)

	var app = New()
	app.servers = append(app.servers, s1)
	err := app.Start()
	fmt.Println(err)
}
