package checkapi

import (
	"github.com/nico-phil/service/fondation/logger"
	"github.com/nico-phil/service/fondation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
}

func Routes(app *web.App, cfg Config) {
	api := NewAPI(cfg.Build, cfg.Log)

	app.HandleFunc("GET /readiness", api.readiness)
	app.HandleFunc("GET /liveness", api.liveness)
	app.HandleFunc("GET /testerr", api.testerr)
}
