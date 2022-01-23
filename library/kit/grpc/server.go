package grpc

import (
	"context"
	"go_advanced/library/kit/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

type ServerOption func(*Server)

func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

func Address(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

func Middleware(m ...middleware.Middleware) ServerOption {
	return func(s *Server) {
		s.middleware = m
	}
}

type Server struct {
	*grpc.Server
	health     *health.Server
	network    string
	address    string
	middleware []middleware.Middleware
}

func (s *Server) Start(ctx context.Context) error {
	log.Print("grpc server start")
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

func (s *Server) Stop(ctx context.Context) error {
	log.Print("grpc server stop")
	s.GracefulStop()
	s.health.Shutdown()
	return nil
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		health:  health.NewServer(),
	}
	for _, o := range opts {
		o(srv)
	}
	interceptor := []grpc.UnaryServerInterceptor{
		srv.unaryServerInterceptor(),
	}
	grpcOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptor...),
	}
	srv.Server = grpc.NewServer(grpcOptions...)
	grpc_health_v1.RegisterHealthServer(srv.Server, srv.health)
	return srv
}

// 中间件
func (s *Server) unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
		if len(s.middleware) > 0 {
			h = middleware.Chain(s.middleware...)(h)
		}
		reply, err := h(ctx, req)
		return reply, err
	}
}
