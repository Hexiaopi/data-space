package app

import (
	"github.com/hexiaopi/data-space/internal/config"
	"github.com/hexiaopi/data-space/internal/router"
	"github.com/hexiaopi/data-space/pkg/app"
	"github.com/hexiaopi/data-space/pkg/server/http"
)

func Run() error {
	router := router.InitRouter()
	httpServer := http.NewServer(
		router,
		http.WithServerHost(config.HTTP.Host),
		http.WithServerPort(config.HTTP.Port),
		http.WithLogger(config.Logger),
	)
	if err := app.NewApp(app.WithServer(httpServer), app.WithLogger(config.Logger)).Run(); err != nil {
		return err
	}
	return nil
}
