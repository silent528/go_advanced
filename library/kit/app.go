package kit

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	ctx     context.Context
	cancel  func()
	options options
}

func New(opts ...Options) *App {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx:     ctx,
		cancel:  cancel,
		options: o,
	}
}

func (a *App) Start() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	for _, srv := range a.options.servers {
		s := srv
		eg.Go(func() error {
			<-ctx.Done()
			toCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return s.Stop(toCtx)
		})
		eg.Go(func() error {
			return s.Start(ctx)
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
				log.Print("signal stop")
				a.cancel()
				return nil
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
