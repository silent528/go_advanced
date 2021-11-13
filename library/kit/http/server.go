package http

import (
	"context"
	"log"
	"net"
	"net/http"
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
	*http.Server
	network string
	address string
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	log.Print("http server start", s.Addr)
	return s.Serve(lis)
}

func (s *Server) Stop(ctx context.Context) error {
	log.Print("http server stop", s.Addr)
	return s.Shutdown(ctx)
}

func NewServer(handle http.Handler, opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
	}
	for _, o := range opts {
		o(srv)
	}
	srv.Server = &http.Server{
		Handler: handle,
	}
	return srv
}
