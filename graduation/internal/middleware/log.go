package middleware

import (
	"context"
	"fmt"
	"go_advanced/library/kit/middleware"
)

func LoggingMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		fmt.Println("logging middleware in", req)
		reply, err = handler(ctx, req)
		fmt.Println("logging middleware out", reply)
		return
	}
}
