package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/hexiaopi/data-space/pkg/logger"
)

type Server struct {
	engine  *http.Server
	handler http.Handler
	log     logger.Logger
	host    string
	port    int
}
type Option func(s *Server)

func NewServer(handler http.Handler, opts ...Option) *Server {
	s := &Server{
		handler: handler,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}
func WithServerPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithLogger(log logger.Logger) Option {
	return func(s *Server) {
		s.log = log
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.engine = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s.handler,
	}
	s.log.Infof("Server is running on %s:%d", s.host, s.port)
	if err := s.engine.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.Errorf("listen: %s\n", err)
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	s.log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.engine.Shutdown(ctx); err != nil {
		s.log.Errorf("Server forced to shutdown err:%", err)
	}
	s.log.Info("Server exiting")
	return nil
}

func (s *Server) Run(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	errChan := make(chan error)
	go func() {
		if err := s.Start(ctx); err != nil {
			errChan <- err
		}
	}()
	select {
	case <-ctx.Done():
		s.Stop()
		return nil
	case err := <-errChan:
		return err
	}
}
