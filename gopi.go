package gopi

import "github.com/wardonne/gopi/app"

var application = app.New()

func WithRouter() {
	application.WithRouter()
}

func Configure(file string) {
	application.Configuration.MustLoad(file)
}
