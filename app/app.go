package app

import (
	"os"
	"path/filepath"

	"github.com/wardonne/gopi/config"
	"github.com/wardonne/gopi/logger"
	"github.com/wardonne/gopi/router"
)

type App struct {
	BasePath      string
	Router        *router.Router
	Logger        *logger.Logger
	Configuration *config.Configuration
}

func New() *App {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	app := &App{
		BasePath: filepath.Dir(executable),
	}
	app.Configuration = config.New()
	return app
}

func (app *App) WithRouter() {
	app.Router = router.New()
}

func (app *App) WithValidator() {
}
