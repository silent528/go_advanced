package kit

import (
	"go_advanced/library/kit/ts"
)

type Options func(o *options)

type options struct {
	servers []ts.Server
}

func Server(servers ...ts.Server) Options {
	return func(o *options) { o.servers = servers }
}
