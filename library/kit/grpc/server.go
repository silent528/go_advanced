package grpc

import (
	"context"
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

type Server struct {
	*grpc.Server
	health  *health.Server
	network string
	address string
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
	srv.Server = grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(srv.Server, srv.health)
	return srv
}
