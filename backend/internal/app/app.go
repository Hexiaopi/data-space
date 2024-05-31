package app

import (
	"github.com/hexiaopi/data-space/internal/router"
	"github.com/hexiaopi/data-space/pkg/app"
	"github.com/hexiaopi/data-space/pkg/server/http"
)

func Run() error {
	router := router.InitRouter()
	httpServer := http.NewServer(
		router,
		http.WithServerHost("127.0.0.1"),
		http.WithServerPort(8080),
	)
	if err := app.NewApp(app.WithServer(httpServer)).Run(); err != nil {
		return err
	}
	return nil
}
