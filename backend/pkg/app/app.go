package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hexiaopi/data-space/pkg/server"
)

type App struct {
	name    string
	servers []server.Server
}

type Option func(a *App)

func NewApp(opts ...Option) *App {
	a := &App{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func WithServer(servers ...server.Server) Option {
	return func(a *App) {
		a.servers = servers
	}
}

func WithName(name string) Option {
	return func(a *App) {
		a.name = name
	}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, len(a.servers))

	for _, srv := range a.servers {
		go func(srv server.Server) {
			if err := srv.Run(ctx); err != nil {
				errChan <- err
			}
		}(srv)
	}

	select {
	case <-signals:
		log.Println("Received termination signal")
		cancel()
	case err := <-errChan:
		cancel()
		return err
	}
	return nil
}
