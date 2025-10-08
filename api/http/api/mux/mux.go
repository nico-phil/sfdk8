package mux

import (
	"context"

	"github.com/nico-phil/service/api/http/api/mid"
	"github.com/nico-phil/service/fondation/logger"
	"github.com/nico-phil/service/fondation/web"
)

type Config struct {
	Build string
	Log   *logger.Logger
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg Config)
}

func WebAPI(cfg Config, routeAdder RouteAdder) *web.App {

	logger := func(ctx context.Context, msg string, v ...any) {
		cfg.Log.Info(ctx, msg, v...)
	}
	app := web.NewApp(logger, mid.Logger(cfg.Log))

	routeAdder.Add(app, cfg)

	return app
}
