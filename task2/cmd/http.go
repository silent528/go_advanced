package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func (s *Server) Start() error {
	log.Print("server start", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	log.Print("server stop", s.Addr)
	return s.Shutdown(ctx)
}

type App struct {
	ctx     context.Context
	cancel  func()
	servers []*Server
}

func New() *App {
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (a *App) Start() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	for _, srv := range a.servers {
		s := srv
		eg.Go(func() error {
			<-ctx.Done()
			toCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return s.Stop(toCtx)
		})
		eg.Go(func() error {
			return s.Start()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				log.Print("sign stop")
				a.cancel()
				return nil
			}
		}
	})
	return eg.Wait()
}

func NewServer(address string, handle http.Handler) *Server {
	srv := &Server{}
	srv.Server = &http.Server{
		Addr:    address,
		Handler: handle,
	}
	return srv
}
